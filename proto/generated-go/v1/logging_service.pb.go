// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/logging_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogEntry_Action int32

const (
	LogEntry_ACTION_UNSPECIFIED LogEntry_Action = 0
	// In worksapce resource only.
	//
	// ACTION_MEMBER_CREATE is the type for creating a new member.
	LogEntry_ACTION_MEMBER_CREATE LogEntry_Action = 1
	// ACTION_MEMBER_ROLE_UPDATE is the type for updating a member's role.
	LogEntry_ACTION_MEMBER_ROLE_UPDATE LogEntry_Action = 2
	// ACTION_MEMBER_ACTIVATE_UPDATE is the type for activating members.
	LogEntry_ACTION_MEMBER_ACTIVATE LogEntry_Action = 3
	// ACTION_MEMBER_DEACTIVE is the type for deactiving members.
	LogEntry_ACTION_MEMBER_DEACTIVE LogEntry_Action = 4
	// In project resource only.
	//
	// ACTION_ISSUE_CREATE is the type for creating a new issue.
	LogEntry_ACTION_ISSUE_CREATE LogEntry_Action = 5
	// ACTION_ISSUE_COMMENT_CREATE is the type for creating a new comment on an issue.
	LogEntry_ACTION_ISSUE_COMMENT_CREATE LogEntry_Action = 6
	// ACTION_ISSUE_FIELD_UPDATE is the type for updating an issue's field.
	LogEntry_ACTION_ISSUE_FIELD_UPDATE LogEntry_Action = 7
	// ACTION_ISSUE_STATUS_UPDATE is the type for updating an issue's status.
	LogEntry_ACTION_ISSUE_STATUS_UPDATE LogEntry_Action = 8
	// ACTION_PIPELINE_STAGE_STATUS_UPDATE is the type for stage begins or ends.
	LogEntry_ACTION_PIPELINE_STAGE_STATUS_UPDATE LogEntry_Action = 9
	// ACTION_PIPELINE_TASK_STATUS_UPDATE is the type for updating pipeline task status.
	LogEntry_ACTION_PIPELINE_TASK_STATUS_UPDATE LogEntry_Action = 10
	// ACTION_PIPELINE_TASK_FILE_COMMIT is the type for committing pipeline task files.
	LogEntry_ACTION_PIPELINE_TASK_FILE_COMMIT LogEntry_Action = 11
	// ACTION_PIPELINE_TASK_STATEMENT_UPDATE is the type for updating pipeline task SQL statement.
	LogEntry_ACTION_PIPELINE_TASK_STATEMENT_UPDATE LogEntry_Action = 12
	// ACTION_PIPELINE_TASK_EARLIEST_ALLOWED_DATE_UPDATE is the type for updating pipeline task the earliest allowed time.
	LogEntry_ACITON_PIPELINE_TASK_EARLIEST_ALLOWED_DATE_UPDATE LogEntry_Action = 13
	// ACTION_PROJECT_MEMBER_CREATE is the type for creating a new project member.
	LogEntry_ACTION_PROJECT_MEMBER_CREATE LogEntry_Action = 14
	// ACTION_PROJECT_MEMBER_ROLE_UPDATE is the type for updating a project member's role.
	LogEntry_ACTION_PROJECT_MEMBER_ROLE_UPDATE LogEntry_Action = 15
	// ACTION_PROJECT_MEMBER_DELETE is the type for deleting a project member.
	LogEntry_ACTION_PROJECT_MEMBER_DELETE LogEntry_Action = 16
	// ACTION_PROJECT_REPOSITORY_PUSH is the type for pushing to a project repository.
	LogEntry_ACTION_PROJECT_REPOSITORY_PUSH LogEntry_Action = 17
	// ACTION_PROJECT_DATABASE_TRANSFER is the type for transferring a database to a project.
	LogEntry_ACTION_PROJECT_DTABASE_TRANSFER LogEntry_Action = 18
	// ACTION_PROJECT_DATABASE_RECOVERY_PITR_DONE is the type for database PITR recovery done.
	LogEntry_ACTION_PROJECT_DATABASE_RECOVERY_PITR_DONE LogEntry_Action = 19
	// Both in workspace and project resource.
	//
	// ACTION_SQL_EDITOR_QUERY is the type for SQL editor query.
	// If user runs SQL in Read-only mode, this action will belong to project resource.
	// If user runs SQL in Read-write mode, this action will belong to workspace resource.
	LogEntry_ACTION_SQL_EDITOR_QUERY LogEntry_Action = 20
)

