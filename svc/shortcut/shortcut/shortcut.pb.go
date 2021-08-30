// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: shortcut.proto

package shortcut

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	api "shortlink/api"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InternalCreateShortcutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUrl string `protobuf:"bytes,1,opt,name=target_url,json=targetUrl,proto3" json:"target_url,omitempty"`
}

func (x *InternalCreateShortcutRequest) Reset() {
	*x = InternalCreateShortcutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortcut_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalCreateShortcutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalCreateShortcutRequest) ProtoMessage() {}

func (x *InternalCreateShortcutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortcut_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalCreateShortcutRequest.ProtoReflect.Descriptor instead.
func (*InternalCreateShortcutRequest) Descriptor() ([]byte, []int) {
	return file_shortcut_proto_rawDescGZIP(), []int{0}
}

func (x *InternalCreateShortcutRequest) GetTargetUrl() string {
	if x != nil {
		return x.TargetUrl
	}
	return ""
}

type InternalDeleteShortcutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlToken string `protobuf:"bytes,1,opt,name=url_token,json=urlToken,proto3" json:"url_token,omitempty"`
}

func (x *InternalDeleteShortcutRequest) Reset() {
	*x = InternalDeleteShortcutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortcut_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalDeleteShortcutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalDeleteShortcutRequest) ProtoMessage() {}

func (x *InternalDeleteShortcutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortcut_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalDeleteShortcutRequest.ProtoReflect.Descriptor instead.
func (*InternalDeleteShortcutRequest) Descriptor() ([]byte, []int) {
	return file_shortcut_proto_rawDescGZIP(), []int{1}
}

func (x *InternalDeleteShortcutRequest) GetUrlToken() string {
	if x != nil {
		return x.UrlToken
	}
	return ""
}

var File_shortcut_proto protoreflect.FileDescriptor

var file_shortcut_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x1d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x55, 0x72, 0x6c, 0x22, 0x3c, 0x0a, 0x1d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x72, 0x6c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x32, 0x91, 0x02, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x59, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63,
	0x75, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69,
	0x6e, 0x6b, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_shortcut_proto_rawDescOnce sync.Once
	file_shortcut_proto_rawDescData = file_shortcut_proto_rawDesc
)

func file_shortcut_proto_rawDescGZIP() []byte {
	file_shortcut_proto_rawDescOnce.Do(func() {
		file_shortcut_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortcut_proto_rawDescData)
	})
	return file_shortcut_proto_rawDescData
}

var file_shortcut_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shortcut_proto_goTypes = []interface{}{
	(*InternalCreateShortcutRequest)(nil), // 0: protos.InternalCreateShortcutRequest
	(*InternalDeleteShortcutRequest)(nil), // 1: protos.InternalDeleteShortcutRequest
	(*api.GetShortcutRequest)(nil),        // 2: protos.GetShortcutRequest
	(*api.CreateShortcutResponse)(nil),    // 3: protos.CreateShortcutResponse
	(*api.DeleteShortcutResponse)(nil),    // 4: protos.DeleteShortcutResponse
	(*api.GetShortcutResponse)(nil),       // 5: protos.GetShortcutResponse
}
var file_shortcut_proto_depIdxs = []int32{
	0, // 0: protos.ShortcutService.CreateShortcut:input_type -> protos.InternalCreateShortcutRequest
	1, // 1: protos.ShortcutService.DeleteShortcut:input_type -> protos.InternalDeleteShortcutRequest
	2, // 2: protos.ShortcutService.GetShortcut:input_type -> protos.GetShortcutRequest
	3, // 3: protos.ShortcutService.CreateShortcut:output_type -> protos.CreateShortcutResponse
	4, // 4: protos.ShortcutService.DeleteShortcut:output_type -> protos.DeleteShortcutResponse
	5, // 5: protos.ShortcutService.GetShortcut:output_type -> protos.GetShortcutResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortcut_proto_init() }
func file_shortcut_proto_init() {
	if File_shortcut_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortcut_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalCreateShortcutRequest); i {
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
		file_shortcut_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalDeleteShortcutRequest); i {
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
			RawDescriptor: file_shortcut_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortcut_proto_goTypes,
		DependencyIndexes: file_shortcut_proto_depIdxs,
		MessageInfos:      file_shortcut_proto_msgTypes,
	}.Build()
	File_shortcut_proto = out.File
	file_shortcut_proto_rawDesc = nil
	file_shortcut_proto_goTypes = nil
	file_shortcut_proto_depIdxs = nil
}
