// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: checkout.proto

package checkout

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PurchaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *PurchaseResponse) Reset() {
	*x = PurchaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseResponse) ProtoMessage() {}

func (x *PurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseResponse.ProtoReflect.Descriptor instead.
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{0}
}

func (x *PurchaseResponse) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type AddToCartParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64  `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Sku   uint32 `protobuf:"varint,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AddToCartParams) Reset() {
	*x = AddToCartParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartParams) ProtoMessage() {}

func (x *AddToCartParams) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartParams.ProtoReflect.Descriptor instead.
func (*AddToCartParams) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{1}
}

func (x *AddToCartParams) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *AddToCartParams) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *AddToCartParams) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeleteFromCartParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64  `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Sku   uint32 `protobuf:"varint,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *DeleteFromCartParams) Reset() {
	*x = DeleteFromCartParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFromCartParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFromCartParams) ProtoMessage() {}

func (x *DeleteFromCartParams) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFromCartParams.ProtoReflect.Descriptor instead.
func (*DeleteFromCartParams) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteFromCartParams) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *DeleteFromCartParams) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *DeleteFromCartParams) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListCartParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User int64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ListCartParams) Reset() {
	*x = ListCartParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartParams) ProtoMessage() {}

func (x *ListCartParams) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartParams.ProtoReflect.Descriptor instead.
func (*ListCartParams) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{3}
}

func (x *ListCartParams) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

type ListCartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price uint32 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ListCartItem) Reset() {
	*x = ListCartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartItem) ProtoMessage() {}

func (x *ListCartItem) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartItem.ProtoReflect.Descriptor instead.
func (*ListCartItem) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{4}
}

func (x *ListCartItem) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *ListCartItem) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListCartItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListCartItem) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ListCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*ListCartItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice uint32          `protobuf:"varint,2,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
}

func (x *ListCartResponse) Reset() {
	*x = ListCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartResponse) ProtoMessage() {}

func (x *ListCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartResponse.ProtoReflect.Descriptor instead.
func (*ListCartResponse) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{5}
}

func (x *ListCartResponse) GetItems() []*ListCartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListCartResponse) GetTotalPrice() uint32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type PurchaseParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User int64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *PurchaseParams) Reset() {
	*x = PurchaseParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseParams) ProtoMessage() {}

func (x *PurchaseParams) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseParams.ProtoReflect.Descriptor instead.
func (*PurchaseParams) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{6}
}

func (x *PurchaseParams) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

var File_checkout_proto protoreflect.FileDescriptor

var file_checkout_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61,
	0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73,
	0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x60,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x6b, 0x75,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x60, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xa0, 0x02, 0x0a, 0x08, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x40, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43,
	0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x18, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x08, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1a, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkout_proto_rawDescOnce sync.Once
	file_checkout_proto_rawDescData = file_checkout_proto_rawDesc
)

func file_checkout_proto_rawDescGZIP() []byte {
	file_checkout_proto_rawDescOnce.Do(func() {
		file_checkout_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkout_proto_rawDescData)
	})
	return file_checkout_proto_rawDescData
}

var file_checkout_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_checkout_proto_goTypes = []interface{}{
	(*PurchaseResponse)(nil),     // 0: checkout.PurchaseResponse
	(*AddToCartParams)(nil),      // 1: checkout.AddToCartParams
	(*DeleteFromCartParams)(nil), // 2: checkout.DeleteFromCartParams
	(*ListCartParams)(nil),       // 3: checkout.ListCartParams
	(*ListCartItem)(nil),         // 4: checkout.ListCartItem
	(*ListCartResponse)(nil),     // 5: checkout.ListCartResponse
	(*PurchaseParams)(nil),       // 6: checkout.PurchaseParams
	(*emptypb.Empty)(nil),        // 7: google.protobuf.Empty
}
var file_checkout_proto_depIdxs = []int32{
	4, // 0: checkout.ListCartResponse.items:type_name -> checkout.ListCartItem
	1, // 1: checkout.Checkout.AddToCart:input_type -> checkout.AddToCartParams
	2, // 2: checkout.Checkout.DeleteFromCart:input_type -> checkout.DeleteFromCartParams
	3, // 3: checkout.Checkout.ListCart:input_type -> checkout.ListCartParams
	6, // 4: checkout.Checkout.Purchase:input_type -> checkout.PurchaseParams
	7, // 5: checkout.Checkout.AddToCart:output_type -> google.protobuf.Empty
	7, // 6: checkout.Checkout.DeleteFromCart:output_type -> google.protobuf.Empty
	5, // 7: checkout.Checkout.ListCart:output_type -> checkout.ListCartResponse
	0, // 8: checkout.Checkout.Purchase:output_type -> checkout.PurchaseResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_checkout_proto_init() }
func file_checkout_proto_init() {
	if File_checkout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseResponse); i {
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
		file_checkout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartParams); i {
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
		file_checkout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFromCartParams); i {
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
		file_checkout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartParams); i {
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
		file_checkout_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartItem); i {
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
		file_checkout_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartResponse); i {
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
		file_checkout_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseParams); i {
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
			RawDescriptor: file_checkout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkout_proto_goTypes,
		DependencyIndexes: file_checkout_proto_depIdxs,
		MessageInfos:      file_checkout_proto_msgTypes,
	}.Build()
	File_checkout_proto = out.File
	file_checkout_proto_rawDesc = nil
	file_checkout_proto_goTypes = nil
	file_checkout_proto_depIdxs = nil
}
