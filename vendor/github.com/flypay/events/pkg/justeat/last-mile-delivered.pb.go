// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: justeat/last-mile-delivered.proto

package justeat

import (
	_ "github.com/flypay/events/pkg/flyt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// LastMileDelivered is emitted by JustEat when a meal was delivered.
type LastMileDelivered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If present, use as order ID.
	PartnerOrderId string `protobuf:"bytes,1,opt,name=partner_order_id,json=PartnerOrderId,proto3" json:"partner_order_id,omitempty"`
	// Use this as order ID if PartnerOrderId is not present.
	OrderReference            string                 `protobuf:"bytes,2,opt,name=order_reference,json=OrderReference,proto3" json:"order_reference,omitempty"`
	FriendlyOrderReference    string                 `protobuf:"bytes,3,opt,name=friendly_order_reference,json=FriendlyOrderReference,proto3" json:"friendly_order_reference,omitempty"`
	RestaurantId              int32                  `protobuf:"varint,4,opt,name=restaurant_id,json=RestaurantId,proto3" json:"restaurant_id,omitempty"`
	DeliveryPartnerName       string                 `protobuf:"bytes,5,opt,name=delivery_partner_name,json=DeliveryPartnerName,proto3" json:"delivery_partner_name,omitempty"`
	DeliveryPartnerId         *int32                 `protobuf:"varint,6,opt,name=delivery_partner_id,json=DeliveryPartnerId,proto3,oneof" json:"delivery_partner_id,omitempty"`
	DeliveryPartnerPickupTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=delivery_partner_pickup_time,json=DeliveryPartnerPickupTime,proto3,oneof" json:"delivery_partner_pickup_time,omitempty"`
	DeliveryAddress           *Location              `protobuf:"bytes,8,opt,name=delivery_address,json=DeliveryAddress,proto3" json:"delivery_address,omitempty"`
	RestaurantAddress         *Location              `protobuf:"bytes,9,opt,name=restaurant_address,json=RestaurantAddress,proto3" json:"restaurant_address,omitempty"`
	OriginalEventTimestamp    *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=original_event_timestamp,json=OriginalEventTimestamp,proto3" json:"original_event_timestamp,omitempty"`
	Confidence                *Confidence            `protobuf:"bytes,11,opt,name=confidence,json=Confidence,proto3" json:"confidence,omitempty"`
	TimeStamp                 *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=time_stamp,json=TimeStamp,proto3" json:"time_stamp,omitempty"`
	RaisingComponent          string                 `protobuf:"bytes,13,opt,name=raising_component,json=RaisingComponent,proto3" json:"raising_component,omitempty"`
	Tenant                    string                 `protobuf:"bytes,14,opt,name=tenant,json=Tenant,proto3" json:"tenant,omitempty"`
	Conversation              string                 `protobuf:"bytes,15,opt,name=conversation,json=Conversation,proto3" json:"conversation,omitempty"`
}

func (x *LastMileDelivered) Reset() {
	*x = LastMileDelivered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_last_mile_delivered_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastMileDelivered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastMileDelivered) ProtoMessage() {}

func (x *LastMileDelivered) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_last_mile_delivered_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastMileDelivered.ProtoReflect.Descriptor instead.
func (*LastMileDelivered) Descriptor() ([]byte, []int) {
	return file_justeat_last_mile_delivered_proto_rawDescGZIP(), []int{0}
}

func (x *LastMileDelivered) GetPartnerOrderId() string {
	if x != nil {
		return x.PartnerOrderId
	}
	return ""
}

func (x *LastMileDelivered) GetOrderReference() string {
	if x != nil {
		return x.OrderReference
	}
	return ""
}

func (x *LastMileDelivered) GetFriendlyOrderReference() string {
	if x != nil {
		return x.FriendlyOrderReference
	}
	return ""
}

func (x *LastMileDelivered) GetRestaurantId() int32 {
	if x != nil {
		return x.RestaurantId
	}
	return 0
}

func (x *LastMileDelivered) GetDeliveryPartnerName() string {
	if x != nil {
		return x.DeliveryPartnerName
	}
	return ""
}

func (x *LastMileDelivered) GetDeliveryPartnerId() int32 {
	if x != nil && x.DeliveryPartnerId != nil {
		return *x.DeliveryPartnerId
	}
	return 0
}

func (x *LastMileDelivered) GetDeliveryPartnerPickupTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveryPartnerPickupTime
	}
	return nil
}

func (x *LastMileDelivered) GetDeliveryAddress() *Location {
	if x != nil {
		return x.DeliveryAddress
	}
	return nil
}

func (x *LastMileDelivered) GetRestaurantAddress() *Location {
	if x != nil {
		return x.RestaurantAddress
	}
	return nil
}

