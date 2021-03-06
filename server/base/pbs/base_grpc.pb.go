// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbs

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

// BaseServiceClient is the client API for BaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseServiceClient interface {
	// A simple unary RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	Ping(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
}

type baseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseServiceClient(cc grpc.ClientConnInterface) BaseServiceClient {
	return &baseServiceClient{cc}
}

func (c *baseServiceClient) Ping(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/base.BaseService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseServiceServer is the server API for BaseService service.
// All implementations must embed UnimplementedBaseServiceServer
// for forward compatibility
type BaseServiceServer interface {
	// A simple unary RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	Ping(context.Context, *Input) (*Output, error)
	mustEmbedUnimplementedBaseServiceServer()
}

// UnimplementedBaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBaseServiceServer struct {
}

func (UnimplementedBaseServiceServer) Ping(context.Context, *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBaseServiceServer) mustEmbedUnimplementedBaseServiceServer() {}

// UnsafeBaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseServiceServer will
// result in compilation errors.
type UnsafeBaseServiceServer interface {
	mustEmbedUnimplementedBaseServiceServer()
}

func RegisterBaseServiceServer(s grpc.ServiceRegistrar, srv BaseServiceServer) {
	s.RegisterService(&BaseService_ServiceDesc, srv)
}

func _BaseService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.BaseService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).Ping(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

// BaseService_ServiceDesc is the grpc.ServiceDesc for BaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "base.BaseService",
	HandlerType: (*BaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BaseService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base.proto",
}
