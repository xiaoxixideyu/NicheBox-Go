// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: box-content.proto

package box_content

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

// GetPostList
type GetPostListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoxID int64  `protobuf:"varint,1,opt,name=BoxID,proto3" json:"BoxID,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Size  int32  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	Order string `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetPostListRequest) Reset() {
	*x = GetPostListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostListRequest) ProtoMessage() {}

func (x *GetPostListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostListRequest.ProtoReflect.Descriptor instead.
func (*GetPostListRequest) Descriptor() ([]byte, []int) {
	return file_box_content_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostListRequest) GetBoxID() int64 {
	if x != nil {
		return x.BoxID
	}
	return 0
}

func (x *GetPostListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPostListRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetPostListRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type GetPostListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDs []int64 `protobuf:"varint,1,rep,packed,name=IDs,proto3" json:"IDs,omitempty"`
}

func (x *GetPostListResponse) Reset() {
	*x = GetPostListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostListResponse) ProtoMessage() {}

func (x *GetPostListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostListResponse.ProtoReflect.Descriptor instead.
func (*GetPostListResponse) Descriptor() ([]byte, []int) {
	return file_box_content_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostListResponse) GetIDs() []int64 {
	if x != nil {
		return x.IDs
	}
	return nil
}

type ModifiedPostInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID    int64  `protobuf:"varint,1,opt,name=PostID,proto3" json:"PostID,omitempty"`
	Time      string `protobuf:"bytes,2,opt,name=Time,proto3" json:"Time,omitempty"`
	BoxID     int64  `protobuf:"varint,3,opt,name=BoxID,proto3" json:"BoxID,omitempty"`
	InfoCount int32  `protobuf:"varint,4,opt,name=InfoCount,proto3" json:"InfoCount,omitempty"`
}

func (x *ModifiedPostInfo) Reset() {
	*x = ModifiedPostInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifiedPostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifiedPostInfo) ProtoMessage() {}

func (x *ModifiedPostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_box_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifiedPostInfo.ProtoReflect.Descriptor instead.
func (*ModifiedPostInfo) Descriptor() ([]byte, []int) {
	return file_box_content_proto_rawDescGZIP(), []int{2}
}

func (x *ModifiedPostInfo) GetPostID() int64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

func (x *ModifiedPostInfo) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *ModifiedPostInfo) GetBoxID() int64 {
	if x != nil {
		return x.BoxID
	}
	return 0
}

func (x *ModifiedPostInfo) GetInfoCount() int32 {
	if x != nil {
		return x.InfoCount
	}
	return 0
}

// UpdateNewPosts
type UpdateNewPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPosts []*ModifiedPostInfo `protobuf:"bytes,1,rep,name=NewPosts,proto3" json:"NewPosts,omitempty"`
}

func (x *UpdateNewPostsRequest) Reset() {
	*x = UpdateNewPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNewPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNewPostsRequest) ProtoMessage() {}

func (x *UpdateNewPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNewPostsRequest.ProtoReflect.Descriptor instead.
func (*UpdateNewPostsRequest) Descriptor() ([]byte, []int) {
	return file_box_content_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNewPostsRequest) GetNewPosts() []*ModifiedPostInfo {
	if x != nil {
		return x.NewPosts
	}
	return nil
}

type UpdateNewPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateNewPostsResponse) Reset() {
	*x = UpdateNewPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_content_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNewPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNewPostsResponse) ProtoMessage() {}

func (x *UpdateNewPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_content_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNewPostsResponse.ProtoReflect.Descriptor instead.
func (*UpdateNewPostsResponse) Descriptor() ([]byte, []int) {
	return file_box_content_proto_rawDescGZIP(), []int{4}
}

// UpdateDeletedPosts
type UpdateDeletedPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletedPosts []*ModifiedPostInfo `protobuf:"bytes,1,rep,name=DeletedPosts,proto3" json:"DeletedPosts,omitempty"`
}

func (x *UpdateDeletedPostsRequest) Reset() {
	*x = UpdateDeletedPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_content_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeletedPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeletedPostsRequest) ProtoMessage() {}

func (x *UpdateDeletedPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_content_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeletedPostsRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeletedPostsRequest) Descriptor() ([]byte, []int) {
	return file_box_content_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDeletedPostsRequest) GetDeletedPosts() []*ModifiedPostInfo {
	if x != nil {
		return x.DeletedPosts
	}
	return nil
}

type UpdateDeletedPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDeletedPostsResponse) Reset() {
	*x = UpdateDeletedPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_content_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeletedPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeletedPostsResponse) ProtoMessage() {}

func (x *UpdateDeletedPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_content_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeletedPostsResponse.ProtoReflect.Descriptor instead.
func (*UpdateDeletedPostsResponse) Descriptor() ([]byte, []int) {
	return file_box_content_proto_rawDescGZIP(), []int{6}
}

var File_box_content_proto protoreflect.FileDescriptor

var file_box_content_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x78, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x62, 0x6f, 0x78, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x42,
	0x6f, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x42, 0x6f, 0x78, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22,
	0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x49, 0x44, 0x73, 0x22, 0x72, 0x0a, 0x10, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x6f, 0x78, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x42, 0x6f, 0x78, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x6f, 0x78, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x4e, 0x65, 0x77,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x63, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x6f, 0x78, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xbe, 0x02, 0x0a, 0x0a, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x24, 0x2e, 0x62, 0x6f, 0x78, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x78, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12,
	0x27, 0x2e, 0x62, 0x6f, 0x78, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x6f, 0x78, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x2b, 0x2e, 0x62, 0x6f, 0x78, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x62, 0x6f, 0x78, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x62, 0x6f, 0x78, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_box_content_proto_rawDescOnce sync.Once
	file_box_content_proto_rawDescData = file_box_content_proto_rawDesc
)

func file_box_content_proto_rawDescGZIP() []byte {
	file_box_content_proto_rawDescOnce.Do(func() {
		file_box_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_box_content_proto_rawDescData)
	})
	return file_box_content_proto_rawDescData
}

var file_box_content_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_box_content_proto_goTypes = []interface{}{
	(*GetPostListRequest)(nil),         // 0: boxcontentclient.GetPostListRequest
	(*GetPostListResponse)(nil),        // 1: boxcontentclient.GetPostListResponse
	(*ModifiedPostInfo)(nil),           // 2: boxcontentclient.ModifiedPostInfo
	(*UpdateNewPostsRequest)(nil),      // 3: boxcontentclient.UpdateNewPostsRequest
	(*UpdateNewPostsResponse)(nil),     // 4: boxcontentclient.UpdateNewPostsResponse
	(*UpdateDeletedPostsRequest)(nil),  // 5: boxcontentclient.UpdateDeletedPostsRequest
	(*UpdateDeletedPostsResponse)(nil), // 6: boxcontentclient.UpdateDeletedPostsResponse
}
var file_box_content_proto_depIdxs = []int32{
	2, // 0: boxcontentclient.UpdateNewPostsRequest.NewPosts:type_name -> boxcontentclient.ModifiedPostInfo
	2, // 1: boxcontentclient.UpdateDeletedPostsRequest.DeletedPosts:type_name -> boxcontentclient.ModifiedPostInfo
	0, // 2: boxcontentclient.BoxContent.GetPostList:input_type -> boxcontentclient.GetPostListRequest
	3, // 3: boxcontentclient.BoxContent.UpdateNewPosts:input_type -> boxcontentclient.UpdateNewPostsRequest
	5, // 4: boxcontentclient.BoxContent.UpdateDeletedPosts:input_type -> boxcontentclient.UpdateDeletedPostsRequest
	1, // 5: boxcontentclient.BoxContent.GetPostList:output_type -> boxcontentclient.GetPostListResponse
	4, // 6: boxcontentclient.BoxContent.UpdateNewPosts:output_type -> boxcontentclient.UpdateNewPostsResponse
	6, // 7: boxcontentclient.BoxContent.UpdateDeletedPosts:output_type -> boxcontentclient.UpdateDeletedPostsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_box_content_proto_init() }
func file_box_content_proto_init() {
	if File_box_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_box_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostListRequest); i {
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
		file_box_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostListResponse); i {
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
		file_box_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifiedPostInfo); i {
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
		file_box_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNewPostsRequest); i {
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
		file_box_content_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNewPostsResponse); i {
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
		file_box_content_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeletedPostsRequest); i {
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
		file_box_content_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeletedPostsResponse); i {
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
			RawDescriptor: file_box_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_box_content_proto_goTypes,
		DependencyIndexes: file_box_content_proto_depIdxs,
		MessageInfos:      file_box_content_proto_msgTypes,
	}.Build()
	File_box_content_proto = out.File
	file_box_content_proto_rawDesc = nil
	file_box_content_proto_goTypes = nil
	file_box_content_proto_depIdxs = nil
}
