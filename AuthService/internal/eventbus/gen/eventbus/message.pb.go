// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: eventbus/message.proto

package routeguide

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRegistrated struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRegistrated) Reset() {
	*x = UserRegistrated{}
	mi := &file_eventbus_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegistrated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegistrated) ProtoMessage() {}

func (x *UserRegistrated) ProtoReflect() protoreflect.Message {
	mi := &file_eventbus_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegistrated.ProtoReflect.Descriptor instead.
func (*UserRegistrated) Descriptor() ([]byte, []int) {
	return file_eventbus_message_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegistrated) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRegistrated) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_eventbus_message_proto protoreflect.FileDescriptor

const file_eventbus_message_proto_rawDesc = "" +
	"\n" +
	"\x16eventbus/message.proto\"?\n" +
	"\x0fUserRegistrated\x12\x16\n" +
	"\x06UserId\x18\x01 \x01(\x03R\x06UserId\x12\x14\n" +
	"\x05Email\x18\x02 \x01(\tR\x05EmailB8Z6google.golang.org/grpc/examples/route_guide/routeguideb\x06proto3"

var (
	file_eventbus_message_proto_rawDescOnce sync.Once
	file_eventbus_message_proto_rawDescData []byte
)

func file_eventbus_message_proto_rawDescGZIP() []byte {
	file_eventbus_message_proto_rawDescOnce.Do(func() {
		file_eventbus_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eventbus_message_proto_rawDesc), len(file_eventbus_message_proto_rawDesc)))
	})
	return file_eventbus_message_proto_rawDescData
}

var file_eventbus_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eventbus_message_proto_goTypes = []any{
	(*UserRegistrated)(nil), // 0: UserRegistrated
}
var file_eventbus_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eventbus_message_proto_init() }
func file_eventbus_message_proto_init() {
	if File_eventbus_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eventbus_message_proto_rawDesc), len(file_eventbus_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eventbus_message_proto_goTypes,
		DependencyIndexes: file_eventbus_message_proto_depIdxs,
		MessageInfos:      file_eventbus_message_proto_msgTypes,
	}.Build()
	File_eventbus_message_proto = out.File
	file_eventbus_message_proto_goTypes = nil
	file_eventbus_message_proto_depIdxs = nil
}
