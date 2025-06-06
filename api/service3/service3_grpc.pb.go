// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: api/service3/service3.proto

package service3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Service3_ScheduleScript_FullMethodName = "/Service3/ScheduleScript"
)

// Service3Client is the client API for Service3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Service3Client interface {
	ScheduleScript(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
}

type service3Client struct {
	cc grpc.ClientConnInterface
}

func NewService3Client(cc grpc.ClientConnInterface) Service3Client {
	return &service3Client{cc}
}

func (c *service3Client) ScheduleScript(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, Service3_ScheduleScript_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service3Server is the server API for Service3 service.
// All implementations must embed UnimplementedService3Server
// for forward compatibility.
type Service3Server interface {
	ScheduleScript(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	mustEmbedUnimplementedService3Server()
}

// UnimplementedService3Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedService3Server struct{}

func (UnimplementedService3Server) ScheduleScript(context.Context, *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleScript not implemented")
}
func (UnimplementedService3Server) mustEmbedUnimplementedService3Server() {}
func (UnimplementedService3Server) testEmbeddedByValue()                  {}

// UnsafeService3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Service3Server will
// result in compilation errors.
type UnsafeService3Server interface {
	mustEmbedUnimplementedService3Server()
}

func RegisterService3Server(s grpc.ServiceRegistrar, srv Service3Server) {
	// If the following call pancis, it indicates UnimplementedService3Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Service3_ServiceDesc, srv)
}

func _Service3_ScheduleScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service3Server).ScheduleScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service3_ScheduleScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service3Server).ScheduleScript(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service3_ServiceDesc is the grpc.ServiceDesc for Service3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Service3",
	HandlerType: (*Service3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleScript",
			Handler:    _Service3_ScheduleScript_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service3/service3.proto",
}
