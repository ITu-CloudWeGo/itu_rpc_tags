// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.2
// source: idl/tags.proto

package tags_service

import (
	context "context"
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

// 传入的pid和标签，但是存入的是pid和tid
type PidTidCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int64  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *PidTidCreateRequest) Reset() {
	*x = PidTidCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PidTidCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PidTidCreateRequest) ProtoMessage() {}

func (x *PidTidCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PidTidCreateRequest.ProtoReflect.Descriptor instead.
func (*PidTidCreateRequest) Descriptor() ([]byte, []int) {
	return file_idl_tags_proto_rawDescGZIP(), []int{0}
}

func (x *PidTidCreateRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *PidTidCreateRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type PidTidCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Pid    int64  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Tid    int64  `protobuf:"varint,4,opt,name=tid,proto3" json:"tid,omitempty"`
}

func (x *PidTidCreateResponse) Reset() {
	*x = PidTidCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PidTidCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PidTidCreateResponse) ProtoMessage() {}

func (x *PidTidCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PidTidCreateResponse.ProtoReflect.Descriptor instead.
func (*PidTidCreateResponse) Descriptor() ([]byte, []int) {
	return file_idl_tags_proto_rawDescGZIP(), []int{1}
}

func (x *PidTidCreateResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PidTidCreateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PidTidCreateResponse) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *PidTidCreateResponse) GetTid() int64 {
	if x != nil {
		return x.Tid
	}
	return 0
}

type GetTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid int64 `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
}

func (x *GetTagRequest) Reset() {
	*x = GetTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tags_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagRequest) ProtoMessage() {}

func (x *GetTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tags_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagRequest.ProtoReflect.Descriptor instead.
func (*GetTagRequest) Descriptor() ([]byte, []int) {
	return file_idl_tags_proto_rawDescGZIP(), []int{2}
}

func (x *GetTagRequest) GetTid() int64 {
	if x != nil {
		return x.Tid
	}
	return 0
}

type GetTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Tag    string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *GetTagResponse) Reset() {
	*x = GetTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tags_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagResponse) ProtoMessage() {}

func (x *GetTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tags_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagResponse.ProtoReflect.Descriptor instead.
func (*GetTagResponse) Descriptor() ([]byte, []int) {
	return file_idl_tags_proto_rawDescGZIP(), []int{3}
}

func (x *GetTagResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetTagResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetTagResponse) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type GetTagIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int64 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *GetTagIDRequest) Reset() {
	*x = GetTagIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tags_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagIDRequest) ProtoMessage() {}

func (x *GetTagIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tags_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagIDRequest.ProtoReflect.Descriptor instead.
func (*GetTagIDRequest) Descriptor() ([]byte, []int) {
	return file_idl_tags_proto_rawDescGZIP(), []int{4}
}

func (x *GetTagIDRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type GetTagIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Tids   []int64 `protobuf:"varint,3,rep,packed,name=tids,proto3" json:"tids,omitempty"`
}

func (x *GetTagIDResponse) Reset() {
	*x = GetTagIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tags_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagIDResponse) ProtoMessage() {}

func (x *GetTagIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tags_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagIDResponse.ProtoReflect.Descriptor instead.
func (*GetTagIDResponse) Descriptor() ([]byte, []int) {
	return file_idl_tags_proto_rawDescGZIP(), []int{5}
}

func (x *GetTagIDResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetTagIDResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetTagIDResponse) GetTids() []int64 {
	if x != nil {
		return x.Tids
	}
	return nil
}

var File_idl_tags_proto protoreflect.FileDescriptor

var file_idl_tags_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x39, 0x0a, 0x13, 0x50, 0x69, 0x64, 0x54, 0x69, 0x64,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x22, 0x64, 0x0a, 0x14, 0x50, 0x69, 0x64, 0x54, 0x69, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x74, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x67, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x50, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x64, 0x73, 0x32,
	0xcc, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x0c, 0x50, 0x69, 0x64, 0x54, 0x69, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x50, 0x69, 0x64, 0x54, 0x69, 0x64, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x61, 0x67,
	0x73, 0x2e, 0x50, 0x69, 0x64, 0x54, 0x69, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67,
	0x12, 0x13, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x67, 0x49, 0x44, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x69, 0x64, 0x12,
	0x15, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x54, 0x75,
	0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x57, 0x65, 0x47, 0x6f, 0x2f, 0x69, 0x74, 0x75, 0x5f, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65,
	0x6e, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_tags_proto_rawDescOnce sync.Once
	file_idl_tags_proto_rawDescData = file_idl_tags_proto_rawDesc
)

func file_idl_tags_proto_rawDescGZIP() []byte {
	file_idl_tags_proto_rawDescOnce.Do(func() {
		file_idl_tags_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_tags_proto_rawDescData)
	})
	return file_idl_tags_proto_rawDescData
}

var file_idl_tags_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_idl_tags_proto_goTypes = []interface{}{
	(*PidTidCreateRequest)(nil),  // 0: tags.PidTidCreateRequest
	(*PidTidCreateResponse)(nil), // 1: tags.PidTidCreateResponse
	(*GetTagRequest)(nil),        // 2: tags.GetTagRequest
	(*GetTagResponse)(nil),       // 3: tags.GetTagResponse
	(*GetTagIDRequest)(nil),      // 4: tags.GetTagIDRequest
	(*GetTagIDResponse)(nil),     // 5: tags.GetTagIDResponse
}
var file_idl_tags_proto_depIdxs = []int32{
	0, // 0: tags.TagsService.PidTidCreate:input_type -> tags.PidTidCreateRequest
	2, // 1: tags.TagsService.GetTag:input_type -> tags.GetTagRequest
	4, // 2: tags.TagsService.GetTagIDsWithPid:input_type -> tags.GetTagIDRequest
	1, // 3: tags.TagsService.PidTidCreate:output_type -> tags.PidTidCreateResponse
	3, // 4: tags.TagsService.GetTag:output_type -> tags.GetTagResponse
	5, // 5: tags.TagsService.GetTagIDsWithPid:output_type -> tags.GetTagIDResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_tags_proto_init() }
func file_idl_tags_proto_init() {
	if File_idl_tags_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_tags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PidTidCreateRequest); i {
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
		file_idl_tags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PidTidCreateResponse); i {
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
		file_idl_tags_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagRequest); i {
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
		file_idl_tags_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagResponse); i {
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
		file_idl_tags_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagIDRequest); i {
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
		file_idl_tags_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagIDResponse); i {
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
			RawDescriptor: file_idl_tags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_tags_proto_goTypes,
		DependencyIndexes: file_idl_tags_proto_depIdxs,
		MessageInfos:      file_idl_tags_proto_msgTypes,
	}.Build()
	File_idl_tags_proto = out.File
	file_idl_tags_proto_rawDesc = nil
	file_idl_tags_proto_goTypes = nil
	file_idl_tags_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.12.1. DO NOT EDIT.

type TagsService interface {
	PidTidCreate(ctx context.Context, req *PidTidCreateRequest) (res *PidTidCreateResponse, err error)
	GetTag(ctx context.Context, req *GetTagRequest) (res *GetTagResponse, err error)
	GetTagIDsWithPid(ctx context.Context, req *GetTagIDRequest) (res *GetTagIDResponse, err error)
}
