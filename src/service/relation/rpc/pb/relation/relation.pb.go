// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: relation.proto

package relation

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

// Follow
type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Fid int64 `protobuf:"varint,2,opt,name=Fid,proto3" json:"Fid,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{0}
}

func (x *FollowRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FollowRequest) GetFid() int64 {
	if x != nil {
		return x.Fid
	}
	return 0
}

type FollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FollowResponse) Reset() {
	*x = FollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowResponse) ProtoMessage() {}

func (x *FollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowResponse.ProtoReflect.Descriptor instead.
func (*FollowResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{1}
}

// Unfollow
type UnfollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Fid int64 `protobuf:"varint,2,opt,name=Fid,proto3" json:"Fid,omitempty"`
}

func (x *UnfollowRequest) Reset() {
	*x = UnfollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowRequest) ProtoMessage() {}

func (x *UnfollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowRequest.ProtoReflect.Descriptor instead.
func (*UnfollowRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{2}
}

func (x *UnfollowRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UnfollowRequest) GetFid() int64 {
	if x != nil {
		return x.Fid
	}
	return 0
}

type UnfollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnfollowResponse) Reset() {
	*x = UnfollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowResponse) ProtoMessage() {}

func (x *UnfollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowResponse.ProtoReflect.Descriptor instead.
func (*UnfollowResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{3}
}

// Relation
type RelationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fid          int64  `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	Relationship string `protobuf:"bytes,2,opt,name=Relationship,proto3" json:"Relationship,omitempty"`
	UpdateTime   string `protobuf:"bytes,3,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *RelationMessage) Reset() {
	*x = RelationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationMessage) ProtoMessage() {}

func (x *RelationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationMessage.ProtoReflect.Descriptor instead.
func (*RelationMessage) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{4}
}

func (x *RelationMessage) GetFid() int64 {
	if x != nil {
		return x.Fid
	}
	return 0
}

func (x *RelationMessage) GetRelationship() string {
	if x != nil {
		return x.Relationship
	}
	return ""
}

func (x *RelationMessage) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

// GetFollowers
type GetFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64  `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Size  int32  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	Order string `protobuf:"bytes,4,opt,name=Order,proto3" json:"Order,omitempty"`
}

func (x *GetFollowersRequest) Reset() {
	*x = GetFollowersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersRequest) ProtoMessage() {}

