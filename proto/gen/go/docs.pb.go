// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: docs.proto

package docsv1

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

type UserTDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *UserTDocument) Reset() {
	*x = UserTDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTDocument) ProtoMessage() {}

func (x *UserTDocument) ProtoReflect() protoreflect.Message {
	mi := &file_docs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTDocument.ProtoReflect.Descriptor instead.
func (*UserTDocument) Descriptor() ([]byte, []int) {
	return file_docs_proto_rawDescGZIP(), []int{0}
}

func (x *UserTDocument) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UserTDocument) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url            string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	PubDate        uint64 `protobuf:"varint,2,opt,name=PubDate,proto3" json:"PubDate,omitempty"`
	FetchTime      uint64 `protobuf:"varint,3,opt,name=FetchTime,proto3" json:"FetchTime,omitempty"`
	Text           string `protobuf:"bytes,4,opt,name=Text,proto3" json:"Text,omitempty"`
	FirstFetchTime uint64 `protobuf:"varint,5,opt,name=FirstFetchTime,proto3" json:"FirstFetchTime,omitempty"`
}

func (x *TDocument) Reset() {
	*x = TDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDocument) ProtoMessage() {}

func (x *TDocument) ProtoReflect() protoreflect.Message {
	mi := &file_docs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDocument.ProtoReflect.Descriptor instead.
func (*TDocument) Descriptor() ([]byte, []int) {
	return file_docs_proto_rawDescGZIP(), []int{1}
}

func (x *TDocument) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *TDocument) GetPubDate() uint64 {
	if x != nil {
		return x.PubDate
	}
	return 0
}

func (x *TDocument) GetFetchTime() uint64 {
	if x != nil {
		return x.FetchTime
	}
	return 0
}

func (x *TDocument) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TDocument) GetFirstFetchTime() uint64 {
	if x != nil {
		return x.FirstFetchTime
	}
	return 0
}

var File_docs_proto protoreflect.FileDescriptor

var file_docs_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x35, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x91,
	0x01, 0x0a, 0x09, 0x54, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x75, 0x62, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x50, 0x75, 0x62, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0e, 0x46, 0x69, 0x72, 0x73, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x32, 0x7f, 0x0a, 0x09, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x54, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x14, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x77, 0x6c, 0x63, 0x6d, 0x74, 0x75, 0x6e, 0x6b, 0x6e,
	0x77, 0x6e, 0x64, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x64, 0x6f,
	0x63, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_docs_proto_rawDescOnce sync.Once
	file_docs_proto_rawDescData = file_docs_proto_rawDesc
)

func file_docs_proto_rawDescGZIP() []byte {
	file_docs_proto_rawDescOnce.Do(func() {
		file_docs_proto_rawDescData = protoimpl.X.CompressGZIP(file_docs_proto_rawDescData)
	})
	return file_docs_proto_rawDescData
}

var file_docs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_docs_proto_goTypes = []any{
	(*UserTDocument)(nil), // 0: documents.UserTDocument
	(*TDocument)(nil),     // 1: documents.TDocument
}
var file_docs_proto_depIdxs = []int32{
	0, // 0: documents.Documents.Create:input_type -> documents.UserTDocument
	0, // 1: documents.Documents.Update:input_type -> documents.UserTDocument
	1, // 2: documents.Documents.Create:output_type -> documents.TDocument
	1, // 3: documents.Documents.Update:output_type -> documents.TDocument
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_docs_proto_init() }
func file_docs_proto_init() {
	if File_docs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_docs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserTDocument); i {
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
		file_docs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TDocument); i {
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
			RawDescriptor: file_docs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_docs_proto_goTypes,
		DependencyIndexes: file_docs_proto_depIdxs,
		MessageInfos:      file_docs_proto_msgTypes,
	}.Build()
	File_docs_proto = out.File
	file_docs_proto_rawDesc = nil
	file_docs_proto_goTypes = nil
	file_docs_proto_depIdxs = nil
}
