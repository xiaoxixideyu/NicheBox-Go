// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: user.proto

package user

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

// Register
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// CheckEmail
type CheckEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *CheckEmailRequest) Reset() {
	*x = CheckEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEmailRequest) ProtoMessage() {}

func (x *CheckEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEmailRequest.ProtoReflect.Descriptor instead.
func (*CheckEmailRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *CheckEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CheckEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=Exists,proto3" json:"Exists,omitempty"`
}

func (x *CheckEmailResponse) Reset() {
	*x = CheckEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEmailResponse) ProtoMessage() {}

func (x *CheckEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEmailResponse.ProtoReflect.Descriptor instead.
func (*CheckEmailResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *CheckEmailResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

// GetUidByEmailAndPwd
type GetUidByEmailAndPwdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *GetUidByEmailAndPwdRequest) Reset() {
	*x = GetUidByEmailAndPwdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUidByEmailAndPwdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUidByEmailAndPwdRequest) ProtoMessage() {}

func (x *GetUidByEmailAndPwdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUidByEmailAndPwdRequest.ProtoReflect.Descriptor instead.
func (*GetUidByEmailAndPwdRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUidByEmailAndPwdRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUidByEmailAndPwdRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetUidByEmailAndPwdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *GetUidByEmailAndPwdResponse) Reset() {
	*x = GetUidByEmailAndPwdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUidByEmailAndPwdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUidByEmailAndPwdResponse) ProtoMessage() {}

func (x *GetUidByEmailAndPwdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUidByEmailAndPwdResponse.ProtoReflect.Descriptor instead.
func (*GetUidByEmailAndPwdResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetUidByEmailAndPwdResponse) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// CheckUid
type CheckUidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *CheckUidRequest) Reset() {
	*x = CheckUidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUidRequest) ProtoMessage() {}

func (x *CheckUidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUidRequest.ProtoReflect.Descriptor instead.
func (*CheckUidRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *CheckUidRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type CheckUidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=Exists,proto3" json:"Exists,omitempty"`
}

func (x *CheckUidResponse) Reset() {
	*x = CheckUidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUidResponse) ProtoMessage() {}

func (x *CheckUidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUidResponse.ProtoReflect.Descriptor instead.
func (*CheckUidResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *CheckUidResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x24, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x22, 0x4e, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x69, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x6e, 0x64, 0x50, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x2f, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x55, 0x69, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x6e, 0x64, 0x50, 0x77, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69,
	0x64, 0x22, 0x23, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x32, 0xc9, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x66, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x69, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x41, 0x6e, 0x64, 0x50, 0x77, 0x64, 0x12, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x69, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x6e, 0x64, 0x50, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x69, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x6e, 0x64, 0x50, 0x77, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x55, 0x69, 0x64, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x55, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),             // 0: userclient.RegisterRequest
	(*RegisterResponse)(nil),            // 1: userclient.RegisterResponse
	(*CheckEmailRequest)(nil),           // 2: userclient.CheckEmailRequest
	(*CheckEmailResponse)(nil),          // 3: userclient.CheckEmailResponse
	(*GetUidByEmailAndPwdRequest)(nil),  // 4: userclient.GetUidByEmailAndPwdRequest
	(*GetUidByEmailAndPwdResponse)(nil), // 5: userclient.GetUidByEmailAndPwdResponse
	(*CheckUidRequest)(nil),             // 6: userclient.CheckUidRequest
	(*CheckUidResponse)(nil),            // 7: userclient.CheckUidResponse
}
var file_user_proto_depIdxs = []int32{
	0, // 0: userclient.User.Register:input_type -> userclient.RegisterRequest
	2, // 1: userclient.User.CheckEmail:input_type -> userclient.CheckEmailRequest
	4, // 2: userclient.User.GetUidByEmailAndPwd:input_type -> userclient.GetUidByEmailAndPwdRequest
	6, // 3: userclient.User.CheckUid:input_type -> userclient.CheckUidRequest
	1, // 4: userclient.User.Register:output_type -> userclient.RegisterResponse
	3, // 5: userclient.User.CheckEmail:output_type -> userclient.CheckEmailResponse
	5, // 6: userclient.User.GetUidByEmailAndPwd:output_type -> userclient.GetUidByEmailAndPwdResponse
	7, // 7: userclient.User.CheckUid:output_type -> userclient.CheckUidResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEmailRequest); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEmailResponse); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUidByEmailAndPwdRequest); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUidByEmailAndPwdResponse); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUidRequest); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUidResponse); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
