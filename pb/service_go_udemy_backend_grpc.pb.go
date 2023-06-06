// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: service_go_udemy_backend.proto

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

// GolangUdemyBackendClient is the client API for GolangUdemyBackend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GolangUdemyBackendClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type golangUdemyBackendClient struct {
	cc grpc.ClientConnInterface
}

func NewGolangUdemyBackendClient(cc grpc.ClientConnInterface) GolangUdemyBackendClient {
	return &golangUdemyBackendClient{cc}
}

func (c *golangUdemyBackendClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.GolangUdemyBackend/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golangUdemyBackendClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/pb.GolangUdemyBackend/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GolangUdemyBackendServer is the server API for GolangUdemyBackend service.
// All implementations must embed UnimplementedGolangUdemyBackendServer
// for forward compatibility
type GolangUdemyBackendServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedGolangUdemyBackendServer()
}

// UnimplementedGolangUdemyBackendServer must be embedded to have forward compatible implementations.
type UnimplementedGolangUdemyBackendServer struct {
}

func (UnimplementedGolangUdemyBackendServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedGolangUdemyBackendServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedGolangUdemyBackendServer) mustEmbedUnimplementedGolangUdemyBackendServer() {}

// UnsafeGolangUdemyBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GolangUdemyBackendServer will
// result in compilation errors.
type UnsafeGolangUdemyBackendServer interface {
	mustEmbedUnimplementedGolangUdemyBackendServer()
}

func RegisterGolangUdemyBackendServer(s grpc.ServiceRegistrar, srv GolangUdemyBackendServer) {
	s.RegisterService(&GolangUdemyBackend_ServiceDesc, srv)
}

func _GolangUdemyBackend_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolangUdemyBackendServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GolangUdemyBackend/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolangUdemyBackendServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GolangUdemyBackend_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolangUdemyBackendServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GolangUdemyBackend/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolangUdemyBackendServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GolangUdemyBackend_ServiceDesc is the grpc.ServiceDesc for GolangUdemyBackend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GolangUdemyBackend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GolangUdemyBackend",
	HandlerType: (*GolangUdemyBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _GolangUdemyBackend_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _GolangUdemyBackend_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_go_udemy_backend.proto",
}