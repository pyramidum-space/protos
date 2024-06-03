// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: motivator/motivator.proto

package motivator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             []byte                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId         int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartTime      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	WeightPerBlock int32                  `protobuf:"varint,4,opt,name=weight_per_block,json=weightPerBlock,proto3" json:"weight_per_block,omitempty"`
	DestroyTime    *durationpb.Duration   `protobuf:"bytes,5,opt,name=destroy_time,json=destroyTime,proto3" json:"destroy_time,omitempty"`
	Blocks         []*Block               `protobuf:"bytes,6,rep,name=blocks,proto3" json:"blocks,omitempty"`
	IsActive       bool                   `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	LastStopTime   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=last_stop_time,json=lastStopTime,proto3,oneof" json:"last_stop_time,omitempty"`
	TaskIds        [][]byte               `protobuf:"bytes,9,rep,name=task_ids,json=taskIds,proto3" json:"task_ids,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Session) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Session) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Session) GetWeightPerBlock() int32 {
	if x != nil {
		return x.WeightPerBlock
	}
	return 0
}

func (x *Session) GetDestroyTime() *durationpb.Duration {
	if x != nil {
		return x.DestroyTime
	}
	return nil
}

func (x *Session) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *Session) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Session) GetLastStopTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastStopTime
	}
	return nil
}

func (x *Session) GetTaskIds() [][]byte {
	if x != nil {
		return x.TaskIds
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Weight     int32                  `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	FinishTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Block) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Block) GetFinishTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishTime
	}
	return nil
}

type StartSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         int32                `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	WeightPerBlock int32                `protobuf:"varint,2,opt,name=weight_per_block,json=weightPerBlock,proto3" json:"weight_per_block,omitempty"`
	DestroyTime    *durationpb.Duration `protobuf:"bytes,3,opt,name=destroy_time,json=destroyTime,proto3" json:"destroy_time,omitempty"`
	TaskIds        [][]byte             `protobuf:"bytes,4,rep,name=task_ids,json=taskIds,proto3" json:"task_ids,omitempty"`
}

func (x *StartSessionRequest) Reset() {
	*x = StartSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSessionRequest) ProtoMessage() {}

func (x *StartSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSessionRequest.ProtoReflect.Descriptor instead.
func (*StartSessionRequest) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{2}
}

func (x *StartSessionRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *StartSessionRequest) GetWeightPerBlock() int32 {
	if x != nil {
		return x.WeightPerBlock
	}
	return 0
}

func (x *StartSessionRequest) GetDestroyTime() *durationpb.Duration {
	if x != nil {
		return x.DestroyTime
	}
	return nil
}

func (x *StartSessionRequest) GetTaskIds() [][]byte {
	if x != nil {
		return x.TaskIds
	}
	return nil
}

type StartSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartSessionResponse) Reset() {
	*x = StartSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSessionResponse) ProtoMessage() {}

func (x *StartSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSessionResponse.ProtoReflect.Descriptor instead.
func (*StartSessionResponse) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{3}
}

func (x *StartSessionResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type SessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SessionRequest) Reset() {
	*x = SessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionRequest) ProtoMessage() {}

func (x *SessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionRequest.ProtoReflect.Descriptor instead.
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{4}
}

func (x *SessionRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type SessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *SessionResponse) Reset() {
	*x = SessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionResponse) ProtoMessage() {}

func (x *SessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionResponse.ProtoReflect.Descriptor instead.
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{5}
}

func (x *SessionResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type CommitResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Weight int32  `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *CommitResultRequest) Reset() {
	*x = CommitResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitResultRequest) ProtoMessage() {}

func (x *CommitResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitResultRequest.ProtoReflect.Descriptor instead.
func (*CommitResultRequest) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{6}
}

func (x *CommitResultRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *CommitResultRequest) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type CommitResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CommitResultResponse) Reset() {
	*x = CommitResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitResultResponse) ProtoMessage() {}

func (x *CommitResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitResultResponse.ProtoReflect.Descriptor instead.
func (*CommitResultResponse) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{7}
}

func (x *CommitResultResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type StopSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StopSessionRequest) Reset() {
	*x = StopSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopSessionRequest) ProtoMessage() {}

func (x *StopSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopSessionRequest.ProtoReflect.Descriptor instead.
func (*StopSessionRequest) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{8}
}

func (x *StopSessionRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type StopSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StopSessionResponse) Reset() {
	*x = StopSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motivator_motivator_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopSessionResponse) ProtoMessage() {}

