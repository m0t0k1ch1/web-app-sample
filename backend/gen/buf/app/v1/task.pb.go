// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: app/v1/task.proto

package appv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskStatus int32

const (
	TaskStatus_TASK_STATUS_UNSPECIFIED TaskStatus = 0
	TaskStatus_TASK_STATUS_UNCOMPLETED TaskStatus = 1
	TaskStatus_TASK_STATUS_COMPLETED   TaskStatus = 2
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "TASK_STATUS_UNSPECIFIED",
		1: "TASK_STATUS_UNCOMPLETED",
		2: "TASK_STATUS_COMPLETED",
	}
	TaskStatus_value = map[string]int32{
		"TASK_STATUS_UNSPECIFIED": 0,
		"TASK_STATUS_UNCOMPLETED": 1,
		"TASK_STATUS_COMPLETED":   2,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_app_v1_task_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_app_v1_task_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{0}
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Status    TaskStatus `protobuf:"varint,3,opt,name=status,proto3,enum=app.v1.TaskStatus" json:"status,omitempty"`
	UpdatedAt int64      `protobuf:"varint,1001,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt int64      `protobuf:"varint,1002,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_TASK_STATUS_UNSPECIFIED
}

func (x *Task) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Task) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type TaskServiceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *TaskServiceCreateRequest) Reset() {
	*x = TaskServiceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceCreateRequest) ProtoMessage() {}

func (x *TaskServiceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceCreateRequest.ProtoReflect.Descriptor instead.
func (*TaskServiceCreateRequest) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskServiceCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type TaskServiceCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TaskServiceCreateResponse) Reset() {
	*x = TaskServiceCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceCreateResponse) ProtoMessage() {}

func (x *TaskServiceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceCreateResponse.ProtoReflect.Descriptor instead.
func (*TaskServiceCreateResponse) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskServiceCreateResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type TaskServiceGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskServiceGetRequest) Reset() {
	*x = TaskServiceGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceGetRequest) ProtoMessage() {}

func (x *TaskServiceGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceGetRequest.ProtoReflect.Descriptor instead.
func (*TaskServiceGetRequest) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskServiceGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TaskServiceGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TaskServiceGetResponse) Reset() {
	*x = TaskServiceGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceGetResponse) ProtoMessage() {}

func (x *TaskServiceGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceGetResponse.ProtoReflect.Descriptor instead.
func (*TaskServiceGetResponse) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{4}
}

func (x *TaskServiceGetResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type TaskServiceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskServiceListRequest) Reset() {
	*x = TaskServiceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceListRequest) ProtoMessage() {}

func (x *TaskServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceListRequest.ProtoReflect.Descriptor instead.
func (*TaskServiceListRequest) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{5}
}

type TaskServiceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskServiceListResponse) Reset() {
	*x = TaskServiceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceListResponse) ProtoMessage() {}

func (x *TaskServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceListResponse.ProtoReflect.Descriptor instead.
func (*TaskServiceListResponse) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{6}
}

func (x *TaskServiceListResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TaskServiceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Status TaskStatus `protobuf:"varint,3,opt,name=status,proto3,enum=app.v1.TaskStatus" json:"status,omitempty"`
}

func (x *TaskServiceUpdateRequest) Reset() {
	*x = TaskServiceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceUpdateRequest) ProtoMessage() {}

func (x *TaskServiceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceUpdateRequest.ProtoReflect.Descriptor instead.
func (*TaskServiceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{7}
}

func (x *TaskServiceUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TaskServiceUpdateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskServiceUpdateRequest) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_TASK_STATUS_UNSPECIFIED
}

type TaskServiceUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TaskServiceUpdateResponse) Reset() {
	*x = TaskServiceUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceUpdateResponse) ProtoMessage() {}

func (x *TaskServiceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceUpdateResponse.ProtoReflect.Descriptor instead.
func (*TaskServiceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{8}
}

func (x *TaskServiceUpdateResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type TaskServiceDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskServiceDeleteRequest) Reset() {
	*x = TaskServiceDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceDeleteRequest) ProtoMessage() {}

func (x *TaskServiceDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceDeleteRequest.ProtoReflect.Descriptor instead.
func (*TaskServiceDeleteRequest) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{9}
}

func (x *TaskServiceDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TaskServiceDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskServiceDeleteResponse) Reset() {
	*x = TaskServiceDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_v1_task_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceDeleteResponse) ProtoMessage() {}

func (x *TaskServiceDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_v1_task_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceDeleteResponse.ProtoReflect.Descriptor instead.
func (*TaskServiceDeleteResponse) Descriptor() ([]byte, []int) {
	return file_app_v1_task_proto_rawDescGZIP(), []int{10}
}

var File_app_v1_task_proto protoreflect.FileDescriptor

var file_app_v1_task_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0xe9, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0xea, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x3b, 0x0a, 0x18, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x01, 0x18, 0x20, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3d, 0x0a, 0x19,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x61, 0x73,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x30, 0x0a, 0x15, 0x54,
	0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a,
	0x16, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x18, 0x0a, 0x16, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x17, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x18, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01,
	0x18, 0x20, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x82, 0x01, 0x04, 0x18, 0x01, 0x18, 0x02, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x3d, 0x0a, 0x19, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x22, 0x33, 0x0a, 0x18, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2a, 0x61, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1b, 0x0a, 0x17, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a,
	0x17, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0x93, 0x03, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x6a, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x18, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x70, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x70, 0x70, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x06, 0x41, 0x70, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x70, 0x70, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07,
	0x41, 0x70, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_v1_task_proto_rawDescOnce sync.Once
	file_app_v1_task_proto_rawDescData = file_app_v1_task_proto_rawDesc
)

func file_app_v1_task_proto_rawDescGZIP() []byte {
	file_app_v1_task_proto_rawDescOnce.Do(func() {
		file_app_v1_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_v1_task_proto_rawDescData)
	})
	return file_app_v1_task_proto_rawDescData
}

