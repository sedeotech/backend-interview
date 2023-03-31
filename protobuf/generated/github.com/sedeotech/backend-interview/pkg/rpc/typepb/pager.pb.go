// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: type/pager.proto

package typepb

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

type PagerMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sorting string `protobuf:"bytes,1,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Order   string `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	Current int64  `protobuf:"varint,3,opt,name=current,proto3" json:"current,omitempty"`
	Total   int64  `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PagerMsg) Reset() {
	*x = PagerMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_pager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagerMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagerMsg) ProtoMessage() {}

func (x *PagerMsg) ProtoReflect() protoreflect.Message {
	mi := &file_type_pager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagerMsg.ProtoReflect.Descriptor instead.
func (*PagerMsg) Descriptor() ([]byte, []int) {
	return file_type_pager_proto_rawDescGZIP(), []int{0}
}

func (x *PagerMsg) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *PagerMsg) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *PagerMsg) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *PagerMsg) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type PageRequestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sorting   string `protobuf:"bytes,1,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Order     string `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	Requested int64  `protobuf:"varint,3,opt,name=requested,proto3" json:"requested,omitempty"`
	Size      int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *PageRequestMsg) Reset() {
	*x = PageRequestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_pager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequestMsg) ProtoMessage() {}

func (x *PageRequestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_type_pager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequestMsg.ProtoReflect.Descriptor instead.
func (*PageRequestMsg) Descriptor() ([]byte, []int) {
	return file_type_pager_proto_rawDescGZIP(), []int{1}
}

func (x *PageRequestMsg) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *PageRequestMsg) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *PageRequestMsg) GetRequested() int64 {
	if x != nil {
		return x.Requested
	}
	return 0
}

func (x *PageRequestMsg) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_type_pager_proto protoreflect.FileDescriptor

var file_type_pager_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x72, 0x0a, 0x0e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x64, 0x65, 0x6f, 0x74, 0x65, 0x63, 0x68,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_type_pager_proto_rawDescOnce sync.Once
	file_type_pager_proto_rawDescData = file_type_pager_proto_rawDesc
)

func file_type_pager_proto_rawDescGZIP() []byte {
	file_type_pager_proto_rawDescOnce.Do(func() {
		file_type_pager_proto_rawDescData = protoimpl.X.CompressGZIP(file_type_pager_proto_rawDescData)
	})
	return file_type_pager_proto_rawDescData
}

var file_type_pager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_type_pager_proto_goTypes = []interface{}{
	(*PagerMsg)(nil),       // 0: type.PagerMsg
	(*PageRequestMsg)(nil), // 1: type.PageRequestMsg
}
var file_type_pager_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_type_pager_proto_init() }
func file_type_pager_proto_init() {
	if File_type_pager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_type_pager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagerMsg); i {
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
		file_type_pager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequestMsg); i {
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
			RawDescriptor: file_type_pager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_type_pager_proto_goTypes,
		DependencyIndexes: file_type_pager_proto_depIdxs,
		MessageInfos:      file_type_pager_proto_msgTypes,
	}.Build()
	File_type_pager_proto = out.File
	file_type_pager_proto_rawDesc = nil
	file_type_pager_proto_goTypes = nil
	file_type_pager_proto_depIdxs = nil
}
