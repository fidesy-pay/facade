// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: api/crypto-service/crypto-service.proto

package crypto_service

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
	CryptoService_AcceptCrypto_FullMethodName = "/crypto_service.CryptoService/AcceptCrypto"
	CryptoService_Transfer_FullMethodName     = "/crypto_service.CryptoService/Transfer"
	CryptoService_CreateWallet_FullMethodName = "/crypto_service.CryptoService/CreateWallet"
	CryptoService_ListWallets_FullMethodName  = "/crypto_service.CryptoService/ListWallets"
)

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoServiceClient interface {
	AcceptCrypto(ctx context.Context, in *AcceptCryptoRequest, opts ...grpc.CallOption) (*AcceptCryptoResponse, error)
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*Wallet, error)
	ListWallets(ctx context.Context, in *ListWalletsRequest, opts ...grpc.CallOption) (*ListWalletsResponse, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) AcceptCrypto(ctx context.Context, in *AcceptCryptoRequest, opts ...grpc.CallOption) (*AcceptCryptoResponse, error) {
	out := new(AcceptCryptoResponse)
	err := c.cc.Invoke(ctx, CryptoService_AcceptCrypto_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, CryptoService_Transfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, CryptoService_CreateWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) ListWallets(ctx context.Context, in *ListWalletsRequest, opts ...grpc.CallOption) (*ListWalletsResponse, error) {
	out := new(ListWalletsResponse)
	err := c.cc.Invoke(ctx, CryptoService_ListWallets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
// All implementations must embed UnimplementedCryptoServiceServer
// for forward compatibility
type CryptoServiceServer interface {
	AcceptCrypto(context.Context, *AcceptCryptoRequest) (*AcceptCryptoResponse, error)
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*Wallet, error)
	ListWallets(context.Context, *ListWalletsRequest) (*ListWalletsResponse, error)
	mustEmbedUnimplementedCryptoServiceServer()
}

// UnimplementedCryptoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (UnimplementedCryptoServiceServer) AcceptCrypto(context.Context, *AcceptCryptoRequest) (*AcceptCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptCrypto not implemented")
}
func (UnimplementedCryptoServiceServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedCryptoServiceServer) CreateWallet(context.Context, *CreateWalletRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedCryptoServiceServer) ListWallets(context.Context, *ListWalletsRequest) (*ListWalletsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWallets not implemented")
}
func (UnimplementedCryptoServiceServer) mustEmbedUnimplementedCryptoServiceServer() {}

// UnsafeCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoServiceServer will
// result in compilation errors.
type UnsafeCryptoServiceServer interface {
	mustEmbedUnimplementedCryptoServiceServer()
}

func RegisterCryptoServiceServer(s grpc.ServiceRegistrar, srv CryptoServiceServer) {
	s.RegisterService(&CryptoService_ServiceDesc, srv)
}

func _CryptoService_AcceptCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).AcceptCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_AcceptCrypto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).AcceptCrypto(ctx, req.(*AcceptCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_ListWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).ListWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_ListWallets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).ListWallets(ctx, req.(*ListWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptoService_ServiceDesc is the grpc.ServiceDesc for CryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crypto_service.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptCrypto",
			Handler:    _CryptoService_AcceptCrypto_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _CryptoService_Transfer_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _CryptoService_CreateWallet_Handler,
		},
		{
			MethodName: "ListWallets",
			Handler:    _CryptoService_ListWallets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/crypto-service/crypto-service.proto",
}
