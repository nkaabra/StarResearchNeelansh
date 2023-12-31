// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: listDemands.proto

package __

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

type ListDemandsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDemandsRequest) Reset() {
	*x = ListDemandsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listDemands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDemandsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDemandsRequest) ProtoMessage() {}

func (x *ListDemandsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listDemands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDemandsRequest.ProtoReflect.Descriptor instead.
func (*ListDemandsRequest) Descriptor() ([]byte, []int) {
	return file_listDemands_proto_rawDescGZIP(), []int{0}
}

type DemandInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName string `protobuf:"bytes,1,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Quantity    int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *DemandInfo) Reset() {
	*x = DemandInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listDemands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemandInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemandInfo) ProtoMessage() {}

func (x *DemandInfo) ProtoReflect() protoreflect.Message {
	mi := &file_listDemands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemandInfo.ProtoReflect.Descriptor instead.
func (*DemandInfo) Descriptor() ([]byte, []int) {
	return file_listDemands_proto_rawDescGZIP(), []int{1}
}

func (x *DemandInfo) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *DemandInfo) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ListDemandsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of current customer demands.
	Demands []*DemandInfo `protobuf:"bytes,1,rep,name=demands,proto3" json:"demands,omitempty"`
}

func (x *ListDemandsResponse) Reset() {
	*x = ListDemandsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listDemands_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDemandsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDemandsResponse) ProtoMessage() {}

func (x *ListDemandsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listDemands_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDemandsResponse.ProtoReflect.Descriptor instead.
func (*ListDemandsResponse) Descriptor() ([]byte, []int) {
	return file_listDemands_proto_rawDescGZIP(), []int{2}
}

func (x *ListDemandsResponse) GetDemands() []*DemandInfo {
	if x != nil {
		return x.Demands
	}
	return nil
}

var File_listDemands_proto protoreflect.FileDescriptor

var file_listDemands_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x0a, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x47, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x32, 0x66, 0x0a, 0x14, 0x44,
	0x65, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_listDemands_proto_rawDescOnce sync.Once
	file_listDemands_proto_rawDescData = file_listDemands_proto_rawDesc
)

func file_listDemands_proto_rawDescGZIP() []byte {
	file_listDemands_proto_rawDescOnce.Do(func() {
		file_listDemands_proto_rawDescData = protoimpl.X.CompressGZIP(file_listDemands_proto_rawDescData)
	})
	return file_listDemands_proto_rawDescData
}

var file_listDemands_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_listDemands_proto_goTypes = []interface{}{
	(*ListDemandsRequest)(nil),  // 0: listDemand.ListDemandsRequest
	(*DemandInfo)(nil),          // 1: listDemand.DemandInfo
	(*ListDemandsResponse)(nil), // 2: listDemand.ListDemandsResponse
}
var file_listDemands_proto_depIdxs = []int32{
	1, // 0: listDemand.ListDemandsResponse.demands:type_name -> listDemand.DemandInfo
	0, // 1: listDemand.DemandListingService.ListDemands:input_type -> listDemand.ListDemandsRequest
	2, // 2: listDemand.DemandListingService.ListDemands:output_type -> listDemand.ListDemandsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_listDemands_proto_init() }
func file_listDemands_proto_init() {
	if File_listDemands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_listDemands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDemandsRequest); i {
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
		file_listDemands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemandInfo); i {
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
		file_listDemands_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDemandsResponse); i {
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
			RawDescriptor: file_listDemands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listDemands_proto_goTypes,
		DependencyIndexes: file_listDemands_proto_depIdxs,
		MessageInfos:      file_listDemands_proto_msgTypes,
	}.Build()
	File_listDemands_proto = out.File
	file_listDemands_proto_rawDesc = nil
	file_listDemands_proto_goTypes = nil
	file_listDemands_proto_depIdxs = nil
}
