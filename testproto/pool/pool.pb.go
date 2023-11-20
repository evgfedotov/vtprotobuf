// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: pool/pool.proto

package pool

import (
	_ "github.com/planetscale/vtprotobuf/vtproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Struct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Struct) Reset() {
	*x = Struct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Struct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Struct) ProtoMessage() {}

func (x *Struct) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Struct.ProtoReflect.Descriptor instead.
func (*Struct) Descriptor() ([]byte, []int) {
	return file_pool_pool_proto_rawDescGZIP(), []int{0}
}

func (x *Struct) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_pool_pool_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type MemoryPoolExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo1 string `protobuf:"bytes,1,opt,name=foo1,proto3" json:"foo1,omitempty"`
	Foo2 uint64 `protobuf:"varint,2,opt,name=foo2,proto3" json:"foo2,omitempty"`
}

func (x *MemoryPoolExtension) Reset() {
	*x = MemoryPoolExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryPoolExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryPoolExtension) ProtoMessage() {}

func (x *MemoryPoolExtension) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryPoolExtension.ProtoReflect.Descriptor instead.
func (*MemoryPoolExtension) Descriptor() ([]byte, []int) {
	return file_pool_pool_proto_rawDescGZIP(), []int{2}
}

func (x *MemoryPoolExtension) GetFoo1() string {
	if x != nil {
		return x.Foo1
	}
	return ""
}

func (x *MemoryPoolExtension) GetFoo2() uint64 {
	if x != nil {
		return x.Foo2
	}
	return 0
}

var File_pool_pool_proto protoreflect.FileDescriptor

var file_pool_pool_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x12, 0x2b, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x41,
	0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x3a, 0x04, 0xa8, 0xa6, 0x1f, 0x01, 0x22, 0x3d, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x3a, 0x04, 0xa8, 0xa6, 0x1f, 0x01, 0x22, 0x43, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x50, 0x6f, 0x6f, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x6f, 0x6f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6f, 0x6f,
	0x31, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x66, 0x6f, 0x6f, 0x32, 0x3a, 0x04, 0xa8, 0xa6, 0x1f, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pool_pool_proto_rawDescOnce sync.Once
	file_pool_pool_proto_rawDescData = file_pool_pool_proto_rawDesc
)

func file_pool_pool_proto_rawDescGZIP() []byte {
	file_pool_pool_proto_rawDescOnce.Do(func() {
		file_pool_pool_proto_rawDescData = protoimpl.X.CompressGZIP(file_pool_pool_proto_rawDescData)
	})
	return file_pool_pool_proto_rawDescData
}

var file_pool_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pool_pool_proto_goTypes = []interface{}{
	(*Struct)(nil),                // 0: Struct
	(*Value)(nil),                 // 1: Value
	(*MemoryPoolExtension)(nil),   // 2: MemoryPoolExtension
	nil,                           // 3: Struct.FieldsEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_pool_pool_proto_depIdxs = []int32{
	3, // 0: Struct.fields:type_name -> Struct.FieldsEntry
	4, // 1: Value.date:type_name -> google.protobuf.Timestamp
	1, // 2: Struct.FieldsEntry.value:type_name -> Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pool_pool_proto_init() }
func file_pool_pool_proto_init() {
	if File_pool_pool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pool_pool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Struct); i {
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
		file_pool_pool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_pool_pool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryPoolExtension); i {
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
			RawDescriptor: file_pool_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pool_pool_proto_goTypes,
		DependencyIndexes: file_pool_pool_proto_depIdxs,
		MessageInfos:      file_pool_pool_proto_msgTypes,
	}.Build()
	File_pool_pool_proto = out.File
	file_pool_pool_proto_rawDesc = nil
	file_pool_pool_proto_goTypes = nil
	file_pool_pool_proto_depIdxs = nil
}
