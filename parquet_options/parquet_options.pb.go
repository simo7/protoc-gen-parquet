// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: parquet_options/parquet_options.proto

package parquet_options

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type TimestampType int32

const (
	TimestampType_TIMESTAMP_MILLIS TimestampType = 0
	TimestampType_TIMESTAMP_MICROS TimestampType = 1
	TimestampType_TIMESTAMP_NANOS  TimestampType = 2
)

// Enum value maps for TimestampType.
var (
	TimestampType_name = map[int32]string{
		0: "TIMESTAMP_MILLIS",
		1: "TIMESTAMP_MICROS",
		2: "TIMESTAMP_NANOS",
	}
	TimestampType_value = map[string]int32{
		"TIMESTAMP_MILLIS": 0,
		"TIMESTAMP_MICROS": 1,
		"TIMESTAMP_NANOS":  2,
	}
)

func (x TimestampType) Enum() *TimestampType {
	p := new(TimestampType)
	*p = x
	return p
}

func (x TimestampType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimestampType) Descriptor() protoreflect.EnumDescriptor {
	return file_parquet_options_parquet_options_proto_enumTypes[0].Descriptor()
}

func (TimestampType) Type() protoreflect.EnumType {
	return &file_parquet_options_parquet_options_proto_enumTypes[0]
}

func (x TimestampType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimestampType.Descriptor instead.
func (TimestampType) EnumDescriptor() ([]byte, []int) {
	return file_parquet_options_parquet_options_proto_rawDescGZIP(), []int{0}
}

type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// As long as table_name is not blank,
	// a schema is generated for top-level messages in each file.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parquet_options_parquet_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_parquet_options_parquet_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_parquet_options_parquet_options_proto_rawDescGZIP(), []int{0}
}

func (x *MessageOptions) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimestampType TimestampType `protobuf:"varint,1,opt,name=timestamp_type,json=timestampType,proto3,enum=parquet_options.TimestampType" json:"timestamp_type,omitempty"`
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parquet_options_parquet_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_parquet_options_parquet_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_parquet_options_parquet_options_proto_rawDescGZIP(), []int{1}
}

func (x *FieldOptions) GetTimestampType() TimestampType {
	if x != nil {
		return x.TimestampType
	}
	return TimestampType_TIMESTAMP_MILLIS
}

var file_parquet_options_parquet_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         1022,
		Name:          "parquet_options.message_opts",
		Tag:           "bytes,1022,opt,name=message_opts",
		Filename:      "parquet_options/parquet_options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         1022,
		Name:          "parquet_options.field_opts",
		Tag:           "bytes,1022,opt,name=field_opts",
		Filename:      "parquet_options/parquet_options.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional parquet_options.MessageOptions message_opts = 1022;
	E_MessageOpts = &file_parquet_options_parquet_options_proto_extTypes[0]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional parquet_options.FieldOptions field_opts = 1022;
	E_FieldOpts = &file_parquet_options_parquet_options_proto_extTypes[1]
)

var File_parquet_options_parquet_options_proto protoreflect.FileDescriptor

var file_parquet_options_parquet_options_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x0c, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x2a, 0x50, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50,
	0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x53, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4e, 0x41, 0x4e,
	0x4f, 0x53, 0x10, 0x02, 0x3a, 0x64, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6f, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfe, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x3a, 0x5c, 0x0a, 0x0a, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfe, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x61,
	0x72, 0x71, 0x75, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parquet_options_parquet_options_proto_rawDescOnce sync.Once
	file_parquet_options_parquet_options_proto_rawDescData = file_parquet_options_parquet_options_proto_rawDesc
)

func file_parquet_options_parquet_options_proto_rawDescGZIP() []byte {
	file_parquet_options_parquet_options_proto_rawDescOnce.Do(func() {
		file_parquet_options_parquet_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_parquet_options_parquet_options_proto_rawDescData)
	})
	return file_parquet_options_parquet_options_proto_rawDescData
}

var file_parquet_options_parquet_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_parquet_options_parquet_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_parquet_options_parquet_options_proto_goTypes = []interface{}{
	(TimestampType)(0),                // 0: parquet_options.TimestampType
	(*MessageOptions)(nil),            // 1: parquet_options.MessageOptions
	(*FieldOptions)(nil),              // 2: parquet_options.FieldOptions
	(*descriptor.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
}
var file_parquet_options_parquet_options_proto_depIdxs = []int32{
	0, // 0: parquet_options.FieldOptions.timestamp_type:type_name -> parquet_options.TimestampType
	3, // 1: parquet_options.message_opts:extendee -> google.protobuf.MessageOptions
	4, // 2: parquet_options.field_opts:extendee -> google.protobuf.FieldOptions
	1, // 3: parquet_options.message_opts:type_name -> parquet_options.MessageOptions
	2, // 4: parquet_options.field_opts:type_name -> parquet_options.FieldOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_parquet_options_parquet_options_proto_init() }
func file_parquet_options_parquet_options_proto_init() {
	if File_parquet_options_parquet_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parquet_options_parquet_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
		file_parquet_options_parquet_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions); i {
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
			RawDescriptor: file_parquet_options_parquet_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_parquet_options_parquet_options_proto_goTypes,
		DependencyIndexes: file_parquet_options_parquet_options_proto_depIdxs,
		EnumInfos:         file_parquet_options_parquet_options_proto_enumTypes,
		MessageInfos:      file_parquet_options_parquet_options_proto_msgTypes,
		ExtensionInfos:    file_parquet_options_parquet_options_proto_extTypes,
	}.Build()
	File_parquet_options_parquet_options_proto = out.File
	file_parquet_options_parquet_options_proto_rawDesc = nil
	file_parquet_options_parquet_options_proto_goTypes = nil
	file_parquet_options_parquet_options_proto_depIdxs = nil
}
