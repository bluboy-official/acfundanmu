// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: CommonStateSignalPKAccept.proto

package acproto

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

type CommonStateSignalPKAccept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *CommonStateSignalPKAccept) Reset() {
	*x = CommonStateSignalPKAccept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalPKAccept_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalPKAccept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalPKAccept) ProtoMessage() {}

func (x *CommonStateSignalPKAccept) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalPKAccept_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalPKAccept.ProtoReflect.Descriptor instead.
func (*CommonStateSignalPKAccept) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalPKAccept_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalPKAccept) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *CommonStateSignalPKAccept) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

var File_CommonStateSignalPKAccept_proto protoreflect.FileDescriptor

var file_CommonStateSignalPKAccept_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x50, 0x4b, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0x37, 0x0a,
	0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x50, 0x4b, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x62, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x61, 0x63, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalPKAccept_proto_rawDescOnce sync.Once
	file_CommonStateSignalPKAccept_proto_rawDescData = file_CommonStateSignalPKAccept_proto_rawDesc
)

func file_CommonStateSignalPKAccept_proto_rawDescGZIP() []byte {
	file_CommonStateSignalPKAccept_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalPKAccept_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalPKAccept_proto_rawDescData)
	})
	return file_CommonStateSignalPKAccept_proto_rawDescData
}

var file_CommonStateSignalPKAccept_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalPKAccept_proto_goTypes = []interface{}{
	(*CommonStateSignalPKAccept)(nil), // 0: AcFunDanmu.CommonStateSignalPKAccept
}
var file_CommonStateSignalPKAccept_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CommonStateSignalPKAccept_proto_init() }
func file_CommonStateSignalPKAccept_proto_init() {
	if File_CommonStateSignalPKAccept_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalPKAccept_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonStateSignalPKAccept); i {
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
			RawDescriptor: file_CommonStateSignalPKAccept_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalPKAccept_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalPKAccept_proto_depIdxs,
		MessageInfos:      file_CommonStateSignalPKAccept_proto_msgTypes,
	}.Build()
	File_CommonStateSignalPKAccept_proto = out.File
	file_CommonStateSignalPKAccept_proto_rawDesc = nil
	file_CommonStateSignalPKAccept_proto_goTypes = nil
	file_CommonStateSignalPKAccept_proto_depIdxs = nil
}
