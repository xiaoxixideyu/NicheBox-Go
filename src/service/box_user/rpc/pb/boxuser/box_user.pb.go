// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: box_user.proto

package boxuser

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

type UserRole int32

const (
	UserRole_Owner  UserRole = 0
	UserRole_Admin  UserRole = 1
	UserRole_Member UserRole = 2
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "Owner",
		1: "Admin",
		2: "Member",
	}
	UserRole_value = map[string]int32{
		"Owner":  0,
		"Admin":  1,
		"Member": 2,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_box_user_proto_enumTypes[0].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_box_user_proto_enumTypes[0]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{0}
}

// AddOwner
type AddOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid int64 `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Uid int64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *AddOwnerRequest) Reset() {
	*x = AddOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOwnerRequest) ProtoMessage() {}

func (x *AddOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOwnerRequest.ProtoReflect.Descriptor instead.
func (*AddOwnerRequest) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{0}
}

func (x *AddOwnerRequest) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *AddOwnerRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type AddOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddOwnerResponse) Reset() {
	*x = AddOwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOwnerResponse) ProtoMessage() {}

func (x *AddOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOwnerResponse.ProtoReflect.Descriptor instead.
func (*AddOwnerResponse) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{1}
}

// AddBoxUser
type AddBoxUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid int64 `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Uid int64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *AddBoxUserRequest) Reset() {
	*x = AddBoxUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBoxUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBoxUserRequest) ProtoMessage() {}

func (x *AddBoxUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBoxUserRequest.ProtoReflect.Descriptor instead.
func (*AddBoxUserRequest) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{2}
}

func (x *AddBoxUserRequest) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *AddBoxUserRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type AddBoxUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddBoxUserResponse) Reset() {
	*x = AddBoxUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBoxUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBoxUserResponse) ProtoMessage() {}

func (x *AddBoxUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBoxUserResponse.ProtoReflect.Descriptor instead.
func (*AddBoxUserResponse) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{3}
}

// RemoveBoxUser
type RemoveBoxUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid int64 `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Uid int64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *RemoveBoxUserRequest) Reset() {
	*x = RemoveBoxUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveBoxUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBoxUserRequest) ProtoMessage() {}

func (x *RemoveBoxUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBoxUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveBoxUserRequest) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveBoxUserRequest) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *RemoveBoxUserRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type RemoveBoxUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveBoxUserResponse) Reset() {
	*x = RemoveBoxUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveBoxUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBoxUserResponse) ProtoMessage() {}

func (x *RemoveBoxUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBoxUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveBoxUserResponse) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{5}
}

// SetRole
type SetRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid  int64    `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Uid  int64    `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Role UserRole `protobuf:"varint,3,opt,name=Role,proto3,enum=boxuserclient.UserRole" json:"Role,omitempty"`
}

func (x *SetRoleRequest) Reset() {
	*x = SetRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoleRequest) ProtoMessage() {}

func (x *SetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoleRequest.ProtoReflect.Descriptor instead.
func (*SetRoleRequest) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{6}
}

func (x *SetRoleRequest) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *SetRoleRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SetRoleRequest) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_Owner
}

type SetRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetRoleResponse) Reset() {
	*x = SetRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoleResponse) ProtoMessage() {}

func (x *SetRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoleResponse.ProtoReflect.Descriptor instead.
func (*SetRoleResponse) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{7}
}

// GetRole
type GetRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid int64 `protobuf:"varint,1,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Uid int64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *GetRoleRequest) Reset() {
	*x = GetRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleRequest) ProtoMessage() {}

func (x *GetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleRequest.ProtoReflect.Descriptor instead.
func (*GetRoleRequest) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetRoleRequest) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *GetRoleRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool     `protobuf:"varint,1,opt,name=Exist,proto3" json:"Exist,omitempty"`
	Role  UserRole `protobuf:"varint,2,opt,name=Role,proto3,enum=boxuserclient.UserRole" json:"Role,omitempty"`
}

func (x *GetRoleResponse) Reset() {
	*x = GetRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleResponse) ProtoMessage() {}

func (x *GetRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleResponse.ProtoReflect.Descriptor instead.
func (*GetRoleResponse) Descriptor() ([]byte, []int) {
	return file_box_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetRoleResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *GetRoleResponse) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_Owner
}

var File_box_user_proto protoreflect.FileDescriptor