var file_app_v1_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_v1_task_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_app_v1_task_proto_goTypes = []interface{}{
	(TaskStatus)(0),                   // 0: app.v1.TaskStatus
	(*Task)(nil),                      // 1: app.v1.Task
	(*TaskServiceCreateRequest)(nil),  // 2: app.v1.TaskServiceCreateRequest
	(*TaskServiceCreateResponse)(nil), // 3: app.v1.TaskServiceCreateResponse
	(*TaskServiceGetRequest)(nil),     // 4: app.v1.TaskServiceGetRequest
	(*TaskServiceGetResponse)(nil),    // 5: app.v1.TaskServiceGetResponse
	(*TaskServiceListRequest)(nil),    // 6: app.v1.TaskServiceListRequest
	(*TaskServiceListResponse)(nil),   // 7: app.v1.TaskServiceListResponse
	(*TaskServiceUpdateRequest)(nil),  // 8: app.v1.TaskServiceUpdateRequest
	(*TaskServiceUpdateResponse)(nil), // 9: app.v1.TaskServiceUpdateResponse
	(*TaskServiceDeleteRequest)(nil),  // 10: app.v1.TaskServiceDeleteRequest
	(*TaskServiceDeleteResponse)(nil), // 11: app.v1.TaskServiceDeleteResponse
}
var file_app_v1_task_proto_depIdxs = []int32{
	0,  // 0: app.v1.Task.status:type_name -> app.v1.TaskStatus
	1,  // 1: app.v1.TaskServiceCreateResponse.task:type_name -> app.v1.Task
	1,  // 2: app.v1.TaskServiceGetResponse.task:type_name -> app.v1.Task
	1,  // 3: app.v1.TaskServiceListResponse.tasks:type_name -> app.v1.Task
	0,  // 4: app.v1.TaskServiceUpdateRequest.status:type_name -> app.v1.TaskStatus
	1,  // 5: app.v1.TaskServiceUpdateResponse.task:type_name -> app.v1.Task
	2,  // 6: app.v1.TaskService.Create:input_type -> app.v1.TaskServiceCreateRequest
	4,  // 7: app.v1.TaskService.Get:input_type -> app.v1.TaskServiceGetRequest
	6,  // 8: app.v1.TaskService.List:input_type -> app.v1.TaskServiceListRequest
	8,  // 9: app.v1.TaskService.Update:input_type -> app.v1.TaskServiceUpdateRequest
	10, // 10: app.v1.TaskService.Delete:input_type -> app.v1.TaskServiceDeleteRequest
	3,  // 11: app.v1.TaskService.Create:output_type -> app.v1.TaskServiceCreateResponse
	5,  // 12: app.v1.TaskService.Get:output_type -> app.v1.TaskServiceGetResponse
	7,  // 13: app.v1.TaskService.List:output_type -> app.v1.TaskServiceListResponse
	9,  // 14: app.v1.TaskService.Update:output_type -> app.v1.TaskServiceUpdateResponse
	11, // 15: app.v1.TaskService.Delete:output_type -> app.v1.TaskServiceDeleteResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_app_v1_task_proto_init() }
func file_app_v1_task_proto_init() {
	if File_app_v1_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_v1_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_app_v1_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceCreateRequest); i {
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
		file_app_v1_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceCreateResponse); i {
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
		file_app_v1_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceGetRequest); i {
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
		file_app_v1_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceGetResponse); i {
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
		file_app_v1_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceListRequest); i {
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
		file_app_v1_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceListResponse); i {
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
		file_app_v1_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceUpdateRequest); i {
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
		file_app_v1_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceUpdateResponse); i {
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
		file_app_v1_task_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceDeleteRequest); i {
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
		file_app_v1_task_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskServiceDeleteResponse); i {
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
			RawDescriptor: file_app_v1_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_v1_task_proto_goTypes,
		DependencyIndexes: file_app_v1_task_proto_depIdxs,
		EnumInfos:         file_app_v1_task_proto_enumTypes,
		MessageInfos:      file_app_v1_task_proto_msgTypes,
	}.Build()
	File_app_v1_task_proto = out.File
	file_app_v1_task_proto_rawDesc = nil
	file_app_v1_task_proto_goTypes = nil
	file_app_v1_task_proto_depIdxs = nil
}
