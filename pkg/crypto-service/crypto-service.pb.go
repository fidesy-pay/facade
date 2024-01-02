// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api/crypto-service/crypto-service.proto

package crypto_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AcceptCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain     string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	InvoiceId string `protobuf:"bytes,3,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
}

func (x *AcceptCryptoRequest) Reset() {
	*x = AcceptCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptCryptoRequest) ProtoMessage() {}

func (x *AcceptCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptCryptoRequest.ProtoReflect.Descriptor instead.
func (*AcceptCryptoRequest) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{0}
}

func (x *AcceptCryptoRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *AcceptCryptoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AcceptCryptoRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

type AcceptCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AcceptCryptoResponse) Reset() {
	*x = AcceptCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptCryptoResponse) ProtoMessage() {}

func (x *AcceptCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptCryptoResponse.ProtoReflect.Descriptor instead.
func (*AcceptCryptoResponse) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{1}
}

func (x *AcceptCryptoResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	InvoiceId string `protobuf:"bytes,2,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	Chain     string `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Token     string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{2}
}

func (x *TransferRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *TransferRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *TransferRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *TransferRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash string `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{3}
}

func (x *TransferResponse) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance float64 `protobuf:"fixed64,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{4}
}

func (x *Wallet) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Wallet) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type CreateWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *CreateWalletRequest) Reset() {
	*x = CreateWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRequest) ProtoMessage() {}

func (x *CreateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRequest.ProtoReflect.Descriptor instead.
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateWalletRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ListWalletsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *ListWalletsRequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListWalletsRequest) Reset() {
	*x = ListWalletsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWalletsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWalletsRequest) ProtoMessage() {}

func (x *ListWalletsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWalletsRequest.ProtoReflect.Descriptor instead.
func (*ListWalletsRequest) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListWalletsRequest) GetFilter() *ListWalletsRequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListWalletsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallets []*Wallet `protobuf:"bytes,1,rep,name=wallets,proto3" json:"wallets,omitempty"`
}

func (x *ListWalletsResponse) Reset() {
	*x = ListWalletsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWalletsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWalletsResponse) ProtoMessage() {}

