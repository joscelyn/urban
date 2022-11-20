// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: urban.proto

package urban

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type HNPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url   string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *HNPost) Reset() {
	*x = HNPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urban_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HNPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HNPost) ProtoMessage() {}

func (x *HNPost) ProtoReflect() protoreflect.Message {
	mi := &file_urban_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HNPost.ProtoReflect.Descriptor instead.
func (*HNPost) Descriptor() ([]byte, []int) {
	return file_urban_proto_rawDescGZIP(), []int{0}
}

func (x *HNPost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *HNPost) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*HNPost `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urban_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urban_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_urban_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetPosts() []*HNPost {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetTopStoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*HNPost `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetTopStoriesResponse) Reset() {
	*x = GetTopStoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urban_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopStoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopStoriesResponse) ProtoMessage() {}

func (x *GetTopStoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urban_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopStoriesResponse.ProtoReflect.Descriptor instead.
func (*GetTopStoriesResponse) Descriptor() ([]byte, []int) {
	return file_urban_proto_rawDescGZIP(), []int{2}
}

func (x *GetTopStoriesResponse) GetPosts() []*HNPost {
	if x != nil {
		return x.Posts
	}
	return nil
}

type WhoisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *WhoisRequest) Reset() {
	*x = WhoisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urban_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoisRequest) ProtoMessage() {}

func (x *WhoisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urban_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoisRequest.ProtoReflect.Descriptor instead.
func (*WhoisRequest) Descriptor() ([]byte, []int) {
	return file_urban_proto_rawDescGZIP(), []int{3}
}

func (x *WhoisRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type WhoisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Karma  int64                  `protobuf:"varint,2,opt,name=karma,proto3" json:"karma,omitempty"`
	About  string                 `protobuf:"bytes,3,opt,name=about,proto3" json:"about,omitempty"`
	Joined *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=joined,proto3" json:"joined,omitempty"`
}

func (x *WhoisResponse) Reset() {
	*x = WhoisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urban_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoisResponse) ProtoMessage() {}

func (x *WhoisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urban_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoisResponse.ProtoReflect.Descriptor instead.
func (*WhoisResponse) Descriptor() ([]byte, []int) {
	return file_urban_proto_rawDescGZIP(), []int{4}
}

func (x *WhoisResponse) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *WhoisResponse) GetKarma() int64 {
	if x != nil {
		return x.Karma
	}
	return 0
}

func (x *WhoisResponse) GetAbout() string {
	if x != nil {
		return x.About
	}
	return ""
}

func (x *WhoisResponse) GetJoined() *timestamppb.Timestamp {
	if x != nil {
		return x.Joined
	}
	return nil
}

var File_urban_proto protoreflect.FileDescriptor

var file_urban_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x72, 0x62, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75,
	0x72, 0x62, 0x61, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x30, 0x0a, 0x06, 0x48, 0x4e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x33, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x72, 0x62, 0x61, 0x6e, 0x2e, 0x48, 0x4e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x3c, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x72, 0x62, 0x61, 0x6e, 0x2e, 0x48, 0x4e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x2a, 0x0a, 0x0c, 0x57, 0x68, 0x6f, 0x69, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x57, 0x68, 0x6f, 0x69, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x61, 0x72,
	0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6b, 0x61, 0x72, 0x6d, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x32, 0x86, 0x01, 0x0a, 0x05, 0x55, 0x72,
	0x62, 0x61, 0x6e, 0x12, 0x47, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x53, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x75,
	0x72, 0x62, 0x61, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05,
	0x57, 0x68, 0x6f, 0x69, 0x73, 0x12, 0x13, 0x2e, 0x75, 0x72, 0x62, 0x61, 0x6e, 0x2e, 0x57, 0x68,
	0x6f, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x72, 0x62,
	0x61, 0x6e, 0x2e, 0x57, 0x68, 0x6f, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x6f, 0x73, 0x63, 0x65, 0x6c, 0x79, 0x6e, 0x2f, 0x75, 0x72, 0x62, 0x61, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_urban_proto_rawDescOnce sync.Once
	file_urban_proto_rawDescData = file_urban_proto_rawDesc
)

func file_urban_proto_rawDescGZIP() []byte {
	file_urban_proto_rawDescOnce.Do(func() {
		file_urban_proto_rawDescData = protoimpl.X.CompressGZIP(file_urban_proto_rawDescData)
	})
	return file_urban_proto_rawDescData
}

var file_urban_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_urban_proto_goTypes = []interface{}{
	(*HNPost)(nil),                // 0: urban.HNPost
	(*ListResponse)(nil),          // 1: urban.ListResponse
	(*GetTopStoriesResponse)(nil), // 2: urban.GetTopStoriesResponse
	(*WhoisRequest)(nil),          // 3: urban.WhoisRequest
	(*WhoisResponse)(nil),         // 4: urban.WhoisResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_urban_proto_depIdxs = []int32{
	0, // 0: urban.ListResponse.posts:type_name -> urban.HNPost
	0, // 1: urban.GetTopStoriesResponse.posts:type_name -> urban.HNPost
	5, // 2: urban.WhoisResponse.joined:type_name -> google.protobuf.Timestamp
	6, // 3: urban.Urban.GetTopStories:input_type -> google.protobuf.Empty
	3, // 4: urban.Urban.Whois:input_type -> urban.WhoisRequest
	2, // 5: urban.Urban.GetTopStories:output_type -> urban.GetTopStoriesResponse
	4, // 6: urban.Urban.Whois:output_type -> urban.WhoisResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_urban_proto_init() }
func file_urban_proto_init() {
	if File_urban_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_urban_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HNPost); i {
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
		file_urban_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_urban_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopStoriesResponse); i {
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
		file_urban_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoisRequest); i {
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
		file_urban_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoisResponse); i {
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
			RawDescriptor: file_urban_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_urban_proto_goTypes,
		DependencyIndexes: file_urban_proto_depIdxs,
		MessageInfos:      file_urban_proto_msgTypes,
	}.Build()
	File_urban_proto = out.File
	file_urban_proto_rawDesc = nil
	file_urban_proto_goTypes = nil
	file_urban_proto_depIdxs = nil
}
