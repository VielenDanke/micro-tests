// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: github.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LookupUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *LookupUserReq) Reset() {
	*x = LookupUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupUserReq) ProtoMessage() {}

func (x *LookupUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_github_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupUserReq.ProtoReflect.Descriptor instead.
func (*LookupUserReq) Descriptor() ([]byte, []int) {
	return file_github_proto_rawDescGZIP(), []int{0}
}

func (x *LookupUserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type LookupUserRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LookupUserRsp) Reset() {
	*x = LookupUserRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupUserRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupUserRsp) ProtoMessage() {}

func (x *LookupUserRsp) ProtoReflect() protoreflect.Message {
	mi := &file_github_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupUserRsp.ProtoReflect.Descriptor instead.
func (*LookupUserRsp) Descriptor() ([]byte, []int) {
	return file_github_proto_rawDescGZIP(), []int{1}
}

func (x *LookupUserRsp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_github_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_github_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_github_proto protoreflect.FileDescriptor

var file_github_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x23, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9f, 0x01, 0x0a, 0x06, 0x47, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x12, 0x94, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x4c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x73,
	0x70, 0x22, 0x58, 0x92, 0x41, 0x3c, 0x2a, 0x0a, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x4a, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a,
	0x0e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x11, 0x0a, 0x0f, 0x1a, 0x0d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x3b, 0x5a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x2f, 0x75, 0x6e, 0x69, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_proto_rawDescOnce sync.Once
	file_github_proto_rawDescData = file_github_proto_rawDesc
)

func file_github_proto_rawDescGZIP() []byte {
	file_github_proto_rawDescOnce.Do(func() {
		file_github_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_proto_rawDescData)
	})
	return file_github_proto_rawDescData
}

var file_github_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_proto_goTypes = []interface{}{
	(*LookupUserReq)(nil), // 0: github.LookupUserReq
	(*LookupUserRsp)(nil), // 1: github.LookupUserRsp
	(*Error)(nil),         // 2: github.Error
}
var file_github_proto_depIdxs = []int32{
	0, // 0: github.Github.LookupUser:input_type -> github.LookupUserReq
	1, // 1: github.Github.LookupUser:output_type -> github.LookupUserRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_proto_init() }
func file_github_proto_init() {
	if File_github_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupUserReq); i {
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
		file_github_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupUserRsp); i {
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
		file_github_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_github_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_proto_goTypes,
		DependencyIndexes: file_github_proto_depIdxs,
		MessageInfos:      file_github_proto_msgTypes,
	}.Build()
	File_github_proto = out.File
	file_github_proto_rawDesc = nil
	file_github_proto_goTypes = nil
	file_github_proto_depIdxs = nil
}