func (x *ListWalletsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWalletsResponse.ProtoReflect.Descriptor instead.
func (*ListWalletsResponse) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListWalletsResponse) GetWallets() []*Wallet {
	if x != nil {
		return x.Wallets
	}
	return nil
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Chain   string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetBalanceRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *GetBalanceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance float64 `protobuf:"fixed64,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetBalanceResponse) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type ListWalletsRequest_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIdIn []string `protobuf:"bytes,1,rep,name=client_id_in,json=clientIdIn,proto3" json:"client_id_in,omitempty"`
}

func (x *ListWalletsRequest_Filter) Reset() {
	*x = ListWalletsRequest_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_service_crypto_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWalletsRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWalletsRequest_Filter) ProtoMessage() {}

func (x *ListWalletsRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_service_crypto_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWalletsRequest_Filter.ProtoReflect.Descriptor instead.
func (*ListWalletsRequest_Filter) Descriptor() ([]byte, []int) {
	return file_api_crypto_service_crypto_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ListWalletsRequest_Filter) GetClientIdIn() []string {
	if x != nil {
		return x.ClientIdIn
	}
	return nil
}

var File_api_crypto_service_crypto_service_proto protoreflect.FileDescriptor

var file_api_crypto_service_crypto_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x60, 0x0a, 0x13, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x79, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3c, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x32, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x41, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x1a, 0x2a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x49, 0x6e, 0x22,
	0x47, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x07, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x22, 0x59, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x32, 0xb3, 0x03, 0x0a, 0x0d, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x23, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12,
	0x23, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x56, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x64, 0x65, 0x73, 0x79, 0x2d, 0x70,
	0x61, 0x79, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_crypto_service_crypto_service_proto_rawDescOnce sync.Once
	file_api_crypto_service_crypto_service_proto_rawDescData = file_api_crypto_service_crypto_service_proto_rawDesc
)

func file_api_crypto_service_crypto_service_proto_rawDescGZIP() []byte {
	file_api_crypto_service_crypto_service_proto_rawDescOnce.Do(func() {
		file_api_crypto_service_crypto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_crypto_service_crypto_service_proto_rawDescData)
	})
	return file_api_crypto_service_crypto_service_proto_rawDescData
}

var file_api_crypto_service_crypto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_crypto_service_crypto_service_proto_goTypes = []interface{}{
	(*AcceptCryptoRequest)(nil),       // 0: crypto_service.AcceptCryptoRequest
	(*AcceptCryptoResponse)(nil),      // 1: crypto_service.AcceptCryptoResponse
	(*TransferRequest)(nil),           // 2: crypto_service.TransferRequest
	(*TransferResponse)(nil),          // 3: crypto_service.TransferResponse
	(*Wallet)(nil),                    // 4: crypto_service.Wallet
	(*CreateWalletRequest)(nil),       // 5: crypto_service.CreateWalletRequest
	(*ListWalletsRequest)(nil),        // 6: crypto_service.ListWalletsRequest
	(*ListWalletsResponse)(nil),       // 7: crypto_service.ListWalletsResponse
	(*GetBalanceRequest)(nil),         // 8: crypto_service.GetBalanceRequest
	(*GetBalanceResponse)(nil),        // 9: crypto_service.GetBalanceResponse
	(*ListWalletsRequest_Filter)(nil), // 10: crypto_service.ListWalletsRequest.Filter
}
var file_api_crypto_service_crypto_service_proto_depIdxs = []int32{
	10, // 0: crypto_service.ListWalletsRequest.filter:type_name -> crypto_service.ListWalletsRequest.Filter
	4,  // 1: crypto_service.ListWalletsResponse.wallets:type_name -> crypto_service.Wallet
	0,  // 2: crypto_service.CryptoService.AcceptCrypto:input_type -> crypto_service.AcceptCryptoRequest
	2,  // 3: crypto_service.CryptoService.Transfer:input_type -> crypto_service.TransferRequest
	5,  // 4: crypto_service.CryptoService.CreateWallet:input_type -> crypto_service.CreateWalletRequest
	6,  // 5: crypto_service.CryptoService.ListWallets:input_type -> crypto_service.ListWalletsRequest
	8,  // 6: crypto_service.CryptoService.GetBalance:input_type -> crypto_service.GetBalanceRequest
	1,  // 7: crypto_service.CryptoService.AcceptCrypto:output_type -> crypto_service.AcceptCryptoResponse
	3,  // 8: crypto_service.CryptoService.Transfer:output_type -> crypto_service.TransferResponse
	4,  // 9: crypto_service.CryptoService.CreateWallet:output_type -> crypto_service.Wallet
	7,  // 10: crypto_service.CryptoService.ListWallets:output_type -> crypto_service.ListWalletsResponse
	9,  // 11: crypto_service.CryptoService.GetBalance:output_type -> crypto_service.GetBalanceResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_crypto_service_crypto_service_proto_init() }
func file_api_crypto_service_crypto_service_proto_init() {
	if File_api_crypto_service_crypto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_crypto_service_crypto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptCryptoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptCryptoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWalletsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWalletsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_service_crypto_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWalletsRequest_Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_crypto_service_crypto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_crypto_service_crypto_service_proto_goTypes,
		DependencyIndexes: file_api_crypto_service_crypto_service_proto_depIdxs,
		MessageInfos:      file_api_crypto_service_crypto_service_proto_msgTypes,
	}.Build()
	File_api_crypto_service_crypto_service_proto = out.File
	file_api_crypto_service_crypto_service_proto_rawDesc = nil
	file_api_crypto_service_crypto_service_proto_goTypes = nil
	file_api_crypto_service_crypto_service_proto_depIdxs = nil
}
