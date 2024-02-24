// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: longConn.proto

package longConn

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

// HandShake
type HandShakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       int64  `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	UserAgent string `protobuf:"bytes,2,opt,name=UserAgent,proto3" json:"UserAgent,omitempty"`
	Addr      string `protobuf:"bytes,3,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func (x *HandShakeRequest) Reset() {
	*x = HandShakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longConn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandShakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandShakeRequest) ProtoMessage() {}

func (x *HandShakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_longConn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandShakeRequest.ProtoReflect.Descriptor instead.
func (*HandShakeRequest) Descriptor() ([]byte, []int) {
	return file_longConn_proto_rawDescGZIP(), []int{0}
}

func (x *HandShakeRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *HandShakeRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *HandShakeRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type HandShakeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *HandShakeResponse) Reset() {
	*x = HandShakeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longConn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandShakeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandShakeResponse) ProtoMessage() {}

func (x *HandShakeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_longConn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandShakeResponse.ProtoReflect.Descriptor instead.
func (*HandShakeResponse) Descriptor() ([]byte, []int) {
	return file_longConn_proto_rawDescGZIP(), []int{1}
}

func (x *HandShakeResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_longConn_proto protoreflect.FileDescriptor

var file_longConn_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x6c, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0x56, 0x0a, 0x10, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x22, 0x29, 0x0a, 0x11, 0x48, 0x61, 0x6e, 0x64,
	0x53, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x32, 0x5c, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x6e, 0x12,
	0x50, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x20, 0x2e, 0x6c,
	0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x61,
	0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_longConn_proto_rawDescOnce sync.Once
	file_longConn_proto_rawDescData = file_longConn_proto_rawDesc
)

func file_longConn_proto_rawDescGZIP() []byte {
	file_longConn_proto_rawDescOnce.Do(func() {
		file_longConn_proto_rawDescData = protoimpl.X.CompressGZIP(file_longConn_proto_rawDescData)
	})
	return file_longConn_proto_rawDescData
}

var file_longConn_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_longConn_proto_goTypes = []interface{}{
	(*HandShakeRequest)(nil),  // 0: longconnclient.HandShakeRequest
	(*HandShakeResponse)(nil), // 1: longconnclient.HandShakeResponse
}
var file_longConn_proto_depIdxs = []int32{
	0, // 0: longconnclient.longConn.HandShake:input_type -> longconnclient.HandShakeRequest
	1, // 1: longconnclient.longConn.HandShake:output_type -> longconnclient.HandShakeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_longConn_proto_init() }
func file_longConn_proto_init() {
	if File_longConn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_longConn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandShakeRequest); i {
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
		file_longConn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandShakeResponse); i {
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
			RawDescriptor: file_longConn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_longConn_proto_goTypes,
		DependencyIndexes: file_longConn_proto_depIdxs,
		MessageInfos:      file_longConn_proto_msgTypes,
	}.Build()
	File_longConn_proto = out.File
	file_longConn_proto_rawDesc = nil
	file_longConn_proto_goTypes = nil
	file_longConn_proto_depIdxs = nil
}
