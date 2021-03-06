// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/userProto.proto

package pb

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

// PostUserClient is the client API for PostUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostUserClient interface {
	PostUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UserVote(ctx context.Context, in *UserVoteRequest, opts ...grpc.CallOption) (*UserVoteResponse, error)
	TestStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (PostUser_TestStreamClient, error)
}

type postUserClient struct {
	cc grpc.ClientConnInterface
}

func NewPostUserClient(cc grpc.ClientConnInterface) PostUserClient {
	return &postUserClient{cc}
}

func (c *postUserClient) PostUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/protobuff.PostUser/PostUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postUserClient) UserVote(ctx context.Context, in *UserVoteRequest, opts ...grpc.CallOption) (*UserVoteResponse, error) {
	out := new(UserVoteResponse)
	err := c.cc.Invoke(ctx, "/protobuff.PostUser/UserVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postUserClient) TestStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (PostUser_TestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &PostUser_ServiceDesc.Streams[0], "/protobuff.PostUser/TestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &postUserTestStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PostUser_TestStreamClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type postUserTestStreamClient struct {
	grpc.ClientStream
}

func (x *postUserTestStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PostUserServer is the server API for PostUser service.
// All implementations must embed UnimplementedPostUserServer
// for forward compatibility
type PostUserServer interface {
	PostUser(context.Context, *UserRequest) (*UserResponse, error)
	UserVote(context.Context, *UserVoteRequest) (*UserVoteResponse, error)
	TestStream(*StreamRequest, PostUser_TestStreamServer) error
	mustEmbedUnimplementedPostUserServer()
}

// UnimplementedPostUserServer must be embedded to have forward compatible implementations.
type UnimplementedPostUserServer struct {
}

func (UnimplementedPostUserServer) PostUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUser not implemented")
}
func (UnimplementedPostUserServer) UserVote(context.Context, *UserVoteRequest) (*UserVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserVote not implemented")
}
func (UnimplementedPostUserServer) TestStream(*StreamRequest, PostUser_TestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TestStream not implemented")
}
func (UnimplementedPostUserServer) mustEmbedUnimplementedPostUserServer() {}

// UnsafePostUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostUserServer will
// result in compilation errors.
type UnsafePostUserServer interface {
	mustEmbedUnimplementedPostUserServer()
}

func RegisterPostUserServer(s grpc.ServiceRegistrar, srv PostUserServer) {
	s.RegisterService(&PostUser_ServiceDesc, srv)
}

func _PostUser_PostUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostUserServer).PostUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuff.PostUser/PostUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostUserServer).PostUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostUser_UserVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostUserServer).UserVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuff.PostUser/UserVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostUserServer).UserVote(ctx, req.(*UserVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostUser_TestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostUserServer).TestStream(m, &postUserTestStreamServer{stream})
}

type PostUser_TestStreamServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type postUserTestStreamServer struct {
	grpc.ServerStream
}

func (x *postUserTestStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PostUser_ServiceDesc is the grpc.ServiceDesc for PostUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuff.PostUser",
	HandlerType: (*PostUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostUser",
			Handler:    _PostUser_PostUser_Handler,
		},
		{
			MethodName: "UserVote",
			Handler:    _PostUser_UserVote_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestStream",
			Handler:       _PostUser_TestStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/userProto.proto",
}
