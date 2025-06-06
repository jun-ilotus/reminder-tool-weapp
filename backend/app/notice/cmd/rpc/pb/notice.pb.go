// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: notice.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// --------------------------------noticeLog--------------------------------
type NoticeLog struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                 //id
	UserId        int64                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`         //userId
	UserOpenid    string                 `protobuf:"bytes,3,opt,name=userOpenid,proto3" json:"userOpenid,omitempty"`  //userOpenid
	CreateTime    int64                  `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"` //createTime
	UpdateTime    int64                  `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"` //updateTime
	Status        int64                  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`         //0 寰呭彂閫侊紱1鍙戦€佹垚鍔燂紱鍏朵粬涓哄彂閫佸け璐ョ殑澶辫触鐮?
	Error         string                 `protobuf:"bytes,8,opt,name=error,proto3" json:"error,omitempty"`            //鍙戦€佸け璐ュ師鍥?
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoticeLog) Reset() {
	*x = NoticeLog{}
	mi := &file_notice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoticeLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeLog) ProtoMessage() {}

func (x *NoticeLog) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeLog.ProtoReflect.Descriptor instead.
func (*NoticeLog) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{0}
}

func (x *NoticeLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NoticeLog) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *NoticeLog) GetUserOpenid() string {
	if x != nil {
		return x.UserOpenid
	}
	return ""
}

func (x *NoticeLog) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *NoticeLog) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *NoticeLog) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *NoticeLog) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AddNoticeLogReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`        //userId
	UserOpenid    string                 `protobuf:"bytes,2,opt,name=userOpenid,proto3" json:"userOpenid,omitempty"` //userOpenid
	Status        int64                  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`        //0 寰呭彂閫侊紱1鍙戦€佹垚鍔燂紱鍏朵粬涓哄彂閫佸け璐ョ殑澶辫触鐮
	Error         string                 `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`           //鍙戦€佸け璐ュ師鍥?
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddNoticeLogReq) Reset() {
	*x = AddNoticeLogReq{}
	mi := &file_notice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddNoticeLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoticeLogReq) ProtoMessage() {}

func (x *AddNoticeLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoticeLogReq.ProtoReflect.Descriptor instead.
func (*AddNoticeLogReq) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{1}
}

func (x *AddNoticeLogReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddNoticeLogReq) GetUserOpenid() string {
	if x != nil {
		return x.UserOpenid
	}
	return ""
}

func (x *AddNoticeLogReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddNoticeLogReq) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AddNoticeLogResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddNoticeLogResp) Reset() {
	*x = AddNoticeLogResp{}
	mi := &file_notice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddNoticeLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoticeLogResp) ProtoMessage() {}

func (x *AddNoticeLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoticeLogResp.ProtoReflect.Descriptor instead.
func (*AddNoticeLogResp) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{2}
}

func (x *AddNoticeLogResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateNoticeLogReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                //id
	UserId        int64                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`        //userId
	UserOpenid    string                 `protobuf:"bytes,3,opt,name=userOpenid,proto3" json:"userOpenid,omitempty"` //userOpenid
	Status        int64                  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`        //0 寰呭彂閫侊紱1鍙戦€佹垚鍔燂紱鍏朵粬涓哄彂閫佸け璐ョ殑澶辫触鐮?
	Error         string                 `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`           //鍙戦€佸け璐ュ師鍥?
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNoticeLogReq) Reset() {
	*x = UpdateNoticeLogReq{}
	mi := &file_notice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNoticeLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoticeLogReq) ProtoMessage() {}

func (x *UpdateNoticeLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoticeLogReq.ProtoReflect.Descriptor instead.
func (*UpdateNoticeLogReq) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNoticeLogReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateNoticeLogReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateNoticeLogReq) GetUserOpenid() string {
	if x != nil {
		return x.UserOpenid
	}
	return ""
}

func (x *UpdateNoticeLogReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateNoticeLogReq) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateNoticeLogResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNoticeLogResp) Reset() {
	*x = UpdateNoticeLogResp{}
	mi := &file_notice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNoticeLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoticeLogResp) ProtoMessage() {}

func (x *UpdateNoticeLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoticeLogResp.ProtoReflect.Descriptor instead.
func (*UpdateNoticeLogResp) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{4}
}

type DelNoticeLogReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelNoticeLogReq) Reset() {
	*x = DelNoticeLogReq{}
	mi := &file_notice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelNoticeLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelNoticeLogReq) ProtoMessage() {}

func (x *DelNoticeLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelNoticeLogReq.ProtoReflect.Descriptor instead.
func (*DelNoticeLogReq) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{5}
}

func (x *DelNoticeLogReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelNoticeLogResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelNoticeLogResp) Reset() {
	*x = DelNoticeLogResp{}
	mi := &file_notice_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelNoticeLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelNoticeLogResp) ProtoMessage() {}

func (x *DelNoticeLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelNoticeLogResp.ProtoReflect.Descriptor instead.
func (*DelNoticeLogResp) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{6}
}

type GetNoticeLogByIdReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNoticeLogByIdReq) Reset() {
	*x = GetNoticeLogByIdReq{}
	mi := &file_notice_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNoticeLogByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoticeLogByIdReq) ProtoMessage() {}

func (x *GetNoticeLogByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoticeLogByIdReq.ProtoReflect.Descriptor instead.
func (*GetNoticeLogByIdReq) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{7}
}

func (x *GetNoticeLogByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetNoticeLogByIdResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NoticeLog     *NoticeLog             `protobuf:"bytes,1,opt,name=noticeLog,proto3" json:"noticeLog,omitempty"` //noticeLog
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNoticeLogByIdResp) Reset() {
	*x = GetNoticeLogByIdResp{}
	mi := &file_notice_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNoticeLogByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoticeLogByIdResp) ProtoMessage() {}

func (x *GetNoticeLogByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoticeLogByIdResp.ProtoReflect.Descriptor instead.
func (*GetNoticeLogByIdResp) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{8}
}

func (x *GetNoticeLogByIdResp) GetNoticeLog() *NoticeLog {
	if x != nil {
		return x.NoticeLog
	}
	return nil
}

type SearchNoticeLogReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`             //page
	Limit         int64                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`           //limit
	Id            int64                  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`                 //id
	UserId        int64                  `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`         //userId
	UserOpenid    string                 `protobuf:"bytes,5,opt,name=userOpenid,proto3" json:"userOpenid,omitempty"`  //userOpenid
	CreateTime    int64                  `protobuf:"varint,7,opt,name=createTime,proto3" json:"createTime,omitempty"` //createTime
	UpdateTime    int64                  `protobuf:"varint,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"` //updateTime
	Status        int64                  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`         //0 寰呭彂閫侊紱1鍙戦€佹垚鍔燂紱鍏朵粬涓哄彂閫佸け璐ョ殑澶辫触鐮?
	Error         string                 `protobuf:"bytes,10,opt,name=error,proto3" json:"error,omitempty"`           //鍙戦€佸け璐ュ師鍥?
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchNoticeLogReq) Reset() {
	*x = SearchNoticeLogReq{}
	mi := &file_notice_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchNoticeLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchNoticeLogReq) ProtoMessage() {}

func (x *SearchNoticeLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchNoticeLogReq.ProtoReflect.Descriptor instead.
func (*SearchNoticeLogReq) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{9}
}

func (x *SearchNoticeLogReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchNoticeLogReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchNoticeLogReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchNoticeLogReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SearchNoticeLogReq) GetUserOpenid() string {
	if x != nil {
		return x.UserOpenid
	}
	return ""
}

func (x *SearchNoticeLogReq) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *SearchNoticeLogReq) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *SearchNoticeLogReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SearchNoticeLogReq) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SearchNoticeLogResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NoticeLog     []*NoticeLog           `protobuf:"bytes,1,rep,name=noticeLog,proto3" json:"noticeLog,omitempty"` //noticeLog
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchNoticeLogResp) Reset() {
	*x = SearchNoticeLogResp{}
	mi := &file_notice_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchNoticeLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchNoticeLogResp) ProtoMessage() {}

func (x *SearchNoticeLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchNoticeLogResp.ProtoReflect.Descriptor instead.
func (*SearchNoticeLogResp) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{10}
}

func (x *SearchNoticeLogResp) GetNoticeLog() []*NoticeLog {
	if x != nil {
		return x.NoticeLog
	}
	return nil
}

var File_notice_proto protoreflect.FileDescriptor

var file_notice_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0xc1, 0x01, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x77, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x22, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x25,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a,
	0x09, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52,
	0x09, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x22, 0xf4, 0x01, 0x0a, 0x12, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65,
	0x6e, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x42, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x4c, 0x6f, 0x67, 0x32, 0xcd, 0x02, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x0f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x39, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x4e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x42, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_notice_proto_rawDescOnce sync.Once
	file_notice_proto_rawDescData []byte
)

func file_notice_proto_rawDescGZIP() []byte {
	file_notice_proto_rawDescOnce.Do(func() {
		file_notice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_notice_proto_rawDesc), len(file_notice_proto_rawDesc)))
	})
	return file_notice_proto_rawDescData
}

var file_notice_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_notice_proto_goTypes = []any{
	(*NoticeLog)(nil),            // 0: pb.NoticeLog
	(*AddNoticeLogReq)(nil),      // 1: pb.AddNoticeLogReq
	(*AddNoticeLogResp)(nil),     // 2: pb.AddNoticeLogResp
	(*UpdateNoticeLogReq)(nil),   // 3: pb.UpdateNoticeLogReq
	(*UpdateNoticeLogResp)(nil),  // 4: pb.UpdateNoticeLogResp
	(*DelNoticeLogReq)(nil),      // 5: pb.DelNoticeLogReq
	(*DelNoticeLogResp)(nil),     // 6: pb.DelNoticeLogResp
	(*GetNoticeLogByIdReq)(nil),  // 7: pb.GetNoticeLogByIdReq
	(*GetNoticeLogByIdResp)(nil), // 8: pb.GetNoticeLogByIdResp
	(*SearchNoticeLogReq)(nil),   // 9: pb.SearchNoticeLogReq
	(*SearchNoticeLogResp)(nil),  // 10: pb.SearchNoticeLogResp
}
var file_notice_proto_depIdxs = []int32{
	0,  // 0: pb.GetNoticeLogByIdResp.noticeLog:type_name -> pb.NoticeLog
	0,  // 1: pb.SearchNoticeLogResp.noticeLog:type_name -> pb.NoticeLog
	1,  // 2: pb.notice.AddNoticeLog:input_type -> pb.AddNoticeLogReq
	3,  // 3: pb.notice.UpdateNoticeLog:input_type -> pb.UpdateNoticeLogReq
	5,  // 4: pb.notice.DelNoticeLog:input_type -> pb.DelNoticeLogReq
	7,  // 5: pb.notice.GetNoticeLogById:input_type -> pb.GetNoticeLogByIdReq
	9,  // 6: pb.notice.SearchNoticeLog:input_type -> pb.SearchNoticeLogReq
	2,  // 7: pb.notice.AddNoticeLog:output_type -> pb.AddNoticeLogResp
	4,  // 8: pb.notice.UpdateNoticeLog:output_type -> pb.UpdateNoticeLogResp
	6,  // 9: pb.notice.DelNoticeLog:output_type -> pb.DelNoticeLogResp
	8,  // 10: pb.notice.GetNoticeLogById:output_type -> pb.GetNoticeLogByIdResp
	10, // 11: pb.notice.SearchNoticeLog:output_type -> pb.SearchNoticeLogResp
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_notice_proto_init() }
func file_notice_proto_init() {
	if File_notice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_notice_proto_rawDesc), len(file_notice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notice_proto_goTypes,
		DependencyIndexes: file_notice_proto_depIdxs,
		MessageInfos:      file_notice_proto_msgTypes,
	}.Build()
	File_notice_proto = out.File
	file_notice_proto_goTypes = nil
	file_notice_proto_depIdxs = nil
}
