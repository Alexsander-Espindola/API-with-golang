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
	PostUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
}

type postUserClient struct {
	cc grpc.ClientConnInterface
}

func NewPostUserClient(cc grpc.ClientConnInterface) PostUserClient {
	return &postUserClient{cc}
}

func (c *postUserClient) PostUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/protobuff.PostUser/PostUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostUserServer is the server API for PostUser service.
// All implementations must embed UnimplementedPostUserServer
// for forward compatibility
type PostUserServer interface {
	PostUser(context.Context, *User) (*UserResponse, error)
	mustEmbedUnimplementedPostUserServer()
}

// UnimplementedPostUserServer must be embedded to have forward compatible implementations.
type UnimplementedPostUserServer struct {
}

func (UnimplementedPostUserServer) PostUser(context.Context, *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUser not implemented")
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
	in := new(User)
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
		return srv.(PostUserServer).PostUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/userProto.proto",
}
