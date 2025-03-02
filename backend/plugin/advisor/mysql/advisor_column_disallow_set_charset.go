package mysql

// Framework code is generated by the generator.

import (
	"fmt"

	"github.com/pingcap/tidb/parser/ast"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
)

var (
	_ advisor.Advisor = (*ColumnDisallowSetCharsetAdvisor)(nil)
	_ ast.Visitor     = (*columnDisallowSetCharsetChecker)(nil)
)

func init() {
	advisor.Register(db.MySQL, advisor.MySQLDisallowSetColumnCharset, &ColumnDisallowSetCharsetAdvisor{})
	advisor.Register(db.TiDB, advisor.MySQLDisallowSetColumnCharset, &ColumnDisallowSetCharsetAdvisor{})
	advisor.Register(db.MariaDB, advisor.MySQLDisallowSetColumnCharset, &ColumnDisallowSetCharsetAdvisor{})
}

// ColumnDisallowSetCharsetAdvisor is the advisor checking for disallow set column charset.
type ColumnDisallowSetCharsetAdvisor struct {
}

// Check checks for disallow set column charset.
func (*ColumnDisallowSetCharsetAdvisor) Check(ctx advisor.Context, statement string) ([]advisor.Advice, error) {
	stmtList, errAdvice := parseStatement(statement, ctx.Charset, ctx.Collation)
	if errAdvice != nil {
		return errAdvice, nil
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnDisallowSetCharsetChecker{
		level: level,
		title: string(ctx.Rule.Type),
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList, nil
}

type columnDisallowSetCharsetChecker struct {
	adviceList []advisor.Advice
	level      advisor.Status
	title      string
	text       string
	line       int
}

// Enter implements the ast.Visitor interface.
func (checker *columnDisallowSetCharsetChecker) Enter(in ast.Node) (ast.Node, bool) {
	code := advisor.Ok
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		for _, column := range node.Cols {
			charset := getColumnCharset(column)
			if !checkCharset(charset) {
				code = advisor.SetColumnCharset
				break
			}
		}
	case *ast.AlterTableStmt:
		for _, spec := range node.Specs {
			switch spec.Tp {
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					charset := getColumnCharset(column)
					if !checkCharset(charset) {
						code = advisor.SetColumnCharset
					}
				}
			case ast.AlterTableChangeColumn, ast.AlterTableModifyColumn:
				charset := getColumnCharset(spec.NewColumns[0])
				if !checkCharset(charset) {
					code = advisor.SetColumnCharset
				}
			}
			if code != advisor.Ok {
				break
			}
		}
	}

	if code != advisor.Ok {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  checker.level,
			Code:    advisor.SetColumnCharset,
			Title:   checker.title,
			Content: fmt.Sprintf("Disallow set column charset but \"%s\" does", checker.text),
			Line:    checker.line,
		})
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*columnDisallowSetCharsetChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func checkCharset(charset string) bool {
	switch charset {
	// empty charset or binary for JSON.
	case "", "binary":
		return true
	default:
		return false
	}
}