// Enum value maps for LogEntry_Action.
var (
	LogEntry_Action_name = map[int32]string{
		0:  "ACTION_UNSPECIFIED",
		1:  "ACTION_MEMBER_CREATE",
		2:  "ACTION_MEMBER_ROLE_UPDATE",
		3:  "ACTION_MEMBER_ACTIVATE",
		4:  "ACTION_MEMBER_DEACTIVE",
		5:  "ACTION_ISSUE_CREATE",
		6:  "ACTION_ISSUE_COMMENT_CREATE",
		7:  "ACTION_ISSUE_FIELD_UPDATE",
		8:  "ACTION_ISSUE_STATUS_UPDATE",
		9:  "ACTION_PIPELINE_STAGE_STATUS_UPDATE",
		10: "ACTION_PIPELINE_TASK_STATUS_UPDATE",
		11: "ACTION_PIPELINE_TASK_FILE_COMMIT",
		12: "ACTION_PIPELINE_TASK_STATEMENT_UPDATE",
		13: "ACITON_PIPELINE_TASK_EARLIEST_ALLOWED_DATE_UPDATE",
		14: "ACTION_PROJECT_MEMBER_CREATE",
		15: "ACTION_PROJECT_MEMBER_ROLE_UPDATE",
		16: "ACTION_PROJECT_MEMBER_DELETE",
		17: "ACTION_PROJECT_REPOSITORY_PUSH",
		18: "ACTION_PROJECT_DTABASE_TRANSFER",
		19: "ACTION_PROJECT_DATABASE_RECOVERY_PITR_DONE",
		20: "ACTION_SQL_EDITOR_QUERY",
	}
	LogEntry_Action_value = map[string]int32{
		"ACTION_UNSPECIFIED":                                0,
		"ACTION_MEMBER_CREATE":                              1,
		"ACTION_MEMBER_ROLE_UPDATE":                         2,
		"ACTION_MEMBER_ACTIVATE":                            3,
		"ACTION_MEMBER_DEACTIVE":                            4,
		"ACTION_ISSUE_CREATE":                               5,
		"ACTION_ISSUE_COMMENT_CREATE":                       6,
		"ACTION_ISSUE_FIELD_UPDATE":                         7,
		"ACTION_ISSUE_STATUS_UPDATE":                        8,
		"ACTION_PIPELINE_STAGE_STATUS_UPDATE":               9,
		"ACTION_PIPELINE_TASK_STATUS_UPDATE":                10,
		"ACTION_PIPELINE_TASK_FILE_COMMIT":                  11,
		"ACTION_PIPELINE_TASK_STATEMENT_UPDATE":             12,
		"ACITON_PIPELINE_TASK_EARLIEST_ALLOWED_DATE_UPDATE": 13,
		"ACTION_PROJECT_MEMBER_CREATE":                      14,
		"ACTION_PROJECT_MEMBER_ROLE_UPDATE":                 15,
		"ACTION_PROJECT_MEMBER_DELETE":                      16,
		"ACTION_PROJECT_REPOSITORY_PUSH":                    17,
		"ACTION_PROJECT_DTABASE_TRANSFER":                   18,
		"ACTION_PROJECT_DATABASE_RECOVERY_PITR_DONE":        19,
		"ACTION_SQL_EDITOR_QUERY":                           20,
	}
)

func (x LogEntry_Action) Enum() *LogEntry_Action {
	p := new(LogEntry_Action)
	*p = x
	return p
}

func (x LogEntry_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogEntry_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_logging_service_proto_enumTypes[0].Descriptor()
}

func (LogEntry_Action) Type() protoreflect.EnumType {
	return &file_v1_logging_service_proto_enumTypes[0]
}

func (x LogEntry_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogEntry_Action.Descriptor instead.
func (LogEntry_Action) EnumDescriptor() ([]byte, []int) {
	return file_v1_logging_service_proto_rawDescGZIP(), []int{2, 0}
}

type LogEntry_Level int32

const (
	LogEntry_LEVEL_UNSPECIFIED LogEntry_Level = 0
	// LEVEL_INFO is the type for information.
	LogEntry_LEVEL_INFO LogEntry_Level = 1
	// LEVEL_WARNING is the type for warning.
	LogEntry_LEVEL_WARNING LogEntry_Level = 2
	// LEVEL_ERROR is the type for error.
	LogEntry_LEVEL_ERROR LogEntry_Level = 3
)

// Enum value maps for LogEntry_Level.
var (
	LogEntry_Level_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "LEVEL_INFO",
		2: "LEVEL_WARNING",
		3: "LEVEL_ERROR",
	}
	LogEntry_Level_value = map[string]int32{
		"LEVEL_UNSPECIFIED": 0,
		"LEVEL_INFO":        1,
		"LEVEL_WARNING":     2,
		"LEVEL_ERROR":       3,
	}
)