var file_box_user_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0x35, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x42, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x41, 0x64,
	0x64, 0x42, 0x6f, 0x78, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x42, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x55, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x78, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x42, 0x6f, 0x78, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x42, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42,
	0x6f, 0x78, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61,
	0x0a, 0x0e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x42,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x55, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x42, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x2a, 0x2c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x02, 0x32, 0xea,
	0x03, 0x0a, 0x07, 0x42, 0x6f, 0x78, 0x55, 0x73, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x42,
	0x6f, 0x78, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x78, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x78, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x78, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x62,
	0x6f, 0x78, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x42, 0x6f, 0x78, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x78, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x62,
	0x6f, 0x78, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f,
	0x78, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x62, 0x6f, 0x78, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_box_user_proto_rawDescOnce sync.Once
	file_box_user_proto_rawDescData = file_box_user_proto_rawDesc
)

func file_box_user_proto_rawDescGZIP() []byte {
	file_box_user_proto_rawDescOnce.Do(func() {
		file_box_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_box_user_proto_rawDescData)
	})
	return file_box_user_proto_rawDescData
}

var file_box_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_box_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_box_user_proto_goTypes = []interface{}{
	(UserRole)(0),                 // 0: boxuserclient.UserRole
	(*AddOwnerRequest)(nil),       // 1: boxuserclient.AddOwnerRequest
	(*AddOwnerResponse)(nil),      // 2: boxuserclient.AddOwnerResponse
	(*AddBoxUserRequest)(nil),     // 3: boxuserclient.AddBoxUserRequest
	(*AddBoxUserResponse)(nil),    // 4: boxuserclient.AddBoxUserResponse
	(*RemoveBoxUserRequest)(nil),  // 5: boxuserclient.RemoveBoxUserRequest
	(*RemoveBoxUserResponse)(nil), // 6: boxuserclient.RemoveBoxUserResponse
	(*SetRoleRequest)(nil),        // 7: boxuserclient.SetRoleRequest
	(*SetRoleResponse)(nil),       // 8: boxuserclient.SetRoleResponse
	(*GetRoleRequest)(nil),        // 9: boxuserclient.GetRoleRequest
	(*GetRoleResponse)(nil),       // 10: boxuserclient.GetRoleResponse
}
var file_box_user_proto_depIdxs = []int32{
	0,  // 0: boxuserclient.SetRoleRequest.Role:type_name -> boxuserclient.UserRole
	0,  // 1: boxuserclient.GetRoleResponse.Role:type_name -> boxuserclient.UserRole
	1,  // 2: boxuserclient.BoxUser.AddOwner:input_type -> boxuserclient.AddOwnerRequest
	1,  // 3: boxuserclient.BoxUser.AddOwnerRevert:input_type -> boxuserclient.AddOwnerRequest
	3,  // 4: boxuserclient.BoxUser.AddBoxUser:input_type -> boxuserclient.AddBoxUserRequest
	5,  // 5: boxuserclient.BoxUser.RemoveBoxUser:input_type -> boxuserclient.RemoveBoxUserRequest
	7,  // 6: boxuserclient.BoxUser.SetRole:input_type -> boxuserclient.SetRoleRequest
	9,  // 7: boxuserclient.BoxUser.GetRole:input_type -> boxuserclient.GetRoleRequest
	1,  // 8: boxuserclient.BoxUser.AddOwner:output_type -> boxuserclient.AddOwnerRequest
	1,  // 9: boxuserclient.BoxUser.AddOwnerRevert:output_type -> boxuserclient.AddOwnerRequest
	4,  // 10: boxuserclient.BoxUser.AddBoxUser:output_type -> boxuserclient.AddBoxUserResponse
	6,  // 11: boxuserclient.BoxUser.RemoveBoxUser:output_type -> boxuserclient.RemoveBoxUserResponse
	8,  // 12: boxuserclient.BoxUser.SetRole:output_type -> boxuserclient.SetRoleResponse
	10, // 13: boxuserclient.BoxUser.GetRole:output_type -> boxuserclient.GetRoleResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_box_user_proto_init() }
func file_box_user_proto_init() {
	if File_box_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_box_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOwnerRequest); i {
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
		file_box_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOwnerResponse); i {
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
		file_box_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBoxUserRequest); i {
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
		file_box_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBoxUserResponse); i {
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
		file_box_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveBoxUserRequest); i {
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
		file_box_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveBoxUserResponse); i {
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
		file_box_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRoleRequest); i {
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
		file_box_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRoleResponse); i {
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
		file_box_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleRequest); i {
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
		file_box_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleResponse); i {
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
			RawDescriptor: file_box_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_box_user_proto_goTypes,
		DependencyIndexes: file_box_user_proto_depIdxs,
		EnumInfos:         file_box_user_proto_enumTypes,
		MessageInfos:      file_box_user_proto_msgTypes,
	}.Build()
	File_box_user_proto = out.File
	file_box_user_proto_rawDesc = nil
	file_box_user_proto_goTypes = nil
	file_box_user_proto_depIdxs = nil
}
