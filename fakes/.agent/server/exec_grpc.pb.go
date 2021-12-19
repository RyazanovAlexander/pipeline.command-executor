// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package server

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

// ExecServiceClient is the client API for ExecService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecServiceClient interface {
	ExecuteCommands(ctx context.Context, in *ExecCommands, opts ...grpc.CallOption) (*ExecResult, error)
}

type execServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExecServiceClient(cc grpc.ClientConnInterface) ExecServiceClient {
	return &execServiceClient{cc}
}

func (c *execServiceClient) ExecuteCommands(ctx context.Context, in *ExecCommands, opts ...grpc.CallOption) (*ExecResult, error) {
	out := new(ExecResult)
	err := c.cc.Invoke(ctx, "/exec.ExecService/ExecuteCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecServiceServer is the server API for ExecService service.
// All implementations must embed UnimplementedExecServiceServer
// for forward compatibility
type ExecServiceServer interface {
	ExecuteCommands(context.Context, *ExecCommands) (*ExecResult, error)
	mustEmbedUnimplementedExecServiceServer()
}

// UnimplementedExecServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExecServiceServer struct {
}

func (UnimplementedExecServiceServer) ExecuteCommands(context.Context, *ExecCommands) (*ExecResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCommands not implemented")
}
func (UnimplementedExecServiceServer) mustEmbedUnimplementedExecServiceServer() {}

// UnsafeExecServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecServiceServer will
// result in compilation errors.
type UnsafeExecServiceServer interface {
	mustEmbedUnimplementedExecServiceServer()
}

func RegisterExecServiceServer(s grpc.ServiceRegistrar, srv ExecServiceServer) {
	s.RegisterService(&ExecService_ServiceDesc, srv)
}

func _ExecService_ExecuteCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecCommands)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecServiceServer).ExecuteCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exec.ExecService/ExecuteCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecServiceServer).ExecuteCommands(ctx, req.(*ExecCommands))
	}
	return interceptor(ctx, in, info, handler)
}

// ExecService_ServiceDesc is the grpc.ServiceDesc for ExecService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExecService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exec.ExecService",
	HandlerType: (*ExecServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteCommands",
			Handler:    _ExecService_ExecuteCommands_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exec.proto",
}
