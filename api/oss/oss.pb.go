// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: oss-proto/oss.proto

package oss

import (
	_ "github.com/go-cinch/common/proto/params"
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PreSignedCategory int32

const (
	PreSignedCategory_GET  PreSignedCategory = 0
	PreSignedCategory_PUT  PreSignedCategory = 1
	PreSignedCategory_HEAD PreSignedCategory = 2
)

// Enum value maps for PreSignedCategory.
var (
	PreSignedCategory_name = map[int32]string{
		0: "GET",
		1: "PUT",
		2: "HEAD",
	}
	PreSignedCategory_value = map[string]int32{
		"GET":  0,
		"PUT":  1,
		"HEAD": 2,
	}
)

func (x PreSignedCategory) Enum() *PreSignedCategory {
	p := new(PreSignedCategory)
	*p = x
	return p
}

func (x PreSignedCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PreSignedCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_oss_proto_oss_proto_enumTypes[0].Descriptor()
}

func (PreSignedCategory) Type() protoreflect.EnumType {
	return &file_oss_proto_oss_proto_enumTypes[0]
}

func (x PreSignedCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PreSignedCategory.Descriptor instead.
func (PreSignedCategory) EnumDescriptor() ([]byte, []int) {
	return file_oss_proto_oss_proto_rawDescGZIP(), []int{0}
}

type PreSignedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category PreSignedCategory `protobuf:"varint,1,opt,name=category,proto3,enum=oss.v1.PreSignedCategory" json:"category,omitempty"`
	Name     string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PreSignedRequest) Reset() {
	*x = PreSignedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oss_proto_oss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreSignedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreSignedRequest) ProtoMessage() {}

func (x *PreSignedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oss_proto_oss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreSignedRequest.ProtoReflect.Descriptor instead.
func (*PreSignedRequest) Descriptor() ([]byte, []int) {
	return file_oss_proto_oss_proto_rawDescGZIP(), []int{0}
}

func (x *PreSignedRequest) GetCategory() PreSignedCategory {
	if x != nil {
		return x.Category
	}
	return PreSignedCategory_GET
}

func (x *PreSignedRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PreSignedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *PreSignedReply) Reset() {
	*x = PreSignedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oss_proto_oss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreSignedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreSignedReply) ProtoMessage() {}

func (x *PreSignedReply) ProtoReflect() protoreflect.Message {
	mi := &file_oss_proto_oss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreSignedReply.ProtoReflect.Descriptor instead.
func (*PreSignedReply) Descriptor() ([]byte, []int) {
	return file_oss_proto_oss_proto_rawDescGZIP(), []int{1}
}

func (x *PreSignedReply) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type OcrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *OcrRequest) Reset() {
	*x = OcrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oss_proto_oss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OcrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OcrRequest) ProtoMessage() {}

func (x *OcrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oss_proto_oss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OcrRequest.ProtoReflect.Descriptor instead.
func (*OcrRequest) Descriptor() ([]byte, []int) {
	return file_oss_proto_oss_proto_rawDescGZIP(), []int{2}
}

func (x *OcrRequest) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

type OcrDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Original     string      `protobuf:"bytes,1,opt,name=original,proto3" json:"original,omitempty"`
	Msg          string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	ParseLatency int64       `protobuf:"varint,4,opt,name=parseLatency,proto3" json:"parseLatency,omitempty"`
	DetLatency   int64       `protobuf:"varint,5,opt,name=detLatency,proto3" json:"detLatency,omitempty"`
	ClsLatency   int64       `protobuf:"varint,6,opt,name=clsLatency,proto3" json:"clsLatency,omitempty"`
	RecLatency   int64       `protobuf:"varint,7,opt,name=recLatency,proto3" json:"recLatency,omitempty"`
	Points       []*OcrPoint `protobuf:"bytes,8,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *OcrDetail) Reset() {
	*x = OcrDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oss_proto_oss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OcrDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OcrDetail) ProtoMessage() {}

func (x *OcrDetail) ProtoReflect() protoreflect.Message {
	mi := &file_oss_proto_oss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OcrDetail.ProtoReflect.Descriptor instead.
func (*OcrDetail) Descriptor() ([]byte, []int) {
	return file_oss_proto_oss_proto_rawDescGZIP(), []int{3}
}

func (x *OcrDetail) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

func (x *OcrDetail) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OcrDetail) GetParseLatency() int64 {
	if x != nil {
		return x.ParseLatency
	}
	return 0
}

func (x *OcrDetail) GetDetLatency() int64 {
	if x != nil {
		return x.DetLatency
	}
	return 0
}

func (x *OcrDetail) GetClsLatency() int64 {
	if x != nil {
		return x.ClsLatency
	}
	return 0
}

func (x *OcrDetail) GetRecLatency() int64 {
	if x != nil {
		return x.RecLatency
	}
	return 0
}

func (x *OcrDetail) GetPoints() []*OcrPoint {
	if x != nil {
		return x.Points
	}
	return nil
}

type OcrPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TextRegion string `protobuf:"bytes,1,opt,name=textRegion,proto3" json:"textRegion,omitempty"`
	Confidence string `protobuf:"bytes,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Text       string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *OcrPoint) Reset() {
	*x = OcrPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oss_proto_oss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OcrPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OcrPoint) ProtoMessage() {}

func (x *OcrPoint) ProtoReflect() protoreflect.Message {
	mi := &file_oss_proto_oss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OcrPoint.ProtoReflect.Descriptor instead.
func (*OcrPoint) Descriptor() ([]byte, []int) {
	return file_oss_proto_oss_proto_rawDescGZIP(), []int{4}
}

func (x *OcrPoint) GetTextRegion() string {
	if x != nil {
		return x.TextRegion
	}
	return ""
}

func (x *OcrPoint) GetConfidence() string {
	if x != nil {
		return x.Confidence
	}
	return ""
}

func (x *OcrPoint) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type OcrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List    []*OcrDetail `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Latency int64        `protobuf:"varint,2,opt,name=latency,proto3" json:"latency,omitempty"`
}

func (x *OcrReply) Reset() {
	*x = OcrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oss_proto_oss_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OcrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OcrReply) ProtoMessage() {}

func (x *OcrReply) ProtoReflect() protoreflect.Message {
	mi := &file_oss_proto_oss_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OcrReply.ProtoReflect.Descriptor instead.
func (*OcrReply) Descriptor() ([]byte, []int) {
	return file_oss_proto_oss_proto_rawDescGZIP(), []int{5}
}

func (x *OcrReply) GetList() []*OcrDetail {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *OcrReply) GetLatency() int64 {
	if x != nil {
		return x.Latency
	}
	return 0
}

var File_oss_proto_oss_proto protoreflect.FileDescriptor

var file_oss_proto_oss_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x73, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x69, 0x6e, 0x63, 0x68, 0x2f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5d, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x22, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x22, 0x20, 0x0a, 0x0a, 0x4f, 0x63, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xe7, 0x01, 0x0a, 0x09, 0x4f, 0x63, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x73, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x73, 0x65, 0x4c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x4c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x73, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6c, 0x73, 0x4c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x4c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x63, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x22, 0x5e, 0x0a, 0x08, 0x4f, 0x63, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x4b, 0x0a, 0x08, 0x4f, 0x63, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x73, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x63, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2a, 0x2f, 0x0a,
	0x11, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50,
	0x55, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10, 0x02, 0x32, 0x9a,
	0x01, 0x0a, 0x03, 0x4f, 0x73, 0x73, 0x12, 0x55, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x12, 0x18, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x6f, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a,
	0x22, 0x0b, 0x2f, 0x70, 0x72, 0x65, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x3c, 0x0a,
	0x03, 0x4f, 0x63, 0x72, 0x12, 0x12, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x63,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x63, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x09, 0x3a, 0x01, 0x2a, 0x22, 0x04, 0x2f, 0x6f, 0x63, 0x72, 0x42, 0x8c, 0x01, 0xba, 0x47,
	0x66, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x73, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x18, 0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x6f, 0x73, 0x73, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x63, 0x73, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e,
	0x30, 0x2a, 0x22, 0x3a, 0x20, 0x0a, 0x1e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x0e, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x2a, 0x06, 0x62,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x32, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x0a, 0x06, 0x6f, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x0a, 0x4f, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x0b, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x73, 0x73, 0x3b, 0x6f, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_oss_proto_oss_proto_rawDescOnce sync.Once
	file_oss_proto_oss_proto_rawDescData = file_oss_proto_oss_proto_rawDesc
)

func file_oss_proto_oss_proto_rawDescGZIP() []byte {
	file_oss_proto_oss_proto_rawDescOnce.Do(func() {
		file_oss_proto_oss_proto_rawDescData = protoimpl.X.CompressGZIP(file_oss_proto_oss_proto_rawDescData)
	})
	return file_oss_proto_oss_proto_rawDescData
}

var file_oss_proto_oss_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_oss_proto_oss_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_oss_proto_oss_proto_goTypes = []interface{}{
	(PreSignedCategory)(0),   // 0: oss.v1.PreSignedCategory
	(*PreSignedRequest)(nil), // 1: oss.v1.PreSignedRequest
	(*PreSignedReply)(nil),   // 2: oss.v1.PreSignedReply
	(*OcrRequest)(nil),       // 3: oss.v1.OcrRequest
	(*OcrDetail)(nil),        // 4: oss.v1.OcrDetail
	(*OcrPoint)(nil),         // 5: oss.v1.OcrPoint
	(*OcrReply)(nil),         // 6: oss.v1.OcrReply
}
var file_oss_proto_oss_proto_depIdxs = []int32{
	0, // 0: oss.v1.PreSignedRequest.category:type_name -> oss.v1.PreSignedCategory
	5, // 1: oss.v1.OcrDetail.points:type_name -> oss.v1.OcrPoint
	4, // 2: oss.v1.OcrReply.list:type_name -> oss.v1.OcrDetail
	1, // 3: oss.v1.Oss.PreSigned:input_type -> oss.v1.PreSignedRequest
	3, // 4: oss.v1.Oss.Ocr:input_type -> oss.v1.OcrRequest
	2, // 5: oss.v1.Oss.PreSigned:output_type -> oss.v1.PreSignedReply
	6, // 6: oss.v1.Oss.Ocr:output_type -> oss.v1.OcrReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_oss_proto_oss_proto_init() }
func file_oss_proto_oss_proto_init() {
	if File_oss_proto_oss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oss_proto_oss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreSignedRequest); i {
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
		file_oss_proto_oss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreSignedReply); i {
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
		file_oss_proto_oss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OcrRequest); i {
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
		file_oss_proto_oss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OcrDetail); i {
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
		file_oss_proto_oss_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OcrPoint); i {
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
		file_oss_proto_oss_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OcrReply); i {
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
			RawDescriptor: file_oss_proto_oss_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oss_proto_oss_proto_goTypes,
		DependencyIndexes: file_oss_proto_oss_proto_depIdxs,
		EnumInfos:         file_oss_proto_oss_proto_enumTypes,
		MessageInfos:      file_oss_proto_oss_proto_msgTypes,
	}.Build()
	File_oss_proto_oss_proto = out.File
	file_oss_proto_oss_proto_rawDesc = nil
	file_oss_proto_oss_proto_goTypes = nil
	file_oss_proto_oss_proto_depIdxs = nil
}
