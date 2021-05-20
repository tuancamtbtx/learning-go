// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/common/entities/test.proto

package entities

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

type ItemRank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Rank float64 `protobuf:"fixed64,2,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *ItemRank) Reset() {
	*x = ItemRank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_entities_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemRank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemRank) ProtoMessage() {}

func (x *ItemRank) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_entities_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemRank.ProtoReflect.Descriptor instead.
func (*ItemRank) Descriptor() ([]byte, []int) {
	return file_proto_common_entities_test_proto_rawDescGZIP(), []int{0}
}

func (x *ItemRank) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemRank) GetRank() float64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type ListRankItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ItemRank `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListRankItems) Reset() {
	*x = ListRankItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_entities_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRankItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRankItems) ProtoMessage() {}

func (x *ListRankItems) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_entities_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRankItems.ProtoReflect.Descriptor instead.
func (*ListRankItems) Descriptor() ([]byte, []int) {
	return file_proto_common_entities_test_proto_rawDescGZIP(), []int{1}
}

func (x *ListRankItems) GetItems() []*ItemRank {
	if x != nil {
		return x.Items
	}
	return nil
}

type NestedRankItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items map[string]*ListRankItems `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NestedRankItems) Reset() {
	*x = NestedRankItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_entities_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedRankItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedRankItems) ProtoMessage() {}

func (x *NestedRankItems) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_entities_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedRankItems.ProtoReflect.Descriptor instead.
func (*NestedRankItems) Descriptor() ([]byte, []int) {
	return file_proto_common_entities_test_proto_rawDescGZIP(), []int{2}
}

func (x *NestedRankItems) GetItems() map[string]*ListRankItems {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_common_entities_test_proto protoreflect.FileDescriptor

var file_proto_common_entities_test_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x61,
	0x6e, 0x6b, 0x22, 0x3f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x0f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x61,
	0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x40, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x61,
	0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x57, 0x0a, 0x0a, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61,
	0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_common_entities_test_proto_rawDescOnce sync.Once
	file_proto_common_entities_test_proto_rawDescData = file_proto_common_entities_test_proto_rawDesc
)

func file_proto_common_entities_test_proto_rawDescGZIP() []byte {
	file_proto_common_entities_test_proto_rawDescOnce.Do(func() {
		file_proto_common_entities_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_entities_test_proto_rawDescData)
	})
	return file_proto_common_entities_test_proto_rawDescData
}

var file_proto_common_entities_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_common_entities_test_proto_goTypes = []interface{}{
	(*ItemRank)(nil),        // 0: proto.entities.ItemRank
	(*ListRankItems)(nil),   // 1: proto.entities.ListRankItems
	(*NestedRankItems)(nil), // 2: proto.entities.NestedRankItems
	nil,                     // 3: proto.entities.NestedRankItems.ItemsEntry
}
var file_proto_common_entities_test_proto_depIdxs = []int32{
	0, // 0: proto.entities.ListRankItems.items:type_name -> proto.entities.ItemRank
	3, // 1: proto.entities.NestedRankItems.items:type_name -> proto.entities.NestedRankItems.ItemsEntry
	1, // 2: proto.entities.NestedRankItems.ItemsEntry.value:type_name -> proto.entities.ListRankItems
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_common_entities_test_proto_init() }
func file_proto_common_entities_test_proto_init() {
	if File_proto_common_entities_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_entities_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemRank); i {
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
		file_proto_common_entities_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRankItems); i {
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
		file_proto_common_entities_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedRankItems); i {
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
			RawDescriptor: file_proto_common_entities_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_entities_test_proto_goTypes,
		DependencyIndexes: file_proto_common_entities_test_proto_depIdxs,
		MessageInfos:      file_proto_common_entities_test_proto_msgTypes,
	}.Build()
	File_proto_common_entities_test_proto = out.File
	file_proto_common_entities_test_proto_rawDesc = nil
	file_proto_common_entities_test_proto_goTypes = nil
	file_proto_common_entities_test_proto_depIdxs = nil
}
