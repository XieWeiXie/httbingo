// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/v1/api.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 业务状态码
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 业务描述
	Data    *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`       // 业务数据
}

func (x *ResponseReply) Reset() {
	*x = ResponseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseReply) ProtoMessage() {}

func (x *ResponseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseReply.ProtoReflect.Descriptor instead.
func (*ResponseReply) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseReply) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type HttpMethodGetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 主键
}

func (x *HttpMethodGetReq) Reset() {
	*x = HttpMethodGetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMethodGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMethodGetReq) ProtoMessage() {}

func (x *HttpMethodGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMethodGetReq.ProtoReflect.Descriptor instead.
func (*HttpMethodGetReq) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *HttpMethodGetReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type HttpMethodGetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    // 主键
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // 数据
}

func (x *HttpMethodGetReply) Reset() {
	*x = HttpMethodGetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMethodGetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMethodGetReply) ProtoMessage() {}

func (x *HttpMethodGetReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMethodGetReply.ProtoReflect.Descriptor instead.
func (*HttpMethodGetReply) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *HttpMethodGetReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HttpMethodGetReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HttpMethodPutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    // 主键
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // 修改项
}

func (x *HttpMethodPutReq) Reset() {
	*x = HttpMethodPutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMethodPutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMethodPutReq) ProtoMessage() {}

func (x *HttpMethodPutReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMethodPutReq.ProtoReflect.Descriptor instead.
func (*HttpMethodPutReq) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *HttpMethodPutReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HttpMethodPutReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HttpMethodPutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HttpMethodPutReply) Reset() {
	*x = HttpMethodPutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMethodPutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMethodPutReply) ProtoMessage() {}

func (x *HttpMethodPutReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMethodPutReply.ProtoReflect.Descriptor instead.
func (*HttpMethodPutReply) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{4}
}

func (x *HttpMethodPutReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HttpMethodPutReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HttpMethodDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *HttpMethodDeleteReq) Reset() {
	*x = HttpMethodDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMethodDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMethodDeleteReq) ProtoMessage() {}

func (x *HttpMethodDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMethodDeleteReq.ProtoReflect.Descriptor instead.
func (*HttpMethodDeleteReq) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *HttpMethodDeleteReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type HttpMethodDeleteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HttpMethodDeleteReply) Reset() {
	*x = HttpMethodDeleteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMethodDeleteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMethodDeleteReply) ProtoMessage() {}

func (x *HttpMethodDeleteReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMethodDeleteReply.ProtoReflect.Descriptor instead.
func (*HttpMethodDeleteReply) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{6}
}

type HttpMethodPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HttpMethodPostReq) Reset() {
	*x = HttpMethodPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMethodPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMethodPostReq) ProtoMessage() {}

func (x *HttpMethodPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMethodPostReq.ProtoReflect.Descriptor instead.
func (*HttpMethodPostReq) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{7}
}

func (x *HttpMethodPostReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HttpMethodPostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HttpMethodPostReply) Reset() {
	*x = HttpMethodPostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMethodPostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMethodPostReply) ProtoMessage() {}

func (x *HttpMethodPostReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMethodPostReply.ProtoReflect.Descriptor instead.
func (*HttpMethodPostReply) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{8}
}

func (x *HttpMethodPostReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HttpMethodPostReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StatusCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *StatusCodeReq) Reset() {
	*x = StatusCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusCodeReq) ProtoMessage() {}

func (x *StatusCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusCodeReq.ProtoReflect.Descriptor instead.
func (*StatusCodeReq) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{9}
}

func (x *StatusCodeReq) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type StatusCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *StatusCodeReply) Reset() {
	*x = StatusCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusCodeReply) ProtoMessage() {}

func (x *StatusCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusCodeReply.ProtoReflect.Descriptor instead.
func (*StatusCodeReply) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{10}
}

func (x *StatusCodeReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type HeadersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeadersReq) Reset() {
	*x = HeadersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeadersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeadersReq) ProtoMessage() {}

func (x *HeadersReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeadersReq.ProtoReflect.Descriptor instead.
func (*HeadersReq) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{11}
}

type HeadersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers map[string]*anypb.Any `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HeadersReply) Reset() {
	*x = HeadersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeadersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeadersReply) ProtoMessage() {}

func (x *HeadersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeadersReply.ProtoReflect.Descriptor instead.
func (*HeadersReply) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{12}
}

