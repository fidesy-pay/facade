// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api/invoices-service/invoices-service.proto

package invoices_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InvoiceStatus int32

const (
	InvoiceStatus_UNKNOWN_STATUS InvoiceStatus = 0
	InvoiceStatus_NEW            InvoiceStatus = 1
	InvoiceStatus_PENDING        InvoiceStatus = 2
	InvoiceStatus_FAILED         InvoiceStatus = 3
	InvoiceStatus_SUCCESS        InvoiceStatus = 4
	InvoiceStatus_EXPIRED        InvoiceStatus = 5
)

// Enum value maps for InvoiceStatus.
var (
	InvoiceStatus_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "NEW",
		2: "PENDING",
		3: "FAILED",
		4: "SUCCESS",
		5: "EXPIRED",
	}
	InvoiceStatus_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"NEW":            1,
		"PENDING":        2,
		"FAILED":         3,
		"SUCCESS":        4,
		"EXPIRED":        5,
	}
)

func (x InvoiceStatus) Enum() *InvoiceStatus {
	p := new(InvoiceStatus)
	*p = x
	return p
}

func (x InvoiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvoiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_invoices_service_invoices_service_proto_enumTypes[0].Descriptor()
}

func (InvoiceStatus) Type() protoreflect.EnumType {
	return &file_api_invoices_service_invoices_service_proto_enumTypes[0]
}

func (x InvoiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvoiceStatus.Descriptor instead.
func (InvoiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{0}
}

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId    string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	UsdAmount   float64                `protobuf:"fixed64,3,opt,name=usd_amount,json=usdAmount,proto3" json:"usd_amount,omitempty"`
	TokenAmount float64                `protobuf:"fixed64,4,opt,name=token_amount,json=tokenAmount,proto3" json:"token_amount,omitempty"`
	Chain       string                 `protobuf:"bytes,5,opt,name=chain,proto3" json:"chain,omitempty"`
	Token       string                 `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	Status      InvoiceStatus          `protobuf:"varint,7,opt,name=status,proto3,enum=invoices_service.InvoiceStatus" json:"status,omitempty"`
	Address     string                 `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Invoice) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Invoice) GetUsdAmount() float64 {
	if x != nil {
		return x.UsdAmount
	}
	return 0
}

func (x *Invoice) GetTokenAmount() float64 {
	if x != nil {
		return x.TokenAmount
	}
	return 0
}

func (x *Invoice) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *Invoice) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Invoice) GetStatus() InvoiceStatus {
	if x != nil {
		return x.Status
	}
	return InvoiceStatus_UNKNOWN_STATUS
}

func (x *Invoice) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Invoice) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string  `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	UsdAmount float64 `protobuf:"fixed64,2,opt,name=usd_amount,json=usdAmount,proto3" json:"usd_amount,omitempty"`
}

func (x *CreateInvoiceRequest) Reset() {
	*x = CreateInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceRequest) ProtoMessage() {}

func (x *CreateInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvoiceRequest.ProtoReflect.Descriptor instead.
func (*CreateInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInvoiceRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CreateInvoiceRequest) GetUsdAmount() float64 {
	if x != nil {
		return x.UsdAmount
	}
	return 0
}

type CreateInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Invoice identifier
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateInvoiceResponse) Reset() {
	*x = CreateInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceResponse) ProtoMessage() {}

func (x *CreateInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvoiceResponse.ProtoReflect.Descriptor instead.
func (*CreateInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateInvoiceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CheckInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Invoice identifier
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckInvoiceRequest) Reset() {
	*x = CheckInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInvoiceRequest) ProtoMessage() {}

func (x *CheckInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInvoiceRequest.ProtoReflect.Descriptor instead.
func (*CheckInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{3}
}

func (x *CheckInvoiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CheckInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoice *Invoice `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
}

func (x *CheckInvoiceResponse) Reset() {
	*x = CheckInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInvoiceResponse) ProtoMessage() {}

func (x *CheckInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInvoiceResponse.ProtoReflect.Descriptor instead.
func (*CheckInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{4}
}

func (x *CheckInvoiceResponse) GetInvoice() *Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

type UpdateInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Chain string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UpdateInvoiceRequest) Reset() {
	*x = UpdateInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInvoiceRequest) ProtoMessage() {}