func (x *StopSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_motivator_motivator_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopSessionResponse.ProtoReflect.Descriptor instead.
func (*StopSessionResponse) Descriptor() ([]byte, []int) {
	return file_motivator_motivator_proto_rawDescGZIP(), []int{9}
}

func (x *StopSessionResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_motivator_motivator_proto protoreflect.FileDescriptor

var file_motivator_motivator_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x6f, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73,
	0x74, 0x6f, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x6c, 0x0a, 0x05, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3b, 0x0a, 0x0b,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x13, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x65, 0x72, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x73, 0x22, 0x26, 0x0a,
	0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x24, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc4, 0x02, 0x0a,
	0x10, 0x4d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e,
	0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x79, 0x72, 0x61, 0x6d, 0x69, 0x64, 0x75, 0x6d, 0x2d, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x6f, 0x72, 0x3b, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_motivator_motivator_proto_rawDescOnce sync.Once
	file_motivator_motivator_proto_rawDescData = file_motivator_motivator_proto_rawDesc
)

func file_motivator_motivator_proto_rawDescGZIP() []byte {
	file_motivator_motivator_proto_rawDescOnce.Do(func() {
		file_motivator_motivator_proto_rawDescData = protoimpl.X.CompressGZIP(file_motivator_motivator_proto_rawDescData)
	})
	return file_motivator_motivator_proto_rawDescData
}

var file_motivator_motivator_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_motivator_motivator_proto_goTypes = []interface{}{
	(*Session)(nil),               // 0: motivator.Session
	(*Block)(nil),                 // 1: motivator.Block
	(*StartSessionRequest)(nil),   // 2: motivator.StartSessionRequest
	(*StartSessionResponse)(nil),  // 3: motivator.StartSessionResponse
	(*SessionRequest)(nil),        // 4: motivator.SessionRequest
	(*SessionResponse)(nil),       // 5: motivator.SessionResponse
	(*CommitResultRequest)(nil),   // 6: motivator.CommitResultRequest
	(*CommitResultResponse)(nil),  // 7: motivator.CommitResultResponse
	(*StopSessionRequest)(nil),    // 8: motivator.StopSessionRequest
	(*StopSessionResponse)(nil),   // 9: motivator.StopSessionResponse
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 11: google.protobuf.Duration
}
var file_motivator_motivator_proto_depIdxs = []int32{
	10, // 0: motivator.Session.start_time:type_name -> google.protobuf.Timestamp
	11, // 1: motivator.Session.destroy_time:type_name -> google.protobuf.Duration
	1,  // 2: motivator.Session.blocks:type_name -> motivator.Block
	10, // 3: motivator.Session.last_stop_time:type_name -> google.protobuf.Timestamp
	10, // 4: motivator.Block.finish_time:type_name -> google.protobuf.Timestamp
	11, // 5: motivator.StartSessionRequest.destroy_time:type_name -> google.protobuf.Duration
	0,  // 6: motivator.SessionResponse.session:type_name -> motivator.Session
	2,  // 7: motivator.MotivatorService.StartSession:input_type -> motivator.StartSessionRequest
	4,  // 8: motivator.MotivatorService.Session:input_type -> motivator.SessionRequest
	6,  // 9: motivator.MotivatorService.CommitResult:input_type -> motivator.CommitResultRequest
	8,  // 10: motivator.MotivatorService.StopSession:input_type -> motivator.StopSessionRequest
	3,  // 11: motivator.MotivatorService.StartSession:output_type -> motivator.StartSessionResponse
	5,  // 12: motivator.MotivatorService.Session:output_type -> motivator.SessionResponse
	7,  // 13: motivator.MotivatorService.CommitResult:output_type -> motivator.CommitResultResponse
	9,  // 14: motivator.MotivatorService.StopSession:output_type -> motivator.StopSessionResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_motivator_motivator_proto_init() }
func file_motivator_motivator_proto_init() {
	if File_motivator_motivator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_motivator_motivator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_motivator_motivator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_motivator_motivator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartSessionRequest); i {
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
		file_motivator_motivator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartSessionResponse); i {
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
		file_motivator_motivator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionRequest); i {
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
		file_motivator_motivator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionResponse); i {
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
		file_motivator_motivator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitResultRequest); i {
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
		file_motivator_motivator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitResultResponse); i {
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
		file_motivator_motivator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopSessionRequest); i {
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
		file_motivator_motivator_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopSessionResponse); i {
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
	file_motivator_motivator_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_motivator_motivator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_motivator_motivator_proto_goTypes,
		DependencyIndexes: file_motivator_motivator_proto_depIdxs,
		MessageInfos:      file_motivator_motivator_proto_msgTypes,
	}.Build()
	File_motivator_motivator_proto = out.File
	file_motivator_motivator_proto_rawDesc = nil
	file_motivator_motivator_proto_goTypes = nil
	file_motivator_motivator_proto_depIdxs = nil
}
