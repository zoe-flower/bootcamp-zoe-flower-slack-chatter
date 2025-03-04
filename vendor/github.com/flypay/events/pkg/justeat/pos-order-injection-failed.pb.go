// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: justeat/pos-order-injection-failed.proto

package justeat

import (
	_ "github.com/flypay/events/pkg/flyt"
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

type PosOrderInjectionFailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     string `protobuf:"bytes,1,opt,name=order_id,json=OrderId,proto3" json:"order_id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,json=Description,proto3" json:"description,omitempty"`
	ErrorCode   string `protobuf:"bytes,3,opt,name=error_code,json=ErrorCode,proto3" json:"error_code,omitempty"`
	// e.g 2021-05-05T13:01:49.4976625Z
	TimeStamp        string `protobuf:"bytes,4,opt,name=time_stamp,json=TimeStamp,proto3" json:"time_stamp,omitempty"`
	RaisingComponent string `protobuf:"bytes,5,opt,name=raising_component,json=RaisingComponent,proto3" json:"raising_component,omitempty"`
	Tenant           string `protobuf:"bytes,6,opt,name=tenant,json=Tenant,proto3" json:"tenant,omitempty"`
	Conversation     string `protobuf:"bytes,7,opt,name=conversation,json=Conversation,proto3" json:"conversation,omitempty"`
}

func (x *PosOrderInjectionFailed) Reset() {
	*x = PosOrderInjectionFailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_pos_order_injection_failed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosOrderInjectionFailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosOrderInjectionFailed) ProtoMessage() {}

func (x *PosOrderInjectionFailed) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_pos_order_injection_failed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosOrderInjectionFailed.ProtoReflect.Descriptor instead.
func (*PosOrderInjectionFailed) Descriptor() ([]byte, []int) {
	return file_justeat_pos_order_injection_failed_proto_rawDescGZIP(), []int{0}
}

func (x *PosOrderInjectionFailed) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PosOrderInjectionFailed) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PosOrderInjectionFailed) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *PosOrderInjectionFailed) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *PosOrderInjectionFailed) GetRaisingComponent() string {
	if x != nil {
		return x.RaisingComponent
	}
	return ""
}

func (x *PosOrderInjectionFailed) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *PosOrderInjectionFailed) GetConversation() string {
	if x != nil {
		return x.Conversation
	}
	return ""
}

var File_justeat_pos_order_injection_failed_proto protoreflect.FileDescriptor

var file_justeat_pos_order_injection_failed_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2f, 0x70, 0x6f, 0x73, 0x2d, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6a, 0x75, 0x73, 0x74,
	0x65, 0x61, 0x74, 0x1a, 0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6a, 0x75, 0x73, 0x74,
	0x65, 0x61, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb8, 0x02, 0x0a, 0x17, 0x50, 0x6f, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x61, 0x69, 0x73,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x61, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x39, 0x82, 0xb5, 0x18, 0x1a, 0x70, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x66, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0xa2, 0xbb, 0x18, 0x17, 0x50, 0x6f, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x42, 0x8d, 0x01, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x42, 0x1c, 0x50, 0x6f,
	0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x70, 0x61, 0x79, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x65,
	0x61, 0x74, 0xa2, 0x02, 0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65,
	0x61, 0x74, 0xca, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0xe2, 0x02, 0x13, 0x4a,
	0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_justeat_pos_order_injection_failed_proto_rawDescOnce sync.Once
	file_justeat_pos_order_injection_failed_proto_rawDescData = file_justeat_pos_order_injection_failed_proto_rawDesc
)

func file_justeat_pos_order_injection_failed_proto_rawDescGZIP() []byte {
	file_justeat_pos_order_injection_failed_proto_rawDescOnce.Do(func() {
		file_justeat_pos_order_injection_failed_proto_rawDescData = protoimpl.X.CompressGZIP(file_justeat_pos_order_injection_failed_proto_rawDescData)
	})
	return file_justeat_pos_order_injection_failed_proto_rawDescData
}

var file_justeat_pos_order_injection_failed_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_justeat_pos_order_injection_failed_proto_goTypes = []interface{}{
	(*PosOrderInjectionFailed)(nil), // 0: justeat.PosOrderInjectionFailed
}
var file_justeat_pos_order_injection_failed_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_justeat_pos_order_injection_failed_proto_init() }
func file_justeat_pos_order_injection_failed_proto_init() {
	if File_justeat_pos_order_injection_failed_proto != nil {
		return
	}
	file_justeat_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_justeat_pos_order_injection_failed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosOrderInjectionFailed); i {
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
			RawDescriptor: file_justeat_pos_order_injection_failed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_justeat_pos_order_injection_failed_proto_goTypes,
		DependencyIndexes: file_justeat_pos_order_injection_failed_proto_depIdxs,
		MessageInfos:      file_justeat_pos_order_injection_failed_proto_msgTypes,
	}.Build()
	File_justeat_pos_order_injection_failed_proto = out.File
	file_justeat_pos_order_injection_failed_proto_rawDesc = nil
	file_justeat_pos_order_injection_failed_proto_goTypes = nil
	file_justeat_pos_order_injection_failed_proto_depIdxs = nil
}
