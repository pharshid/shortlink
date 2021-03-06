// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shortcut

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	api "shortlink/api"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShortcutServiceClient is the client API for ShortcutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortcutServiceClient interface {
	CreateShortcut(ctx context.Context, in *InternalCreateShortcutRequest, opts ...grpc.CallOption) (*api.CreateShortcutResponse, error)
	DeleteShortcut(ctx context.Context, in *InternalDeleteShortcutRequest, opts ...grpc.CallOption) (*api.DeleteShortcutResponse, error)
	GetShortcut(ctx context.Context, in *api.GetShortcutRequest, opts ...grpc.CallOption) (*api.GetShortcutResponse, error)
}

type shortcutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortcutServiceClient(cc grpc.ClientConnInterface) ShortcutServiceClient {
	return &shortcutServiceClient{cc}
}

func (c *shortcutServiceClient) CreateShortcut(ctx context.Context, in *InternalCreateShortcutRequest, opts ...grpc.CallOption) (*api.CreateShortcutResponse, error) {
	out := new(api.CreateShortcutResponse)
	err := c.cc.Invoke(ctx, "/protos.ShortcutService/CreateShortcut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortcutServiceClient) DeleteShortcut(ctx context.Context, in *InternalDeleteShortcutRequest, opts ...grpc.CallOption) (*api.DeleteShortcutResponse, error) {
	out := new(api.DeleteShortcutResponse)
	err := c.cc.Invoke(ctx, "/protos.ShortcutService/DeleteShortcut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortcutServiceClient) GetShortcut(ctx context.Context, in *api.GetShortcutRequest, opts ...grpc.CallOption) (*api.GetShortcutResponse, error) {
	out := new(api.GetShortcutResponse)
	err := c.cc.Invoke(ctx, "/protos.ShortcutService/GetShortcut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortcutServiceServer is the server API for ShortcutService service.
// All implementations must embed UnimplementedShortcutServiceServer
// for forward compatibility
type ShortcutServiceServer interface {
	CreateShortcut(context.Context, *InternalCreateShortcutRequest) (*api.CreateShortcutResponse, error)
	DeleteShortcut(context.Context, *InternalDeleteShortcutRequest) (*api.DeleteShortcutResponse, error)
	GetShortcut(context.Context, *api.GetShortcutRequest) (*api.GetShortcutResponse, error)
	mustEmbedUnimplementedShortcutServiceServer()
}

// UnimplementedShortcutServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShortcutServiceServer struct {
}

func (UnimplementedShortcutServiceServer) CreateShortcut(context.Context, *InternalCreateShortcutRequest) (*api.CreateShortcutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShortcut not implemented")
}
func (UnimplementedShortcutServiceServer) DeleteShortcut(context.Context, *InternalDeleteShortcutRequest) (*api.DeleteShortcutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShortcut not implemented")
}
func (UnimplementedShortcutServiceServer) GetShortcut(context.Context, *api.GetShortcutRequest) (*api.GetShortcutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortcut not implemented")
}
func (UnimplementedShortcutServiceServer) mustEmbedUnimplementedShortcutServiceServer() {}

// UnsafeShortcutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortcutServiceServer will
// result in compilation errors.
type UnsafeShortcutServiceServer interface {
	mustEmbedUnimplementedShortcutServiceServer()
}

func RegisterShortcutServiceServer(s grpc.ServiceRegistrar, srv ShortcutServiceServer) {
	s.RegisterService(&ShortcutService_ServiceDesc, srv)
}

func _ShortcutService_CreateShortcut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalCreateShortcutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortcutServiceServer).CreateShortcut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ShortcutService/CreateShortcut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortcutServiceServer).CreateShortcut(ctx, req.(*InternalCreateShortcutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortcutService_DeleteShortcut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalDeleteShortcutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortcutServiceServer).DeleteShortcut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ShortcutService/DeleteShortcut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortcutServiceServer).DeleteShortcut(ctx, req.(*InternalDeleteShortcutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortcutService_GetShortcut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetShortcutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortcutServiceServer).GetShortcut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ShortcutService/GetShortcut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortcutServiceServer).GetShortcut(ctx, req.(*api.GetShortcutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortcutService_ServiceDesc is the grpc.ServiceDesc for ShortcutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortcutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ShortcutService",
	HandlerType: (*ShortcutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShortcut",
			Handler:    _ShortcutService_CreateShortcut_Handler,
		},
		{
			MethodName: "DeleteShortcut",
			Handler:    _ShortcutService_DeleteShortcut_Handler,
		},
		{
			MethodName: "GetShortcut",
			Handler:    _ShortcutService_GetShortcut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shortcut.proto",
}
