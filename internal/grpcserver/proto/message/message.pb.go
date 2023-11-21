// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: internal/grpcserver/proto/message/message.proto

package proto

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

type MessageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromName    string `protobuf:"bytes,1,opt,name=FromName,proto3" json:"FromName,omitempty"`
	ToUID       string `protobuf:"bytes,2,opt,name=ToUID,proto3" json:"ToUID,omitempty"`
	Type        int32  `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Context     string `protobuf:"bytes,4,opt,name=Context,proto3" json:"Context,omitempty"`
	GroupID     string `protobuf:"bytes,5,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	ContextType int32  `protobuf:"varint,6,opt,name=ContextType,proto3" json:"ContextType,omitempty"`
	CreateTime  int64  `protobuf:"varint,7,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *MessageInfo) Reset() {
	*x = MessageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcserver_proto_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageInfo) ProtoMessage() {}

func (x *MessageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcserver_proto_message_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageInfo.ProtoReflect.Descriptor instead.
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return file_internal_grpcserver_proto_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageInfo) GetFromName() string {
	if x != nil {
		return x.FromName
	}
	return ""
}

func (x *MessageInfo) GetToUID() string {
	if x != nil {
		return x.ToUID
	}
	return ""
}

func (x *MessageInfo) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MessageInfo) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *MessageInfo) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *MessageInfo) GetContextType() int32 {
	if x != nil {
		return x.ContextType
	}
	return 0
}

func (x *MessageInfo) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcserver_proto_message_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcserver_proto_message_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_internal_grpcserver_proto_message_message_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_internal_grpcserver_proto_message_message_proto protoreflect.FileDescriptor

var file_internal_grpcserver_proto_message_message_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x72,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x55, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x4a, 0x0a, 0x0d, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_grpcserver_proto_message_message_proto_rawDescOnce sync.Once
	file_internal_grpcserver_proto_message_message_proto_rawDescData = file_internal_grpcserver_proto_message_message_proto_rawDesc
)

func file_internal_grpcserver_proto_message_message_proto_rawDescGZIP() []byte {
	file_internal_grpcserver_proto_message_message_proto_rawDescOnce.Do(func() {
		file_internal_grpcserver_proto_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpcserver_proto_message_message_proto_rawDescData)
	})
	return file_internal_grpcserver_proto_message_message_proto_rawDescData
}

var file_internal_grpcserver_proto_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_grpcserver_proto_message_message_proto_goTypes = []interface{}{
	(*MessageInfo)(nil), // 0: message.MessageInfo
	(*Response)(nil),    // 1: message.Response
}
var file_internal_grpcserver_proto_message_message_proto_depIdxs = []int32{
	0, // 0: message.HandleMessage.StoreMessage:input_type -> message.MessageInfo
	1, // 1: message.HandleMessage.StoreMessage:output_type -> message.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_grpcserver_proto_message_message_proto_init() }
func file_internal_grpcserver_proto_message_message_proto_init() {
	if File_internal_grpcserver_proto_message_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpcserver_proto_message_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageInfo); i {
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
		file_internal_grpcserver_proto_message_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_internal_grpcserver_proto_message_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpcserver_proto_message_message_proto_goTypes,
		DependencyIndexes: file_internal_grpcserver_proto_message_message_proto_depIdxs,
		MessageInfos:      file_internal_grpcserver_proto_message_message_proto_msgTypes,
	}.Build()
	File_internal_grpcserver_proto_message_message_proto = out.File
	file_internal_grpcserver_proto_message_message_proto_rawDesc = nil
	file_internal_grpcserver_proto_message_message_proto_goTypes = nil
	file_internal_grpcserver_proto_message_message_proto_depIdxs = nil
}