func (x *LastMileDelivered) GetOriginalEventTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.OriginalEventTimestamp
	}
	return nil
}

func (x *LastMileDelivered) GetConfidence() *Confidence {
	if x != nil {
		return x.Confidence
	}
	return nil
}

func (x *LastMileDelivered) GetTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

func (x *LastMileDelivered) GetRaisingComponent() string {
	if x != nil {
		return x.RaisingComponent
	}
	return ""
}

func (x *LastMileDelivered) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LastMileDelivered) GetConversation() string {
	if x != nil {
		return x.Conversation
	}
	return ""
}

var File_justeat_last_mile_delivered_proto protoreflect.FileDescriptor

var file_justeat_last_mile_delivered_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2f, 0x6c, 0x61, 0x73, 0x74, 0x2d, 0x6d,
	0x69, 0x6c, 0x65, 0x2d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x1a, 0x15, 0x66, 0x6c,
	0x79, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2f, 0x6c, 0x61, 0x73,
	0x74, 0x2d, 0x6d, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6a, 0x75,
	0x73, 0x74, 0x65, 0x61, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x07, 0x0a, 0x11, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x18, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x16, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x33, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x11, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x60, 0x0a, 0x1c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x19, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x54, 0x0a, 0x18, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x16, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x33, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a,
	0x11, 0x72, 0x61, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x61, 0x69, 0x73, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x34, 0x82, 0xb5, 0x18, 0x1b, 0x6a, 0x75, 0x73, 0x74,
	0x65, 0x61, 0x74, 0x2e, 0x6c, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x6c, 0x65, 0x2e, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0xa2, 0xbb, 0x18, 0x11, 0x4c, 0x61, 0x73, 0x74, 0x4d,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x42, 0x16, 0x0a, 0x14,
	0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x42, 0x1f, 0x0a, 0x1d, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x87, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x75,
	0x73, 0x74, 0x65, 0x61, 0x74, 0x42, 0x16, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x70,
	0x61, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6a, 0x75,
	0x73, 0x74, 0x65, 0x61, 0x74, 0xa2, 0x02, 0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4a, 0x75,
	0x73, 0x74, 0x65, 0x61, 0x74, 0xca, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0xe2,
	0x02, 0x13, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_justeat_last_mile_delivered_proto_rawDescOnce sync.Once
	file_justeat_last_mile_delivered_proto_rawDescData = file_justeat_last_mile_delivered_proto_rawDesc
)

func file_justeat_last_mile_delivered_proto_rawDescGZIP() []byte {
	file_justeat_last_mile_delivered_proto_rawDescOnce.Do(func() {
		file_justeat_last_mile_delivered_proto_rawDescData = protoimpl.X.CompressGZIP(file_justeat_last_mile_delivered_proto_rawDescData)
	})
	return file_justeat_last_mile_delivered_proto_rawDescData
}

var file_justeat_last_mile_delivered_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_justeat_last_mile_delivered_proto_goTypes = []interface{}{
	(*LastMileDelivered)(nil),     // 0: justeat.LastMileDelivered
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Location)(nil),              // 2: justeat.Location
	(*Confidence)(nil),            // 3: justeat.Confidence
}
var file_justeat_last_mile_delivered_proto_depIdxs = []int32{
	1, // 0: justeat.LastMileDelivered.delivery_partner_pickup_time:type_name -> google.protobuf.Timestamp
	2, // 1: justeat.LastMileDelivered.delivery_address:type_name -> justeat.Location
	2, // 2: justeat.LastMileDelivered.restaurant_address:type_name -> justeat.Location
	1, // 3: justeat.LastMileDelivered.original_event_timestamp:type_name -> google.protobuf.Timestamp
	3, // 4: justeat.LastMileDelivered.confidence:type_name -> justeat.Confidence
	1, // 5: justeat.LastMileDelivered.time_stamp:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_justeat_last_mile_delivered_proto_init() }
func file_justeat_last_mile_delivered_proto_init() {
	if File_justeat_last_mile_delivered_proto != nil {
		return
	}
	file_justeat_last_mile_proto_init()
	file_justeat_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_justeat_last_mile_delivered_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastMileDelivered); i {
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
	file_justeat_last_mile_delivered_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_justeat_last_mile_delivered_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_justeat_last_mile_delivered_proto_goTypes,
		DependencyIndexes: file_justeat_last_mile_delivered_proto_depIdxs,
		MessageInfos:      file_justeat_last_mile_delivered_proto_msgTypes,
	}.Build()
	File_justeat_last_mile_delivered_proto = out.File
	file_justeat_last_mile_delivered_proto_rawDesc = nil
	file_justeat_last_mile_delivered_proto_goTypes = nil
	file_justeat_last_mile_delivered_proto_depIdxs = nil
}
