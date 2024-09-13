// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: api/proto/market.proto

package api

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

type Statistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Views  int32   `protobuf:"varint,2,opt,name=views,proto3" json:"views,omitempty"`
	Clicks int32   `protobuf:"varint,3,opt,name=clicks,proto3" json:"clicks,omitempty"`
	Cost   float32 `protobuf:"fixed32,4,opt,name=cost,proto3" json:"cost,omitempty"`
	Cpc    float32 `protobuf:"fixed32,5,opt,name=cpc,proto3" json:"cpc,omitempty"`
	Cpm    float32 `protobuf:"fixed32,6,opt,name=cpm,proto3" json:"cpm,omitempty"`
}

func (x *Statistics) Reset() {
	*x = Statistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_market_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statistics) ProtoMessage() {}

func (x *Statistics) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_market_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statistics.ProtoReflect.Descriptor instead.
func (*Statistics) Descriptor() ([]byte, []int) {
	return file_api_proto_market_proto_rawDescGZIP(), []int{0}
}

func (x *Statistics) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Statistics) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *Statistics) GetClicks() int32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *Statistics) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Statistics) GetCpc() float32 {
	if x != nil {
		return x.Cpc
	}
	return 0
}

func (x *Statistics) GetCpm() float32 {
	if x != nil {
		return x.Cpm
	}
	return 0
}

type ResetStatisticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetStatisticsRequest) Reset() {
	*x = ResetStatisticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_market_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetStatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetStatisticsRequest) ProtoMessage() {}

func (x *ResetStatisticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_market_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetStatisticsRequest.ProtoReflect.Descriptor instead.
func (*ResetStatisticsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_market_proto_rawDescGZIP(), []int{1}
}

type ResetStatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResetStatisticsResponse) Reset() {
	*x = ResetStatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_market_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetStatisticsResponse) ProtoMessage() {}

func (x *ResetStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_market_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetStatisticsResponse.ProtoReflect.Descriptor instead.
func (*ResetStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_market_proto_rawDescGZIP(), []int{2}
}

func (x *ResetStatisticsResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type SaveStatisticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Views  int32   `protobuf:"varint,2,opt,name=views,proto3" json:"views,omitempty"`
	Clicks int32   `protobuf:"varint,3,opt,name=clicks,proto3" json:"clicks,omitempty"`
	Cost   float32 `protobuf:"fixed32,4,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *SaveStatisticsRequest) Reset() {
	*x = SaveStatisticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_market_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveStatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveStatisticsRequest) ProtoMessage() {}

func (x *SaveStatisticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_market_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveStatisticsRequest.ProtoReflect.Descriptor instead.
func (*SaveStatisticsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_market_proto_rawDescGZIP(), []int{3}
}

func (x *SaveStatisticsRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *SaveStatisticsRequest) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *SaveStatisticsRequest) GetClicks() int32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *SaveStatisticsRequest) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type SaveStatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SaveStatisticsResponse) Reset() {
	*x = SaveStatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_market_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveStatisticsResponse) ProtoMessage() {}

func (x *SaveStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_market_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveStatisticsResponse.ProtoReflect.Descriptor instead.
func (*SaveStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_market_proto_rawDescGZIP(), []int{4}
}

func (x *SaveStatisticsResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ShowStatisticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *ShowStatisticsRequest) Reset() {
	*x = ShowStatisticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_market_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowStatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowStatisticsRequest) ProtoMessage() {}

func (x *ShowStatisticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_market_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowStatisticsRequest.ProtoReflect.Descriptor instead.
func (*ShowStatisticsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_market_proto_rawDescGZIP(), []int{5}
}

func (x *ShowStatisticsRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ShowStatisticsRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type ShowStatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statistics []*Statistics `protobuf:"bytes,1,rep,name=statistics,proto3" json:"statistics,omitempty"`
}

func (x *ShowStatisticsResponse) Reset() {
	*x = ShowStatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_market_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowStatisticsResponse) ProtoMessage() {}

func (x *ShowStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_market_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowStatisticsResponse.ProtoReflect.Descriptor instead.
func (*ShowStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_market_proto_rawDescGZIP(), []int{6}
}

func (x *ShowStatisticsResponse) GetStatistics() []*Statistics {
	if x != nil {
		return x.Statistics
	}
	return nil
}

var File_api_proto_market_proto protoreflect.FileDescriptor

var file_api_proto_market_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x86, 0x01,
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x63, 0x70, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x63, 0x70, 0x6d, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x31, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x6d, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x22, 0x30, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x3b, 0x0a, 0x15, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x6f, 0x22, 0x49, 0x0a, 0x16, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x32, 0xf2, 0x01, 0x0a,
	0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x4e, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f,
	0x77, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_market_proto_rawDescOnce sync.Once
	file_api_proto_market_proto_rawDescData = file_api_proto_market_proto_rawDesc
)

func file_api_proto_market_proto_rawDescGZIP() []byte {
	file_api_proto_market_proto_rawDescOnce.Do(func() {
		file_api_proto_market_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_market_proto_rawDescData)
	})
	return file_api_proto_market_proto_rawDescData
}

var file_api_proto_market_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_market_proto_goTypes = []interface{}{
	(*Statistics)(nil),              // 0: api.Statistics
	(*ResetStatisticsRequest)(nil),  // 1: api.ResetStatisticsRequest
	(*ResetStatisticsResponse)(nil), // 2: api.ResetStatisticsResponse
	(*SaveStatisticsRequest)(nil),   // 3: api.SaveStatisticsRequest
	(*SaveStatisticsResponse)(nil),  // 4: api.SaveStatisticsResponse
	(*ShowStatisticsRequest)(nil),   // 5: api.ShowStatisticsRequest
	(*ShowStatisticsResponse)(nil),  // 6: api.ShowStatisticsResponse
}
var file_api_proto_market_proto_depIdxs = []int32{
	0, // 0: api.ShowStatisticsResponse.statistics:type_name -> api.Statistics
	1, // 1: api.Market.ResetStatistics:input_type -> api.ResetStatisticsRequest
	3, // 2: api.Market.SaveStatistics:input_type -> api.SaveStatisticsRequest
	5, // 3: api.Market.ShowStatistics:input_type -> api.ShowStatisticsRequest
	2, // 4: api.Market.ResetStatistics:output_type -> api.ResetStatisticsResponse
	4, // 5: api.Market.SaveStatistics:output_type -> api.SaveStatisticsResponse
	6, // 6: api.Market.ShowStatistics:output_type -> api.ShowStatisticsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_market_proto_init() }
func file_api_proto_market_proto_init() {
	if File_api_proto_market_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_market_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statistics); i {
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
		file_api_proto_market_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetStatisticsRequest); i {
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
		file_api_proto_market_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetStatisticsResponse); i {
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
		file_api_proto_market_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveStatisticsRequest); i {
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
		file_api_proto_market_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveStatisticsResponse); i {
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
		file_api_proto_market_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowStatisticsRequest); i {
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
		file_api_proto_market_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowStatisticsResponse); i {
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
			RawDescriptor: file_api_proto_market_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_market_proto_goTypes,
		DependencyIndexes: file_api_proto_market_proto_depIdxs,
		MessageInfos:      file_api_proto_market_proto_msgTypes,
	}.Build()
	File_api_proto_market_proto = out.File
	file_api_proto_market_proto_rawDesc = nil
	file_api_proto_market_proto_goTypes = nil
	file_api_proto_market_proto_depIdxs = nil
}
