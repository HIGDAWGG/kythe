// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: kythe/proto/identifier.proto

package identifier_go_proto

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

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier         string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Corpus             []string `protobuf:"bytes,2,rep,name=corpus,proto3" json:"corpus,omitempty"`
	Languages          []string `protobuf:"bytes,3,rep,name=languages,proto3" json:"languages,omitempty"`
	PickCanonicalNodes bool     `protobuf:"varint,4,opt,name=pick_canonical_nodes,json=pickCanonicalNodes,proto3" json:"pick_canonical_nodes,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_identifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_identifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_kythe_proto_identifier_proto_rawDescGZIP(), []int{0}
}

func (x *FindRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *FindRequest) GetCorpus() []string {
	if x != nil {
		return x.Corpus
	}
	return nil
}

func (x *FindRequest) GetLanguages() []string {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *FindRequest) GetPickCanonicalNodes() bool {
	if x != nil {
		return x.PickCanonicalNodes
	}
	return false
}

type FindReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches []*FindReply_Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
}

func (x *FindReply) Reset() {
	*x = FindReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_identifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindReply) ProtoMessage() {}

func (x *FindReply) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_identifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindReply.ProtoReflect.Descriptor instead.
func (*FindReply) Descriptor() ([]byte, []int) {
	return file_kythe_proto_identifier_proto_rawDescGZIP(), []int{1}
}

func (x *FindReply) GetMatches() []*FindReply_Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

type FindReply_Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket        string `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	NodeKind      string `protobuf:"bytes,2,opt,name=node_kind,json=nodeKind,proto3" json:"node_kind,omitempty"`
	NodeSubkind   string `protobuf:"bytes,3,opt,name=node_subkind,json=nodeSubkind,proto3" json:"node_subkind,omitempty"`
	BaseName      string `protobuf:"bytes,4,opt,name=base_name,json=baseName,proto3" json:"base_name,omitempty"`
	QualifiedName string `protobuf:"bytes,5,opt,name=qualified_name,json=qualifiedName,proto3" json:"qualified_name,omitempty"`
}

func (x *FindReply_Match) Reset() {
	*x = FindReply_Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_identifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindReply_Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindReply_Match) ProtoMessage() {}

func (x *FindReply_Match) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_identifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindReply_Match.ProtoReflect.Descriptor instead.
func (*FindReply_Match) Descriptor() ([]byte, []int) {
	return file_kythe_proto_identifier_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FindReply_Match) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

func (x *FindReply_Match) GetNodeKind() string {
	if x != nil {
		return x.NodeKind
	}
	return ""
}

func (x *FindReply_Match) GetNodeSubkind() string {
	if x != nil {
		return x.NodeSubkind
	}
	return ""
}

func (x *FindReply_Match) GetBaseName() string {
	if x != nil {
		return x.BaseName
	}
	return ""
}

func (x *FindReply_Match) GetQualifiedName() string {
	if x != nil {
		return x.QualifiedName
	}
	return ""
}

var File_kythe_proto_identifier_proto protoreflect.FileDescriptor

var file_kythe_proto_identifier_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x0b,
	0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x72, 0x70, 0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x72,
	0x70, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x70, 0x69, 0x63, 0x6b, 0x43, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x22, 0xe9, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x1a, 0xa3, 0x01, 0x0a, 0x05, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x6f, 0x64, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x73, 0x75, 0x62, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x6f, 0x64, 0x65, 0x53, 0x75, 0x62, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0x4d, 0x0a, 0x11, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x6b,
	0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x4b,
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x28, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x6b, 0x79, 0x74, 0x68,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_kythe_proto_identifier_proto_rawDescOnce sync.Once
	file_kythe_proto_identifier_proto_rawDescData = file_kythe_proto_identifier_proto_rawDesc
)

func file_kythe_proto_identifier_proto_rawDescGZIP() []byte {
	file_kythe_proto_identifier_proto_rawDescOnce.Do(func() {
		file_kythe_proto_identifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_kythe_proto_identifier_proto_rawDescData)
	})
	return file_kythe_proto_identifier_proto_rawDescData
}

var file_kythe_proto_identifier_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kythe_proto_identifier_proto_goTypes = []interface{}{
	(*FindRequest)(nil),     // 0: kythe.proto.FindRequest
	(*FindReply)(nil),       // 1: kythe.proto.FindReply
	(*FindReply_Match)(nil), // 2: kythe.proto.FindReply.Match
}
var file_kythe_proto_identifier_proto_depIdxs = []int32{
	2, // 0: kythe.proto.FindReply.matches:type_name -> kythe.proto.FindReply.Match
	0, // 1: kythe.proto.IdentifierService.Find:input_type -> kythe.proto.FindRequest
	1, // 2: kythe.proto.IdentifierService.Find:output_type -> kythe.proto.FindReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kythe_proto_identifier_proto_init() }
func file_kythe_proto_identifier_proto_init() {
	if File_kythe_proto_identifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kythe_proto_identifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRequest); i {
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
		file_kythe_proto_identifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindReply); i {
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
		file_kythe_proto_identifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindReply_Match); i {
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
			RawDescriptor: file_kythe_proto_identifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kythe_proto_identifier_proto_goTypes,
		DependencyIndexes: file_kythe_proto_identifier_proto_depIdxs,
		MessageInfos:      file_kythe_proto_identifier_proto_msgTypes,
	}.Build()
	File_kythe_proto_identifier_proto = out.File
	file_kythe_proto_identifier_proto_rawDesc = nil
	file_kythe_proto_identifier_proto_goTypes = nil
	file_kythe_proto_identifier_proto_depIdxs = nil
}