func (x LogEntry_Level) Enum() *LogEntry_Level {
	p := new(LogEntry_Level)
	*p = x
	return p
}

func (x LogEntry_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogEntry_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_logging_service_proto_enumTypes[1].Descriptor()
}

func (LogEntry_Level) Type() protoreflect.EnumType {
	return &file_v1_logging_service_proto_enumTypes[1]
}

func (x LogEntry_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogEntry_Level.Descriptor instead.
func (LogEntry_Level) EnumDescriptor() ([]byte, []int) {
	return file_v1_logging_service_proto_rawDescGZIP(), []int{2, 1}
}

type ListLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource name.
	// Format:
	// projects/{project}
	// workspaces/{workspace}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// filter is the filter to apply on the list logs request,
	// follow the [ebnf](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form) syntax.
	// The field only support in filter:
	// - creator
	// - container
	// - level
	// - action
	// For example:
	// List the logs of type 'ACTION_ISSUE_COMMENT_CREATE' in issue/123: 'action="ACTION_ISSUE_COMMENT_CREATE", container="issue/123"'
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Not used. The maximum number of logs to return.
	// The service may return fewer than this value.
	// If unspecified, at most 100 log entries will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Not used. A page token, received from a previous `ListLogs` call.
	// Provide this to retrieve the subsequent page.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListLogsRequest) Reset() {
	*x = ListLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_logging_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLogsRequest) ProtoMessage() {}

func (x *ListLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_logging_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLogsRequest.ProtoReflect.Descriptor instead.
func (*ListLogsRequest) Descriptor() ([]byte, []int) {
	return file_v1_logging_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListLogsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListLogsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListLogsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListLogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of log entries.
	LogEntries []*LogEntry `protobuf:"bytes,1,rep,name=log_entries,json=logEntries,proto3" json:"log_entries,omitempty"`
	// A token to retrieve next page of log entries.
	// Pass this value in the page_token field in the subsequent call to `ListLogs` method
	// to retrieve the next page of log entries.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListLogsResponse) Reset() {
	*x = ListLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_logging_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLogsResponse) ProtoMessage() {}

func (x *ListLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_logging_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLogsResponse.ProtoReflect.Descriptor instead.
func (*ListLogsResponse) Descriptor() ([]byte, []int) {
	return file_v1_logging_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListLogsResponse) GetLogEntries() []*LogEntry {
	if x != nil {
		return x.LogEntries
	}
	return nil
}

func (x *ListLogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The creator of the log entry.
	// Format: user:{emailid}
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// The timestamp when the backup resource was created initally.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The timestamp when the backup resource was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Action     LogEntry_Action        `protobuf:"varint,4,opt,name=action,proto3,enum=bytebase.v1.LogEntry_Action" json:"action,omitempty"`
	Level      LogEntry_Level         `protobuf:"varint,5,opt,name=level,proto3,enum=bytebase.v1.LogEntry_Level" json:"level,omitempty"`
	// The name of the resource associated with this log entry. For example, the resource user associated with log entry type of "ACTION_MEMBER_CREATE".
	// Format:
	// For ACTION_MEMBER_*: user:emailid
	// For ACTION_ISSUE_*: issues/{issue}
	// For ACTION_PIPELINE_*: pipelines/{pipeline}
	// For ACTION_PROJECT_*: projects/{project}
	// For ACTION_SQL_EDITOR_QUERY: workspaces/{workspace} OR projects/{project}
	ResourceName string `protobuf:"bytes,6,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The payload of the log entry.
	JsonPayload *structpb.Struct `protobuf:"bytes,7,opt,name=json_payload,json=jsonPayload,proto3" json:"json_payload,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_logging_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_logging_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_v1_logging_service_proto_rawDescGZIP(), []int{2}
}

func (x *LogEntry) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *LogEntry) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *LogEntry) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *LogEntry) GetAction() LogEntry_Action {
	if x != nil {
		return x.Action
	}
	return LogEntry_ACTION_UNSPECIFIED
}

func (x *LogEntry) GetLevel() LogEntry_Level {
	if x != nil {
		return x.Level
	}
	return LogEntry_LEVEL_UNSPECIFIED
}

func (x *LogEntry) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *LogEntry) GetJsonPayload() *structpb.Struct {
	if x != nil {
		return x.JsonPayload
	}
	return nil
}

