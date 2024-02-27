// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: box_info.proto

package boxinfo

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

// CreateBid
type CreateBidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBidRequest) Reset() {
	*x = CreateBidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBidRequest) ProtoMessage() {}

func (x *CreateBidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBidRequest.ProtoReflect.Descriptor instead.
func (*CreateBidRequest) Descriptor() ([]byte, []int) {
	return file_box_info_proto_rawDescGZIP(), []int{0}
}

type CreateBidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid int64 `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
}

func (x *CreateBidResponse) Reset() {
	*x = CreateBidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBidResponse) ProtoMessage() {}

func (x *CreateBidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBidResponse.ProtoReflect.Descriptor instead.
func (*CreateBidResponse) Descriptor() ([]byte, []int) {
	return file_box_info_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBidResponse) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

// CreateBox
type CreateBoxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid          int64  `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Introduction string `protobuf:"bytes,3,opt,name=Introduction,proto3" json:"Introduction,omitempty"`
}

func (x *CreateBoxRequest) Reset() {
	*x = CreateBoxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoxRequest) ProtoMessage() {}

func (x *CreateBoxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoxRequest.ProtoReflect.Descriptor instead.
func (*CreateBoxRequest) Descriptor() ([]byte, []int) {
	return file_box_info_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBoxRequest) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *CreateBoxRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBoxRequest) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

type CreateBoxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBoxResponse) Reset() {
	*x = CreateBoxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoxResponse) ProtoMessage() {}

func (x *CreateBoxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoxResponse.ProtoReflect.Descriptor instead.
func (*CreateBoxResponse) Descriptor() ([]byte, []int) {
	return file_box_info_proto_rawDescGZIP(), []int{3}
}

// UpdateBoxInfo
type UpdateBoxInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid          int64  `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Introduction string `protobuf:"bytes,3,opt,name=Introduction,proto3" json:"Introduction,omitempty"`
}

func (x *UpdateBoxInfoRequest) Reset() {
	*x = UpdateBoxInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoxInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoxInfoRequest) ProtoMessage() {}

func (x *UpdateBoxInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoxInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateBoxInfoRequest) Descriptor() ([]byte, []int) {
	return file_box_info_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBoxInfoRequest) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *UpdateBoxInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBoxInfoRequest) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

type UpdateBoxInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBoxInfoResponse) Reset() {
	*x = UpdateBoxInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoxInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoxInfoResponse) ProtoMessage() {}

func (x *UpdateBoxInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoxInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateBoxInfoResponse) Descriptor() ([]byte, []int) {
	return file_box_info_proto_rawDescGZIP(), []int{5}
}

// GetBoxInfo
type GetBoxInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid int64 `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
}

func (x *GetBoxInfoRequest) Reset() {
	*x = GetBoxInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_info_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoxInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoxInfoRequest) ProtoMessage() {}

func (x *GetBoxInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_info_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoxInfoRequest.ProtoReflect.Descriptor instead.
func (*GetBoxInfoRequest) Descriptor() ([]byte, []int) {
	return file_box_info_proto_rawDescGZIP(), []int{6}
}

func (x *GetBoxInfoRequest) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

type GetBoxInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Introduction string `protobuf:"bytes,2,opt,name=Introduction,proto3" json:"Introduction,omitempty"`
}

func (x *GetBoxInfoResponse) Reset() {
	*x = GetBoxInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_info_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoxInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoxInfoResponse) ProtoMessage() {}

func (x *GetBoxInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_info_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoxInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBoxInfoResponse) Descriptor() ([]byte, []int) {
	return file_box_info_proto_rawDescGZIP(), []int{7}
}

func (x *GetBoxInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBoxInfoResponse) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

var File_box_info_proto protoreflect.FileDescriptor

