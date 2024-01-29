// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: raftRpc.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RaftCallClient is the client API for RaftCall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftCallClient interface {
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error)
	RequestPreVote(ctx context.Context, in *RequestPreVoteRequest, opts ...grpc.CallOption) (*RequestPreVoteResponse, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
}

type raftCallClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftCallClient(cc grpc.ClientConnInterface) RaftCallClient {
	return &raftCallClient{cc}
}

func (c *raftCallClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error) {
	out := new(HeartBeatResponse)
	err := c.cc.Invoke(ctx, "/raft.RaftCall/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCallClient) RequestPreVote(ctx context.Context, in *RequestPreVoteRequest, opts ...grpc.CallOption) (*RequestPreVoteResponse, error) {
	out := new(RequestPreVoteResponse)
	err := c.cc.Invoke(ctx, "/raft.RaftCall/RequestPreVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCallClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := c.cc.Invoke(ctx, "/raft.RaftCall/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftCallServer is the server API for RaftCall service.
// All implementations must embed UnimplementedRaftCallServer
// for forward compatibility
type RaftCallServer interface {
	HeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatResponse, error)
	RequestPreVote(context.Context, *RequestPreVoteRequest) (*RequestPreVoteResponse, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	mustEmbedUnimplementedRaftCallServer()
}

// UnimplementedRaftCallServer must be embedded to have forward compatible implementations.
type UnimplementedRaftCallServer struct {
}

func (UnimplementedRaftCallServer) HeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedRaftCallServer) RequestPreVote(context.Context, *RequestPreVoteRequest) (*RequestPreVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPreVote not implemented")
}
func (UnimplementedRaftCallServer) RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (UnimplementedRaftCallServer) mustEmbedUnimplementedRaftCallServer() {}

// UnsafeRaftCallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftCallServer will
// result in compilation errors.
type UnsafeRaftCallServer interface {
	mustEmbedUnimplementedRaftCallServer()
}

func RegisterRaftCallServer(s grpc.ServiceRegistrar, srv RaftCallServer) {
	s.RegisterService(&RaftCall_ServiceDesc, srv)
}

func _RaftCall_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCallServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftCall/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCallServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCall_RequestPreVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPreVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCallServer).RequestPreVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftCall/RequestPreVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCallServer).RequestPreVote(ctx, req.(*RequestPreVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCall_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCallServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftCall/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCallServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RaftCall_ServiceDesc is the grpc.ServiceDesc for RaftCall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftCall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raft.RaftCall",
	HandlerType: (*RaftCallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _RaftCall_HeartBeat_Handler,
		},
		{
			MethodName: "RequestPreVote",
			Handler:    _RaftCall_RequestPreVote_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _RaftCall_RequestVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raftRpc.proto",
}