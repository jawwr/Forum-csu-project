// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: userService.proto

package userService

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error)
	GetUserById(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserByIdClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userService.UserService/GetUserByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserById(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserByIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/userService.UserService/GetUserById", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserByIdClient{stream}
	return x, nil
}

type UserService_GetUserByIdClient interface {
	Send(*UserRequest) error
	Recv() (*User, error)
	grpc.ClientStream
}

type userServiceGetUserByIdClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserByIdClient) Send(m *UserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetUserByIdClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserByToken(context.Context, *Token) (*User, error)
	GetUserById(UserService_GetUserByIdServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserByToken(context.Context, *Token) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByToken not implemented")
}
func (UnimplementedUserServiceServer) GetUserById(UserService_GetUserByIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userService.UserService/GetUserByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserById_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetUserById(&userServiceGetUserByIdServer{stream})
}

type UserService_GetUserByIdServer interface {
	Send(*User) error
	Recv() (*UserRequest, error)
	grpc.ServerStream
}

type userServiceGetUserByIdServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserByIdServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetUserByIdServer) Recv() (*UserRequest, error) {
	m := new(UserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userService.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByToken",
			Handler:    _UserService_GetUserByToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserById",
			Handler:       _UserService_GetUserById_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "userService.proto",
}