func (x *UpdateInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInvoiceRequest.ProtoReflect.Descriptor instead.
func (*UpdateInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateInvoiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateInvoiceRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *UpdateInvoiceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoice *Invoice `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
}

func (x *UpdateInvoiceResponse) Reset() {
	*x = UpdateInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInvoiceResponse) ProtoMessage() {}

func (x *UpdateInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInvoiceResponse.ProtoReflect.Descriptor instead.
func (*UpdateInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateInvoiceResponse) GetInvoice() *Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

type ListInvoicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *ListInvoicesRequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListInvoicesRequest) Reset() {
	*x = ListInvoicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvoicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvoicesRequest) ProtoMessage() {}

func (x *ListInvoicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvoicesRequest.ProtoReflect.Descriptor instead.
func (*ListInvoicesRequest) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListInvoicesRequest) GetFilter() *ListInvoicesRequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListInvoicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoices []*Invoice `protobuf:"bytes,1,rep,name=invoices,proto3" json:"invoices,omitempty"`
}

func (x *ListInvoicesResponse) Reset() {
	*x = ListInvoicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvoicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvoicesResponse) ProtoMessage() {}

func (x *ListInvoicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvoicesResponse.ProtoReflect.Descriptor instead.
func (*ListInvoicesResponse) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListInvoicesResponse) GetInvoices() []*Invoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

type ListInvoicesRequest_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdIn            []string        `protobuf:"bytes,1,rep,name=id_in,json=idIn,proto3" json:"id_in,omitempty"`
	ClientIdIn      []string        `protobuf:"bytes,2,rep,name=client_id_in,json=clientIdIn,proto3" json:"client_id_in,omitempty"`
	InvoiceStatusIn []InvoiceStatus `protobuf:"varint,3,rep,packed,name=invoice_status_in,json=invoiceStatusIn,proto3,enum=invoices_service.InvoiceStatus" json:"invoice_status_in,omitempty"`
}

func (x *ListInvoicesRequest_Filter) Reset() {
	*x = ListInvoicesRequest_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_invoices_service_invoices_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvoicesRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvoicesRequest_Filter) ProtoMessage() {}

func (x *ListInvoicesRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_api_invoices_service_invoices_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvoicesRequest_Filter.ProtoReflect.Descriptor instead.
func (*ListInvoicesRequest_Filter) Descriptor() ([]byte, []int) {
	return file_api_invoices_service_invoices_service_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ListInvoicesRequest_Filter) GetIdIn() []string {
	if x != nil {
		return x.IdIn
	}
	return nil
}

func (x *ListInvoicesRequest_Filter) GetClientIdIn() []string {
	if x != nil {
		return x.ClientIdIn
	}
	return nil
}

func (x *ListInvoicesRequest_Filter) GetInvoiceStatusIn() []InvoiceStatus {
	if x != nil {
		return x.InvoiceStatusIn
	}
	return nil
}

var File_api_invoices_service_invoices_service_proto protoreflect.FileDescriptor

var file_api_invoices_service_invoices_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb2, 0x02, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x64,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x75,
	0x73, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x52, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73,
	0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x75, 0x73, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x14, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x52, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4c, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0xea, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x44, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x8c, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x64, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x49, 0x6e, 0x12, 0x4b, 0x0a, 0x11, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x49, 0x6e, 0x22, 0x4d, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x2a, 0x5f, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x05, 0x32, 0x93, 0x03, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0c, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x64, 0x65, 0x73, 0x79,
	0x2d, 0x70, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_invoices_service_invoices_service_proto_rawDescOnce sync.Once
	file_api_invoices_service_invoices_service_proto_rawDescData = file_api_invoices_service_invoices_service_proto_rawDesc
)

func file_api_invoices_service_invoices_service_proto_rawDescGZIP() []byte {
	file_api_invoices_service_invoices_service_proto_rawDescOnce.Do(func() {
		file_api_invoices_service_invoices_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_invoices_service_invoices_service_proto_rawDescData)
	})
	return file_api_invoices_service_invoices_service_proto_rawDescData
}

var file_api_invoices_service_invoices_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_invoices_service_invoices_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_invoices_service_invoices_service_proto_goTypes = []interface{}{
	(InvoiceStatus)(0),                 // 0: invoices_service.InvoiceStatus
	(*Invoice)(nil),                    // 1: invoices_service.Invoice
	(*CreateInvoiceRequest)(nil),       // 2: invoices_service.CreateInvoiceRequest
	(*CreateInvoiceResponse)(nil),      // 3: invoices_service.CreateInvoiceResponse
	(*CheckInvoiceRequest)(nil),        // 4: invoices_service.CheckInvoiceRequest
	(*CheckInvoiceResponse)(nil),       // 5: invoices_service.CheckInvoiceResponse
	(*UpdateInvoiceRequest)(nil),       // 6: invoices_service.UpdateInvoiceRequest
	(*UpdateInvoiceResponse)(nil),      // 7: invoices_service.UpdateInvoiceResponse
	(*ListInvoicesRequest)(nil),        // 8: invoices_service.ListInvoicesRequest
	(*ListInvoicesResponse)(nil),       // 9: invoices_service.ListInvoicesResponse
	(*ListInvoicesRequest_Filter)(nil), // 10: invoices_service.ListInvoicesRequest.Filter
	(*timestamppb.Timestamp)(nil),      // 11: google.protobuf.Timestamp
}
var file_api_invoices_service_invoices_service_proto_depIdxs = []int32{
	0,  // 0: invoices_service.Invoice.status:type_name -> invoices_service.InvoiceStatus
	11, // 1: invoices_service.Invoice.created_at:type_name -> google.protobuf.Timestamp
	1,  // 2: invoices_service.CheckInvoiceResponse.invoice:type_name -> invoices_service.Invoice
	1,  // 3: invoices_service.UpdateInvoiceResponse.invoice:type_name -> invoices_service.Invoice
	10, // 4: invoices_service.ListInvoicesRequest.filter:type_name -> invoices_service.ListInvoicesRequest.Filter
	1,  // 5: invoices_service.ListInvoicesResponse.invoices:type_name -> invoices_service.Invoice
	0,  // 6: invoices_service.ListInvoicesRequest.Filter.invoice_status_in:type_name -> invoices_service.InvoiceStatus
	2,  // 7: invoices_service.InvoicesService.CreateInvoice:input_type -> invoices_service.CreateInvoiceRequest
	4,  // 8: invoices_service.InvoicesService.CheckInvoice:input_type -> invoices_service.CheckInvoiceRequest
	6,  // 9: invoices_service.InvoicesService.UpdateInvoice:input_type -> invoices_service.UpdateInvoiceRequest
	8,  // 10: invoices_service.InvoicesService.ListInvoices:input_type -> invoices_service.ListInvoicesRequest
	3,  // 11: invoices_service.InvoicesService.CreateInvoice:output_type -> invoices_service.CreateInvoiceResponse
	5,  // 12: invoices_service.InvoicesService.CheckInvoice:output_type -> invoices_service.CheckInvoiceResponse
	7,  // 13: invoices_service.InvoicesService.UpdateInvoice:output_type -> invoices_service.UpdateInvoiceResponse
	9,  // 14: invoices_service.InvoicesService.ListInvoices:output_type -> invoices_service.ListInvoicesResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_invoices_service_invoices_service_proto_init() }
func file_api_invoices_service_invoices_service_proto_init() {
	if File_api_invoices_service_invoices_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_invoices_service_invoices_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
		file_api_invoices_service_invoices_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvoiceRequest); i {
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
		file_api_invoices_service_invoices_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvoiceResponse); i {
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
		file_api_invoices_service_invoices_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInvoiceRequest); i {
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
		file_api_invoices_service_invoices_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInvoiceResponse); i {
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
		file_api_invoices_service_invoices_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInvoiceRequest); i {
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
		file_api_invoices_service_invoices_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInvoiceResponse); i {
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
		file_api_invoices_service_invoices_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvoicesRequest); i {
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
		file_api_invoices_service_invoices_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvoicesResponse); i {
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
		file_api_invoices_service_invoices_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvoicesRequest_Filter); i {
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
			RawDescriptor: file_api_invoices_service_invoices_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_invoices_service_invoices_service_proto_goTypes,
		DependencyIndexes: file_api_invoices_service_invoices_service_proto_depIdxs,
		EnumInfos:         file_api_invoices_service_invoices_service_proto_enumTypes,
		MessageInfos:      file_api_invoices_service_invoices_service_proto_msgTypes,
	}.Build()
	File_api_invoices_service_invoices_service_proto = out.File
	file_api_invoices_service_invoices_service_proto_rawDesc = nil
	file_api_invoices_service_invoices_service_proto_goTypes = nil
	file_api_invoices_service_invoices_service_proto_depIdxs = nil
}