var File_v1_logging_service_proto protoreflect.FileDescriptor

var file_v1_logging_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82,
	0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x77, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc1, 0x09, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x28, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xe4, 0x05, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x4d,
	0x42, 0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a,
	0x16, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x44,
	0x45, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x53, 0x53,
	0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x53,
	0x53, 0x55, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x53, 0x53,
	0x55, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x10, 0x08, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x49, 0x50,
	0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x09, 0x12, 0x26, 0x0a, 0x22, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x0a, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x49,
	0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x0b, 0x12, 0x29, 0x0a, 0x25, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x0c, 0x12, 0x35, 0x0a, 0x31, 0x41, 0x43, 0x49, 0x54, 0x4f, 0x4e, 0x5f, 0x50,
	0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x45, 0x41, 0x52,
	0x4c, 0x49, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x0d, 0x12, 0x20, 0x0a, 0x1c, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4d, 0x45,
	0x4d, 0x42, 0x45, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x0e, 0x12, 0x25, 0x0a,
	0x21, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x0f, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x10, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x4f, 0x52, 0x59, 0x5f, 0x50, 0x55, 0x53, 0x48, 0x10, 0x11, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x54, 0x41,
	0x42, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x12, 0x12,
	0x2e, 0x0a, 0x2a, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43,
	0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x56,
	0x45, 0x52, 0x59, 0x5f, 0x50, 0x49, 0x54, 0x52, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x13, 0x12,
	0x1b, 0x0a, 0x17, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x51, 0x4c, 0x5f, 0x45, 0x44,
	0x49, 0x54, 0x4f, 0x52, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x14, 0x22, 0x52, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03,
	0x32, 0xab, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73,
	0x12, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0xda,
	0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x40, 0x5a, 0x20,
	0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6c, 0x6f, 0x67, 0x73,
	0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x11,
	0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_logging_service_proto_rawDescOnce sync.Once
	file_v1_logging_service_proto_rawDescData = file_v1_logging_service_proto_rawDesc
)

func file_v1_logging_service_proto_rawDescGZIP() []byte {
	file_v1_logging_service_proto_rawDescOnce.Do(func() {
		file_v1_logging_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_logging_service_proto_rawDescData)
	})
	return file_v1_logging_service_proto_rawDescData
}

var file_v1_logging_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_logging_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_logging_service_proto_goTypes = []interface{}{
	(LogEntry_Action)(0),          // 0: bytebase.v1.LogEntry.Action
	(LogEntry_Level)(0),           // 1: bytebase.v1.LogEntry.Level
	(*ListLogsRequest)(nil),       // 2: bytebase.v1.ListLogsRequest
	(*ListLogsResponse)(nil),      // 3: bytebase.v1.ListLogsResponse
	(*LogEntry)(nil),              // 4: bytebase.v1.LogEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 6: google.protobuf.Struct
}
var file_v1_logging_service_proto_depIdxs = []int32{
	4, // 0: bytebase.v1.ListLogsResponse.log_entries:type_name -> bytebase.v1.LogEntry
	5, // 1: bytebase.v1.LogEntry.create_time:type_name -> google.protobuf.Timestamp
	5, // 2: bytebase.v1.LogEntry.update_time:type_name -> google.protobuf.Timestamp
	0, // 3: bytebase.v1.LogEntry.action:type_name -> bytebase.v1.LogEntry.Action
	1, // 4: bytebase.v1.LogEntry.level:type_name -> bytebase.v1.LogEntry.Level
	6, // 5: bytebase.v1.LogEntry.json_payload:type_name -> google.protobuf.Struct
	2, // 6: bytebase.v1.LoggingService.ListLogs:input_type -> bytebase.v1.ListLogsRequest
	3, // 7: bytebase.v1.LoggingService.ListLogs:output_type -> bytebase.v1.ListLogsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_logging_service_proto_init() }
func file_v1_logging_service_proto_init() {
	if File_v1_logging_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_logging_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLogsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_logging_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLogsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_logging_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_logging_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_logging_service_proto_goTypes,
		DependencyIndexes: file_v1_logging_service_proto_depIdxs,
		EnumInfos:         file_v1_logging_service_proto_enumTypes,
		MessageInfos:      file_v1_logging_service_proto_msgTypes,
	}.Build()
	File_v1_logging_service_proto = out.File
	file_v1_logging_service_proto_rawDesc = nil
	file_v1_logging_service_proto_goTypes = nil
	file_v1_logging_service_proto_depIdxs = nil
}
