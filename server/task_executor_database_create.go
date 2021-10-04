package server

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bytebase/bytebase/api"
	"github.com/bytebase/bytebase/plugin/db"
	"go.uber.org/zap"
)

func NewDatabaseCreateTaskExecutor(logger *zap.Logger) TaskExecutor {
	return &DatabaseCreateTaskExecutor{
		l: logger,
	}
}

type DatabaseCreateTaskExecutor struct {
	l *zap.Logger
}

func (exec *DatabaseCreateTaskExecutor) RunOnce(ctx context.Context, server *Server, task *api.Task) (terminated bool, result *api.TaskRunResultPayload, err error) {
	defer func() {
		if r := recover(); r != nil {
			panicErr, ok := r.(error)
			if !ok {
				panicErr = fmt.Errorf("%v", r)
			}
			exec.l.Error("DatabaseCreateTaskExecutor PANIC RECOVER", zap.Error(panicErr))
			terminated = true
			err = fmt.Errorf("encounter internal error when creating database")
		}
	}()

	payload := &api.TaskDatabaseCreatePayload{}
	if err := json.Unmarshal([]byte(task.Payload), payload); err != nil {
		return true, nil, fmt.Errorf("invalid create database payload: %w", err)
	}

	if err := server.ComposeTaskRelationship(ctx, task); err != nil {
		return true, nil, err
	}

	instance := task.Instance
	driver, err := GetDatabaseDriver(task.Instance, "", exec.l)
	if err != nil {
		return true, nil, err
	}
	defer driver.Close(context.Background())

	var statement string
	switch instance.Engine {
	case db.MySQL:
		statement = fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET %s COLLATE %s", payload.DatabaseName, payload.CharacterSet, payload.Collation)
	case db.TiDB:
		statement = fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET %s COLLATE %s", payload.DatabaseName, payload.CharacterSet, payload.Collation)
	case db.Postgres:
		statement = fmt.Sprintf(`CREATE DATABASE "%s"`, payload.DatabaseName)
	}
	exec.l.Debug("Start creating database...",
		zap.String("instance", instance.Name),
		zap.String("database", payload.DatabaseName),
		zap.String("sql", statement),
	)

	// Create a baseline migration history upon creating the database.
	mi := &db.MigrationInfo{
		ReleaseVersion: server.version,
		Version:        defaultMigrationVersionFromTaskId(task.ID),
		Namespace:      payload.DatabaseName,
		Database:       payload.DatabaseName,
		Environment:    instance.Environment.Name,
		Engine:         db.UI,
		Type:           db.Baseline,
		Description:    "Create database",
		CreateDatabase: true,
	}
	creator, err := server.ComposePrincipalById(context.Background(), task.CreatorId)
	if err != nil {
		// If somehow we unable to find the principal, we just emit the error since it's not
		// critical enough to fail the entire operation.
		exec.l.Error("Failed to fetch creator for composing the migration info",
			zap.Int("task_id", task.ID),
			zap.Error(err),
		)
	} else {
		mi.Creator = creator.Name
	}

	issueFind := &api.IssueFind{
		PipelineId: &task.PipelineId,
	}
	issue, err := server.IssueService.FindIssue(ctx, issueFind)
	if err != nil {
		// If somehow we unable to find the issue, we just emit the error since it's not
		// critical enough to fail the entire operation.
		exec.l.Error("Failed to fetch containing issue for composing the migration info",
			zap.Int("task_id", task.ID),
			zap.Error(err),
		)
	} else {
		mi.IssueId = strconv.Itoa(issue.ID)
	}

	migrationId, _, err := driver.ExecuteMigration(ctx, mi, statement)
	if err != nil {
		return true, nil, err
	}

	return true, &api.TaskRunResultPayload{
		Detail:      fmt.Sprintf("Created database %q", payload.DatabaseName),
		MigrationId: migrationId,
		Version:     mi.Version,
	}, nil
}
