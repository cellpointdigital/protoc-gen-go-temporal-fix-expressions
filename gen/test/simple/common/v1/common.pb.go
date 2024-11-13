// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: test/simple/common/v1/common.proto

// buf:lint:ignore PACKAGE_DIRECTORY_MATCH

package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Example int32

const (
	Example_EXAMPLE_UNSPECIFIED Example = 0
	Example_EXAMPLE_FOO         Example = 1
)

// Enum value maps for Example.
var (
	Example_name = map[int32]string{
		0: "EXAMPLE_UNSPECIFIED",
		1: "EXAMPLE_FOO",
	}
	Example_value = map[string]int32{
		"EXAMPLE_UNSPECIFIED": 0,
		"EXAMPLE_FOO":         1,
	}
)

func (x Example) Enum() *Example {
	p := new(Example)
	*p = x
	return p
}

func (x Example) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Example) Descriptor() protoreflect.EnumDescriptor {
	return file_test_simple_common_v1_common_proto_enumTypes[0].Descriptor()
}

func (Example) Type() protoreflect.EnumType {
	return &file_test_simple_common_v1_common_proto_enumTypes[0]
}

func (x Example) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Example.Descriptor instead.
func (Example) EnumDescriptor() ([]byte, []int) {
	return file_test_simple_common_v1_common_proto_rawDescGZIP(), []int{0}
}

type PaginatedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Cursor []byte `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *PaginatedRequest) Reset() {
	*x = PaginatedRequest{}
	mi := &file_test_simple_common_v1_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginatedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginatedRequest) ProtoMessage() {}

func (x *PaginatedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_common_v1_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginatedRequest.ProtoReflect.Descriptor instead.
func (*PaginatedRequest) Descriptor() ([]byte, []int) {
	return file_test_simple_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *PaginatedRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PaginatedRequest) GetCursor() []byte {
	if x != nil {
		return x.Cursor
	}
	return nil
}

type PaginatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*anypb.Any `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextCursor []byte       `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
}

func (x *PaginatedResponse) Reset() {
	*x = PaginatedResponse{}
	mi := &file_test_simple_common_v1_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginatedResponse) ProtoMessage() {}

func (x *PaginatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_common_v1_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginatedResponse.ProtoReflect.Descriptor instead.
func (*PaginatedResponse) Descriptor() ([]byte, []int) {
	return file_test_simple_common_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *PaginatedResponse) GetItems() []*anypb.Any {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PaginatedResponse) GetNextCursor() []byte {
	if x != nil {
		return x.NextCursor
	}
	return nil
}

var File_test_simple_common_v1_common_proto protoreflect.FileDescriptor

var file_test_simple_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x10, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x60, 0x0a,
	0x11, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2a,
	0x33, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x58,
	0x41, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x58, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x46,
	0x4f, 0x4f, 0x10, 0x01, 0x42, 0x86, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x53, 0x43, 0xaa, 0x02, 0x1a, 0x4d, 0x79, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x26, 0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d,
	0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x3a, 0x3a, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_simple_common_v1_common_proto_rawDescOnce sync.Once
	file_test_simple_common_v1_common_proto_rawDescData = file_test_simple_common_v1_common_proto_rawDesc
)

func file_test_simple_common_v1_common_proto_rawDescGZIP() []byte {
	file_test_simple_common_v1_common_proto_rawDescOnce.Do(func() {
		file_test_simple_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_simple_common_v1_common_proto_rawDescData)
	})
	return file_test_simple_common_v1_common_proto_rawDescData
}

var file_test_simple_common_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_simple_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_simple_common_v1_common_proto_goTypes = []any{
	(Example)(0),              // 0: mycompany.simple.common.v1.Example
	(*PaginatedRequest)(nil),  // 1: mycompany.simple.common.v1.PaginatedRequest
	(*PaginatedResponse)(nil), // 2: mycompany.simple.common.v1.PaginatedResponse
	(*anypb.Any)(nil),         // 3: google.protobuf.Any
}
var file_test_simple_common_v1_common_proto_depIdxs = []int32{
	3, // 0: mycompany.simple.common.v1.PaginatedResponse.items:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test_simple_common_v1_common_proto_init() }
func file_test_simple_common_v1_common_proto_init() {
	if File_test_simple_common_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_simple_common_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_simple_common_v1_common_proto_goTypes,
		DependencyIndexes: file_test_simple_common_v1_common_proto_depIdxs,
		EnumInfos:         file_test_simple_common_v1_common_proto_enumTypes,
		MessageInfos:      file_test_simple_common_v1_common_proto_msgTypes,
	}.Build()
	File_test_simple_common_v1_common_proto = out.File
	file_test_simple_common_v1_common_proto_rawDesc = nil
	file_test_simple_common_v1_common_proto_goTypes = nil
	file_test_simple_common_v1_common_proto_depIdxs = nil
}