func (x *GetFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersRequest.ProtoReflect.Descriptor instead.
func (*GetFollowersRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{5}
}

func (x *GetFollowersRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetFollowersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFollowersRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetFollowersRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type GetFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followers []*RelationMessage `protobuf:"bytes,1,rep,name=Followers,proto3" json:"Followers,omitempty"`
}

func (x *GetFollowersResponse) Reset() {
	*x = GetFollowersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersResponse) ProtoMessage() {}

func (x *GetFollowersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersResponse.ProtoReflect.Descriptor instead.
func (*GetFollowersResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{6}
}

func (x *GetFollowersResponse) GetFollowers() []*RelationMessage {
	if x != nil {
		return x.Followers
	}
	return nil
}

// GetFollowings
type GetFollowingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64  `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Size  int32  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	Order string `protobuf:"bytes,4,opt,name=Order,proto3" json:"Order,omitempty"`
}

func (x *GetFollowingsRequest) Reset() {
	*x = GetFollowingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingsRequest) ProtoMessage() {}

func (x *GetFollowingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingsRequest.ProtoReflect.Descriptor instead.
func (*GetFollowingsRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{7}
}

func (x *GetFollowingsRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetFollowingsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFollowingsRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetFollowingsRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type GetFollowingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followings []*RelationMessage `protobuf:"bytes,1,rep,name=Followings,proto3" json:"Followings,omitempty"`
}

func (x *GetFollowingsResponse) Reset() {
	*x = GetFollowingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingsResponse) ProtoMessage() {}

func (x *GetFollowingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingsResponse.ProtoReflect.Descriptor instead.
func (*GetFollowingsResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{8}
}

func (x *GetFollowingsResponse) GetFollowings() []*RelationMessage {
	if x != nil {
		return x.Followings
	}
	return nil
}

// GetFollowerCount
type GetFollowerCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *GetFollowerCountRequest) Reset() {
	*x = GetFollowerCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowerCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowerCountRequest) ProtoMessage() {}

func (x *GetFollowerCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowerCountRequest.ProtoReflect.Descriptor instead.
func (*GetFollowerCountRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{9}
}

func (x *GetFollowerCountRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetFollowerCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *GetFollowerCountResponse) Reset() {
	*x = GetFollowerCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowerCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowerCountResponse) ProtoMessage() {}

func (x *GetFollowerCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowerCountResponse.ProtoReflect.Descriptor instead.
func (*GetFollowerCountResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{10}
}

func (x *GetFollowerCountResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// GetFollowingCount
type GetFollowingCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *GetFollowingCountRequest) Reset() {
	*x = GetFollowingCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingCountRequest) ProtoMessage() {}

func (x *GetFollowingCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingCountRequest.ProtoReflect.Descriptor instead.
func (*GetFollowingCountRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{11}
}

func (x *GetFollowingCountRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetFollowingCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *GetFollowingCountResponse) Reset() {
	*x = GetFollowingCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingCountResponse) ProtoMessage() {}

func (x *GetFollowingCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingCountResponse.ProtoReflect.Descriptor instead.
func (*GetFollowingCountResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{12}
}

func (x *GetFollowingCountResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_relation_proto protoreflect.FileDescriptor

var file_relation_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0x33, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x55, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x46, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x0a, 0x0f, 0x55, 0x6e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x46, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x46, 0x69, 0x64, 0x22, 0x12,
	0x0a, 0x10, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x67, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x46, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x55, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x66, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x58, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x2b, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xac, 0x04, 0x0a, 0x08,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x08, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1f, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55,
	0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x59, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x12, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x68, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relation_proto_rawDescOnce sync.Once
	file_relation_proto_rawDescData = file_relation_proto_rawDesc
)

func file_relation_proto_rawDescGZIP() []byte {
	file_relation_proto_rawDescOnce.Do(func() {
		file_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_relation_proto_rawDescData)
	})
	return file_relation_proto_rawDescData
}

var file_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_relation_proto_goTypes = []interface{}{
	(*FollowRequest)(nil),             // 0: relationclient.FollowRequest
	(*FollowResponse)(nil),            // 1: relationclient.FollowResponse
	(*UnfollowRequest)(nil),           // 2: relationclient.UnfollowRequest
	(*UnfollowResponse)(nil),          // 3: relationclient.UnfollowResponse
	(*RelationMessage)(nil),           // 4: relationclient.RelationMessage
	(*GetFollowersRequest)(nil),       // 5: relationclient.GetFollowersRequest
	(*GetFollowersResponse)(nil),      // 6: relationclient.GetFollowersResponse
	(*GetFollowingsRequest)(nil),      // 7: relationclient.GetFollowingsRequest
	(*GetFollowingsResponse)(nil),     // 8: relationclient.GetFollowingsResponse
	(*GetFollowerCountRequest)(nil),   // 9: relationclient.GetFollowerCountRequest
	(*GetFollowerCountResponse)(nil),  // 10: relationclient.GetFollowerCountResponse
	(*GetFollowingCountRequest)(nil),  // 11: relationclient.GetFollowingCountRequest
	(*GetFollowingCountResponse)(nil), // 12: relationclient.GetFollowingCountResponse
}
var file_relation_proto_depIdxs = []int32{
	4,  // 0: relationclient.GetFollowersResponse.Followers:type_name -> relationclient.RelationMessage
	4,  // 1: relationclient.GetFollowingsResponse.Followings:type_name -> relationclient.RelationMessage
	0,  // 2: relationclient.Relation.Follow:input_type -> relationclient.FollowRequest
	2,  // 3: relationclient.Relation.Unfollow:input_type -> relationclient.UnfollowRequest
	5,  // 4: relationclient.Relation.GetFollowers:input_type -> relationclient.GetFollowersRequest
	7,  // 5: relationclient.Relation.GetFollowings:input_type -> relationclient.GetFollowingsRequest
	9,  // 6: relationclient.Relation.GetFollowerCount:input_type -> relationclient.GetFollowerCountRequest
	11, // 7: relationclient.Relation.GetFollowingCount:input_type -> relationclient.GetFollowingCountRequest
	1,  // 8: relationclient.Relation.Follow:output_type -> relationclient.FollowResponse
	3,  // 9: relationclient.Relation.Unfollow:output_type -> relationclient.UnfollowResponse
	6,  // 10: relationclient.Relation.GetFollowers:output_type -> relationclient.GetFollowersResponse
	8,  // 11: relationclient.Relation.GetFollowings:output_type -> relationclient.GetFollowingsResponse
	10, // 12: relationclient.Relation.GetFollowerCount:output_type -> relationclient.GetFollowerCountResponse
	12, // 13: relationclient.Relation.GetFollowingCount:output_type -> relationclient.GetFollowingCountResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_relation_proto_init() }
func file_relation_proto_init() {
	if File_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowResponse); i {
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
		file_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfollowRequest); i {
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
		file_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfollowResponse); i {
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
		file_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationMessage); i {
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
		file_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersRequest); i {
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
		file_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersResponse); i {
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
		file_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingsRequest); i {
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
		file_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingsResponse); i {
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
		file_relation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowerCountRequest); i {
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
		file_relation_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowerCountResponse); i {
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
		file_relation_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingCountRequest); i {
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
		file_relation_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingCountResponse); i {
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
			RawDescriptor: file_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relation_proto_goTypes,
		DependencyIndexes: file_relation_proto_depIdxs,
		MessageInfos:      file_relation_proto_msgTypes,
	}.Build()
	File_relation_proto = out.File
	file_relation_proto_rawDesc = nil
	file_relation_proto_goTypes = nil
	file_relation_proto_depIdxs = nil
}
