// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.1
// source: raftRpc.proto

package rpc

import (
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

type HeartBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition string `protobuf:"bytes,3,opt,name=partition,proto3" json:"partition,omitempty"`
	Arg       []byte `protobuf:"bytes,1,opt,name=arg,proto3" json:"arg,omitempty"`
}

func (x *HeartBeatRequest) Reset() {
	*x = HeartBeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftRpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatRequest) ProtoMessage() {}

func (x *HeartBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raftRpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatRequest.ProtoReflect.Descriptor instead.
func (*HeartBeatRequest) Descriptor() ([]byte, []int) {
	return file_raftRpc_proto_rawDescGZIP(), []int{0}
}

func (x *HeartBeatRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *HeartBeatRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *HeartBeatRequest) GetArg() []byte {
	if x != nil {
		return x.Arg
	}
	return nil
}

type HeartBeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition string `protobuf:"bytes,3,opt,name=partition,proto3" json:"partition,omitempty"`
	Result    []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HeartBeatResponse) Reset() {
	*x = HeartBeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftRpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatResponse) ProtoMessage() {}

func (x *HeartBeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raftRpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatResponse.ProtoReflect.Descriptor instead.
func (*HeartBeatResponse) Descriptor() ([]byte, []int) {
	return file_raftRpc_proto_rawDescGZIP(), []int{1}
}

func (x *HeartBeatResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *HeartBeatResponse) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *HeartBeatResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

type RequestPreVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition string `protobuf:"bytes,3,opt,name=partition,proto3" json:"partition,omitempty"`
	Arg       []byte `protobuf:"bytes,1,opt,name=arg,proto3" json:"arg,omitempty"`
}

func (x *RequestPreVoteRequest) Reset() {
	*x = RequestPreVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftRpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPreVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPreVoteRequest) ProtoMessage() {}

func (x *RequestPreVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raftRpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPreVoteRequest.ProtoReflect.Descriptor instead.
func (*RequestPreVoteRequest) Descriptor() ([]byte, []int) {
	return file_raftRpc_proto_rawDescGZIP(), []int{2}
}

func (x *RequestPreVoteRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *RequestPreVoteRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *RequestPreVoteRequest) GetArg() []byte {
	if x != nil {
		return x.Arg
	}
	return nil
}

type RequestPreVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition string `protobuf:"bytes,3,opt,name=partition,proto3" json:"partition,omitempty"`
	Result    []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RequestPreVoteResponse) Reset() {
	*x = RequestPreVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftRpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPreVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPreVoteResponse) ProtoMessage() {}

func (x *RequestPreVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raftRpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPreVoteResponse.ProtoReflect.Descriptor instead.
func (*RequestPreVoteResponse) Descriptor() ([]byte, []int) {
	return file_raftRpc_proto_rawDescGZIP(), []int{3}
}

func (x *RequestPreVoteResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *RequestPreVoteResponse) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *RequestPreVoteResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

type RequestVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition string `protobuf:"bytes,3,opt,name=partition,proto3" json:"partition,omitempty"`
	Arg       []byte `protobuf:"bytes,1,opt,name=arg,proto3" json:"arg,omitempty"`
}

func (x *RequestVoteRequest) Reset() {
	*x = RequestVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftRpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteRequest) ProtoMessage() {}

func (x *RequestVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raftRpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteRequest.ProtoReflect.Descriptor instead.
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return file_raftRpc_proto_rawDescGZIP(), []int{4}
}

func (x *RequestVoteRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *RequestVoteRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *RequestVoteRequest) GetArg() []byte {
	if x != nil {
		return x.Arg
	}
	return nil
}

type RequestVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition string `protobuf:"bytes,3,opt,name=partition,proto3" json:"partition,omitempty"`
	Result    []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RequestVoteResponse) Reset() {
	*x = RequestVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftRpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteResponse) ProtoMessage() {}

func (x *RequestVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raftRpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteResponse.ProtoReflect.Descriptor instead.
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return file_raftRpc_proto_rawDescGZIP(), []int{5}
}

func (x *RequestVoteResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *RequestVoteResponse) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *RequestVoteResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_raftRpc_proto protoreflect.FileDescriptor

var file_raftRpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x61, 0x66, 0x74, 0x52, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x72, 0x61, 0x66, 0x74, 0x22, 0x58, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x61, 0x72, 0x67, 0x22,
	0x5f, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x5d, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x65, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x61, 0x72, 0x67, 0x22,
	0x64, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x65, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5a, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x61, 0x72,
	0x67, 0x22, 0x61, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0xdf, 0x01, 0x0a, 0x08, 0x52, 0x61, 0x66, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x3e, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x16,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x65, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x72, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12,
	0x18, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_raftRpc_proto_rawDescOnce sync.Once
	file_raftRpc_proto_rawDescData = file_raftRpc_proto_rawDesc
)

func file_raftRpc_proto_rawDescGZIP() []byte {
	file_raftRpc_proto_rawDescOnce.Do(func() {
		file_raftRpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_raftRpc_proto_rawDescData)
	})
	return file_raftRpc_proto_rawDescData
}

var file_raftRpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_raftRpc_proto_goTypes = []interface{}{
	(*HeartBeatRequest)(nil),       // 0: raft.HeartBeatRequest
	(*HeartBeatResponse)(nil),      // 1: raft.HeartBeatResponse
	(*RequestPreVoteRequest)(nil),  // 2: raft.RequestPreVoteRequest
	(*RequestPreVoteResponse)(nil), // 3: raft.RequestPreVoteResponse
	(*RequestVoteRequest)(nil),     // 4: raft.RequestVoteRequest
	(*RequestVoteResponse)(nil),    // 5: raft.RequestVoteResponse
}
var file_raftRpc_proto_depIdxs = []int32{
	0, // 0: raft.RaftCall.HeartBeat:input_type -> raft.HeartBeatRequest
	2, // 1: raft.RaftCall.RequestPreVote:input_type -> raft.RequestPreVoteRequest
	4, // 2: raft.RaftCall.RequestVote:input_type -> raft.RequestVoteRequest
	1, // 3: raft.RaftCall.HeartBeat:output_type -> raft.HeartBeatResponse
	3, // 4: raft.RaftCall.RequestPreVote:output_type -> raft.RequestPreVoteResponse
	5, // 5: raft.RaftCall.RequestVote:output_type -> raft.RequestVoteResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_raftRpc_proto_init() }
func file_raftRpc_proto_init() {
	if File_raftRpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raftRpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatRequest); i {
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
		file_raftRpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatResponse); i {
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
		file_raftRpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPreVoteRequest); i {
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
		file_raftRpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPreVoteResponse); i {
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
		file_raftRpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteRequest); i {
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
		file_raftRpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteResponse); i {
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
			RawDescriptor: file_raftRpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raftRpc_proto_goTypes,
		DependencyIndexes: file_raftRpc_proto_depIdxs,
		MessageInfos:      file_raftRpc_proto_msgTypes,
	}.Build()
	File_raftRpc_proto = out.File
	file_raftRpc_proto_rawDesc = nil
	file_raftRpc_proto_goTypes = nil
	file_raftRpc_proto_depIdxs = nil
}
