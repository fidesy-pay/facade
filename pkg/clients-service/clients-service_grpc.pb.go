// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: api/clients-service/clients-service.proto

package clients_service

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

const (
	ClientsService_CreateClient_FullMethodName = "/clients_service.ClientsService/CreateClient"
	ClientsService_GetClient_FullMethodName    = "/clients_service.ClientsService/GetClient"
)

// ClientsServiceClient is the client API for ClientsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientsServiceClient interface {
	CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error)
	GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error)
}

type clientsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientsServiceClient(cc grpc.ClientConnInterface) ClientsServiceClient {
	return &clientsServiceClient{cc}
}

func (c *clientsServiceClient) CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, ClientsService_CreateClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServiceClient) GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, ClientsService_GetClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientsServiceServer is the server API for ClientsService service.
// All implementations must embed UnimplementedClientsServiceServer
// for forward compatibility
type ClientsServiceServer interface {
	CreateClient(context.Context, *CreateClientRequest) (*Client, error)
	GetClient(context.Context, *GetClientRequest) (*Client, error)
	mustEmbedUnimplementedClientsServiceServer()
}

// UnimplementedClientsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientsServiceServer struct {
}

func (UnimplementedClientsServiceServer) CreateClient(context.Context, *CreateClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (UnimplementedClientsServiceServer) GetClient(context.Context, *GetClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedClientsServiceServer) mustEmbedUnimplementedClientsServiceServer() {}

// UnsafeClientsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientsServiceServer will
// result in compilation errors.
type UnsafeClientsServiceServer interface {
	mustEmbedUnimplementedClientsServiceServer()
}

func RegisterClientsServiceServer(s grpc.ServiceRegistrar, srv ClientsServiceServer) {
	s.RegisterService(&ClientsService_ServiceDesc, srv)
}

func _ClientsService_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServiceServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientsService_CreateClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServiceServer).CreateClient(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsService_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServiceServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientsService_GetClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServiceServer).GetClient(ctx, req.(*GetClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientsService_ServiceDesc is the grpc.ServiceDesc for ClientsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clients_service.ClientsService",
	HandlerType: (*ClientsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _ClientsService_CreateClient_Handler,
		},
		{
			MethodName: "GetClient",
			Handler:    _ClientsService_GetClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/clients-service/clients-service.proto",
}