func (x *HeadersReply) GetHeaders() map[string]*anypb.Any {
	if x != nil {
		return x.Headers
	}
	return nil
}

type IPReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IPReq) Reset() {
	*x = IPReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPReq) ProtoMessage() {}

func (x *IPReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPReq.ProtoReflect.Descriptor instead.
func (*IPReq) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{13}
}

type IPReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *IPReply) Reset() {
	*x = IPReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPReply) ProtoMessage() {}

func (x *IPReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPReply.ProtoReflect.Descriptor instead.
func (*IPReply) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{14}
}

func (x *IPReply) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

type UserAgentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserAgentReq) Reset() {
	*x = UserAgentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAgentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAgentReq) ProtoMessage() {}

func (x *UserAgentReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAgentReq.ProtoReflect.Descriptor instead.
func (*UserAgentReq) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{15}
}

type UserAgentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAgent string `protobuf:"bytes,1,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
}

func (x *UserAgentReply) Reset() {
	*x = UserAgentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAgentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAgentReply) ProtoMessage() {}

func (x *UserAgentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAgentReply.ProtoReflect.Descriptor instead.
func (*UserAgentReply) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{16}
}

func (x *UserAgentReply) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

var File_api_v1_api_proto protoreflect.FileDescriptor

var file_api_v1_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x10, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x48, 0x74, 0x74, 0x70,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x36, 0x0a, 0x10, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x48, 0x74,
	0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x50, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x48,
	0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x27, 0x0a, 0x11, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a,
	0x13, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x25, 0x0a,
	0x0f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x22, 0xa3, 0x01, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69, 0x6e, 0x67, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x07, 0x0a, 0x05, 0x49, 0x50, 0x52, 0x65,
	0x71, 0x22, 0x21, 0x0a, 0x07, 0x49, 0x50, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x22, 0x0e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x22, 0x2f, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x32, 0xd8, 0x04, 0x0a, 0x0e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x48, 0x74, 0x74, 0x70,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x0d, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x50, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69,
	0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69,
	0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x52, 0x0a, 0x10, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62,
	0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x0e, 0x48, 0x74, 0x74, 0x70,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69, 0x6e,
	0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x40, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69, 0x6e, 0x67,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x36, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x13, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62,
	0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69,
	0x6e, 0x67, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69, 0x6e, 0x67, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58,
	0x69, 0x65, 0x57, 0x65, 0x69, 0x58, 0x69, 0x65, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x69, 0x6e,
	0x67, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_api_proto_rawDescOnce sync.Once
	file_api_v1_api_proto_rawDescData = file_api_v1_api_proto_rawDesc
)

func file_api_v1_api_proto_rawDescGZIP() []byte {
	file_api_v1_api_proto_rawDescOnce.Do(func() {
		file_api_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_api_proto_rawDescData)
	})
	return file_api_v1_api_proto_rawDescData
}

var file_api_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_api_v1_api_proto_goTypes = []interface{}{
	(*ResponseReply)(nil),         // 0: httpbingo.v1.ResponseReply
	(*HttpMethodGetReq)(nil),      // 1: httpbingo.v1.HttpMethodGetReq
	(*HttpMethodGetReply)(nil),    // 2: httpbingo.v1.HttpMethodGetReply
	(*HttpMethodPutReq)(nil),      // 3: httpbingo.v1.HttpMethodPutReq
	(*HttpMethodPutReply)(nil),    // 4: httpbingo.v1.HttpMethodPutReply
	(*HttpMethodDeleteReq)(nil),   // 5: httpbingo.v1.HttpMethodDeleteReq
	(*HttpMethodDeleteReply)(nil), // 6: httpbingo.v1.HttpMethodDeleteReply
	(*HttpMethodPostReq)(nil),     // 7: httpbingo.v1.HttpMethodPostReq
	(*HttpMethodPostReply)(nil),   // 8: httpbingo.v1.HttpMethodPostReply
	(*StatusCodeReq)(nil),         // 9: httpbingo.v1.StatusCodeReq
	(*StatusCodeReply)(nil),       // 10: httpbingo.v1.StatusCodeReply
	(*HeadersReq)(nil),            // 11: httpbingo.v1.HeadersReq
	(*HeadersReply)(nil),          // 12: httpbingo.v1.HeadersReply
	(*IPReq)(nil),                 // 13: httpbingo.v1.IPReq
	(*IPReply)(nil),               // 14: httpbingo.v1.IPReply
	(*UserAgentReq)(nil),          // 15: httpbingo.v1.UserAgentReq
	(*UserAgentReply)(nil),        // 16: httpbingo.v1.UserAgentReply
	nil,                           // 17: httpbingo.v1.HeadersReply.HeadersEntry
	(*anypb.Any)(nil),             // 18: google.protobuf.Any
}
var file_api_v1_api_proto_depIdxs = []int32{
	18, // 0: httpbingo.v1.ResponseReply.data:type_name -> google.protobuf.Any
	17, // 1: httpbingo.v1.HeadersReply.headers:type_name -> httpbingo.v1.HeadersReply.HeadersEntry
	18, // 2: httpbingo.v1.HeadersReply.HeadersEntry.value:type_name -> google.protobuf.Any
	1,  // 3: httpbingo.v1.HttpBinService.HttpMethodGet:input_type -> httpbingo.v1.HttpMethodGetReq
	3,  // 4: httpbingo.v1.HttpBinService.HttpMethodPut:input_type -> httpbingo.v1.HttpMethodPutReq
	5,  // 5: httpbingo.v1.HttpBinService.HttpMethodDelete:input_type -> httpbingo.v1.HttpMethodDeleteReq
	7,  // 6: httpbingo.v1.HttpBinService.HttpMethodPost:input_type -> httpbingo.v1.HttpMethodPostReq
	9,  // 7: httpbingo.v1.HttpBinService.StatusCode:input_type -> httpbingo.v1.StatusCodeReq
	11, // 8: httpbingo.v1.HttpBinService.Headers:input_type -> httpbingo.v1.HeadersReq
	13, // 9: httpbingo.v1.HttpBinService.IP:input_type -> httpbingo.v1.IPReq
	15, // 10: httpbingo.v1.HttpBinService.UserAgent:input_type -> httpbingo.v1.UserAgentReq
	0,  // 11: httpbingo.v1.HttpBinService.HttpMethodGet:output_type -> httpbingo.v1.ResponseReply
	0,  // 12: httpbingo.v1.HttpBinService.HttpMethodPut:output_type -> httpbingo.v1.ResponseReply
	0,  // 13: httpbingo.v1.HttpBinService.HttpMethodDelete:output_type -> httpbingo.v1.ResponseReply
	0,  // 14: httpbingo.v1.HttpBinService.HttpMethodPost:output_type -> httpbingo.v1.ResponseReply
	0,  // 15: httpbingo.v1.HttpBinService.StatusCode:output_type -> httpbingo.v1.ResponseReply
	0,  // 16: httpbingo.v1.HttpBinService.Headers:output_type -> httpbingo.v1.ResponseReply
	0,  // 17: httpbingo.v1.HttpBinService.IP:output_type -> httpbingo.v1.ResponseReply
	0,  // 18: httpbingo.v1.HttpBinService.UserAgent:output_type -> httpbingo.v1.ResponseReply
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_api_proto_init() }
func file_api_v1_api_proto_init() {
	if File_api_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseReply); i {
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
		file_api_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMethodGetReq); i {
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
		file_api_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMethodGetReply); i {
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
		file_api_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMethodPutReq); i {
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
		file_api_v1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMethodPutReply); i {
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
		file_api_v1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMethodDeleteReq); i {
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
		file_api_v1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMethodDeleteReply); i {
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
		file_api_v1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMethodPostReq); i {
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
		file_api_v1_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMethodPostReply); i {
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
		file_api_v1_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusCodeReq); i {
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
		file_api_v1_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusCodeReply); i {
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
		file_api_v1_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeadersReq); i {
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
		file_api_v1_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeadersReply); i {
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
		file_api_v1_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPReq); i {
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
		file_api_v1_api_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPReply); i {
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
		file_api_v1_api_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAgentReq); i {
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
		file_api_v1_api_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAgentReply); i {
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
			RawDescriptor: file_api_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_api_proto_goTypes,
		DependencyIndexes: file_api_v1_api_proto_depIdxs,
		MessageInfos:      file_api_v1_api_proto_msgTypes,
	}.Build()
	File_api_v1_api_proto = out.File
	file_api_v1_api_proto_rawDesc = nil
	file_api_v1_api_proto_goTypes = nil
	file_api_v1_api_proto_depIdxs = nil
}
