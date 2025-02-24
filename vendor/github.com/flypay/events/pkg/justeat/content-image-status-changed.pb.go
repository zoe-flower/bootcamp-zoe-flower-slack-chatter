// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: justeat/content-image-status-changed.proto

package justeat

import (
	_ "github.com/flypay/events/pkg/flyt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContentImageStatusChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlTemplate      string           `protobuf:"bytes,1,opt,name=url_template,json=UrlTemplate,proto3" json:"url_template,omitempty"`
	RawImageUrl      string           `protobuf:"bytes,2,opt,name=raw_image_url,json=RawImageUrl,proto3" json:"raw_image_url,omitempty"`
	ImageContext     *structpb.Struct `protobuf:"bytes,3,opt,name=image_context,json=ImageContext,proto3" json:"image_context,omitempty"`
	ModerationStatus string           `protobuf:"bytes,4,opt,name=moderation_status,json=ModerationStatus,proto3" json:"moderation_status,omitempty"`
	// Deprecated: Do not use.
	RejectionReasons           []string                      `protobuf:"bytes,5,rep,name=rejection_reasons,json=RejectionReasons,proto3" json:"rejection_reasons,omitempty"`
	ModerationRejectionReasons []*ModerationRejectionReasons `protobuf:"bytes,6,rep,name=moderation_rejection_reasons,json=ModerationRejectionReasons,proto3" json:"moderation_rejection_reasons,omitempty"`
}

func (x *ContentImageStatusChanged) Reset() {
	*x = ContentImageStatusChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_content_image_status_changed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentImageStatusChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentImageStatusChanged) ProtoMessage() {}

func (x *ContentImageStatusChanged) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_content_image_status_changed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentImageStatusChanged.ProtoReflect.Descriptor instead.
func (*ContentImageStatusChanged) Descriptor() ([]byte, []int) {
	return file_justeat_content_image_status_changed_proto_rawDescGZIP(), []int{0}
}

func (x *ContentImageStatusChanged) GetUrlTemplate() string {
	if x != nil {
		return x.UrlTemplate
	}
	return ""
}

func (x *ContentImageStatusChanged) GetRawImageUrl() string {
	if x != nil {
		return x.RawImageUrl
	}
	return ""
}

func (x *ContentImageStatusChanged) GetImageContext() *structpb.Struct {
	if x != nil {
		return x.ImageContext
	}
	return nil
}

func (x *ContentImageStatusChanged) GetModerationStatus() string {
	if x != nil {
		return x.ModerationStatus
	}
	return ""
}

// Deprecated: Do not use.
func (x *ContentImageStatusChanged) GetRejectionReasons() []string {
	if x != nil {
		return x.RejectionReasons
	}
	return nil
}

func (x *ContentImageStatusChanged) GetModerationRejectionReasons() []*ModerationRejectionReasons {
	if x != nil {
		return x.ModerationRejectionReasons
	}
	return nil
}

type ModerationRejectionReasons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,json=Id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,json=Value,proto3" json:"value,omitempty"`
}

func (x *ModerationRejectionReasons) Reset() {
	*x = ModerationRejectionReasons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_content_image_status_changed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModerationRejectionReasons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModerationRejectionReasons) ProtoMessage() {}

func (x *ModerationRejectionReasons) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_content_image_status_changed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModerationRejectionReasons.ProtoReflect.Descriptor instead.
func (*ModerationRejectionReasons) Descriptor() ([]byte, []int) {
	return file_justeat_content_image_status_changed_proto_rawDescGZIP(), []int{1}
}

func (x *ModerationRejectionReasons) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ModerationRejectionReasons) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_justeat_content_image_status_changed_proto protoreflect.FileDescriptor

var file_justeat_content_image_status_changed_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2d, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6a, 0x75,
	0x73, 0x74, 0x65, 0x61, 0x74, 0x1a, 0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6a, 0x75, 0x73, 0x74,
	0x65, 0x61, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa4, 0x03, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x72, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x61, 0x77, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x61, 0x77, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x3c, 0x0a, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2f, 0x0a, 0x11, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x10, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x73, 0x12, 0x65, 0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65,
	0x61, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x1a, 0x4d,
	0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x3a, 0x3d, 0x82, 0xb5, 0x18, 0x1c, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0xa2, 0xbb, 0x18, 0x19, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x22, 0x42, 0x0a, 0x1a, 0x4d, 0x6f, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x8f, 0x01, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x42, 0x1e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x70, 0x61,
	0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6a, 0x75, 0x73,
	0x74, 0x65, 0x61, 0x74, 0xa2, 0x02, 0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4a, 0x75, 0x73,
	0x74, 0x65, 0x61, 0x74, 0xca, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0xe2, 0x02,
	0x13, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_justeat_content_image_status_changed_proto_rawDescOnce sync.Once
	file_justeat_content_image_status_changed_proto_rawDescData = file_justeat_content_image_status_changed_proto_rawDesc
)

func file_justeat_content_image_status_changed_proto_rawDescGZIP() []byte {
	file_justeat_content_image_status_changed_proto_rawDescOnce.Do(func() {
		file_justeat_content_image_status_changed_proto_rawDescData = protoimpl.X.CompressGZIP(file_justeat_content_image_status_changed_proto_rawDescData)
	})
	return file_justeat_content_image_status_changed_proto_rawDescData
}

var file_justeat_content_image_status_changed_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_justeat_content_image_status_changed_proto_goTypes = []interface{}{
	(*ContentImageStatusChanged)(nil),  // 0: justeat.ContentImageStatusChanged
	(*ModerationRejectionReasons)(nil), // 1: justeat.ModerationRejectionReasons
	(*structpb.Struct)(nil),            // 2: google.protobuf.Struct
}
var file_justeat_content_image_status_changed_proto_depIdxs = []int32{
	2, // 0: justeat.ContentImageStatusChanged.image_context:type_name -> google.protobuf.Struct
	1, // 1: justeat.ContentImageStatusChanged.moderation_rejection_reasons:type_name -> justeat.ModerationRejectionReasons
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_justeat_content_image_status_changed_proto_init() }
func file_justeat_content_image_status_changed_proto_init() {
	if File_justeat_content_image_status_changed_proto != nil {
		return
	}
	file_justeat_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_justeat_content_image_status_changed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentImageStatusChanged); i {
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
		file_justeat_content_image_status_changed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModerationRejectionReasons); i {
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
			RawDescriptor: file_justeat_content_image_status_changed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_justeat_content_image_status_changed_proto_goTypes,
		DependencyIndexes: file_justeat_content_image_status_changed_proto_depIdxs,
		MessageInfos:      file_justeat_content_image_status_changed_proto_msgTypes,
	}.Build()
	File_justeat_content_image_status_changed_proto = out.File
	file_justeat_content_image_status_changed_proto_rawDesc = nil
	file_justeat_content_image_status_changed_proto_goTypes = nil
	file_justeat_content_image_status_changed_proto_depIdxs = nil
}