var file_box_info_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x62, 0x6f, 0x78, 0x69, 0x6e, 0x66, 0x6f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x42, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x42, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x6e, 0x74, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x42, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49,
	0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x42, 0x69, 0x64, 0x22,
	0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x74,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xae, 0x03,
	0x0a, 0x07, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4e, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x12, 0x1f, 0x2e, 0x62, 0x6f, 0x78, 0x69, 0x6e, 0x66, 0x6f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x6f, 0x78, 0x69, 0x6e, 0x66,
	0x6f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x12, 0x1f, 0x2e, 0x62, 0x6f, 0x78, 0x69, 0x6e, 0x66, 0x6f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x6f, 0x78, 0x69, 0x6e, 0x66,
	0x6f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x12, 0x1f, 0x2e, 0x62,
	0x6f, 0x78, 0x69, 0x6e, 0x66, 0x6f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x62, 0x6f, 0x78, 0x69, 0x6e, 0x66, 0x6f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x23, 0x2e, 0x62, 0x6f, 0x78, 0x69, 0x6e, 0x66, 0x6f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x6f, 0x78, 0x69, 0x6e, 0x66, 0x6f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x78, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x78, 0x69,
	0x6e, 0x66, 0x6f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x78,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6f,
	0x78, 0x69, 0x6e, 0x66, 0x6f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x62, 0x6f, 0x78, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_box_info_proto_rawDescOnce sync.Once
	file_box_info_proto_rawDescData = file_box_info_proto_rawDesc
)

func file_box_info_proto_rawDescGZIP() []byte {
	file_box_info_proto_rawDescOnce.Do(func() {
		file_box_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_box_info_proto_rawDescData)
	})
	return file_box_info_proto_rawDescData
}

var file_box_info_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_box_info_proto_goTypes = []interface{}{
	(*CreateBidRequest)(nil),      // 0: boxinfoclient.CreateBidRequest
	(*CreateBidResponse)(nil),     // 1: boxinfoclient.CreateBidResponse
	(*CreateBoxRequest)(nil),      // 2: boxinfoclient.CreateBoxRequest
	(*CreateBoxResponse)(nil),     // 3: boxinfoclient.CreateBoxResponse
	(*UpdateBoxInfoRequest)(nil),  // 4: boxinfoclient.UpdateBoxInfoRequest
	(*UpdateBoxInfoResponse)(nil), // 5: boxinfoclient.UpdateBoxInfoResponse
	(*GetBoxInfoRequest)(nil),     // 6: boxinfoclient.GetBoxInfoRequest
	(*GetBoxInfoResponse)(nil),    // 7: boxinfoclient.GetBoxInfoResponse
}
var file_box_info_proto_depIdxs = []int32{
	0, // 0: boxinfoclient.BoxInfo.CreateBid:input_type -> boxinfoclient.CreateBidRequest
	2, // 1: boxinfoclient.BoxInfo.CreateBox:input_type -> boxinfoclient.CreateBoxRequest
	2, // 2: boxinfoclient.BoxInfo.CreateBoxRevert:input_type -> boxinfoclient.CreateBoxRequest
	4, // 3: boxinfoclient.BoxInfo.UpdateBoxInfo:input_type -> boxinfoclient.UpdateBoxInfoRequest
	6, // 4: boxinfoclient.BoxInfo.GetBoxInfo:input_type -> boxinfoclient.GetBoxInfoRequest
	1, // 5: boxinfoclient.BoxInfo.CreateBid:output_type -> boxinfoclient.CreateBidResponse
	3, // 6: boxinfoclient.BoxInfo.CreateBox:output_type -> boxinfoclient.CreateBoxResponse
	3, // 7: boxinfoclient.BoxInfo.CreateBoxRevert:output_type -> boxinfoclient.CreateBoxResponse
	5, // 8: boxinfoclient.BoxInfo.UpdateBoxInfo:output_type -> boxinfoclient.UpdateBoxInfoResponse
	7, // 9: boxinfoclient.BoxInfo.GetBoxInfo:output_type -> boxinfoclient.GetBoxInfoResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_box_info_proto_init() }
func file_box_info_proto_init() {
	if File_box_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_box_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBidRequest); i {
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
		file_box_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBidResponse); i {
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
		file_box_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoxRequest); i {
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
		file_box_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoxResponse); i {
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
		file_box_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoxInfoRequest); i {
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
		file_box_info_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoxInfoResponse); i {
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
		file_box_info_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBoxInfoRequest); i {
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
		file_box_info_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBoxInfoResponse); i {
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
			RawDescriptor: file_box_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_box_info_proto_goTypes,
		DependencyIndexes: file_box_info_proto_depIdxs,
		MessageInfos:      file_box_info_proto_msgTypes,
	}.Build()
	File_box_info_proto = out.File
	file_box_info_proto_rawDesc = nil
	file_box_info_proto_goTypes = nil
	file_box_info_proto_depIdxs = nil
}
