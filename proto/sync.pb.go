// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.4
// source: sync.proto

package proto

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

type SetRecordStatus int32

const (
	SetRecordStatus_SUCCESS  SetRecordStatus = 0
	SetRecordStatus_CONFLICT SetRecordStatus = 1
)

// Enum value maps for SetRecordStatus.
var (
	SetRecordStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "CONFLICT",
	}
	SetRecordStatus_value = map[string]int32{
		"SUCCESS":  0,
		"CONFLICT": 1,
	}
)

func (x SetRecordStatus) Enum() *SetRecordStatus {
	p := new(SetRecordStatus)
	*p = x
	return p
}

func (x SetRecordStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetRecordStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_proto_enumTypes[0].Descriptor()
}

func (SetRecordStatus) Type() protoreflect.EnumType {
	return &file_sync_proto_enumTypes[0]
}

func (x SetRecordStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetRecordStatus.Descriptor instead.
func (SetRecordStatus) EnumDescriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{0}
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index   uint64  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Version float32 `protobuf:"fixed32,2,opt,name=version,proto3" json:"version,omitempty"`
	Data    []byte  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Record) GetVersion() float32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Record) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record      *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	RequestTime int64   `protobuf:"varint,2,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	Signature   string  `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SetRecordRequest) Reset() {
	*x = SetRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRecordRequest) ProtoMessage() {}

func (x *SetRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRecordRequest.ProtoReflect.Descriptor instead.
func (*SetRecordRequest) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{1}
}

func (x *SetRecordRequest) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *SetRecordRequest) GetRequestTime() int64 {
	if x != nil {
		return x.RequestTime
	}
	return 0
}

func (x *SetRecordRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type SetRecordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     SetRecordStatus `protobuf:"varint,1,opt,name=status,proto3,enum=sync.SetRecordStatus" json:"status,omitempty"`
	NewVersion int64           `protobuf:"varint,2,opt,name=new_version,json=newVersion,proto3" json:"new_version,omitempty"`
}

func (x *SetRecordReply) Reset() {
	*x = SetRecordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRecordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRecordReply) ProtoMessage() {}

func (x *SetRecordReply) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRecordReply.ProtoReflect.Descriptor instead.
func (*SetRecordReply) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{2}
}

func (x *SetRecordReply) GetStatus() SetRecordStatus {
	if x != nil {
		return x.Status
	}
	return SetRecordStatus_SUCCESS
}

func (x *SetRecordReply) GetNewVersion() int64 {
	if x != nil {
		return x.NewVersion
	}
	return 0
}

type ListChangesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SinceVersion int64  `protobuf:"varint,1,opt,name=since_version,json=sinceVersion,proto3" json:"since_version,omitempty"`
	RequestTime  int64  `protobuf:"varint,2,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	Signature    string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ListChangesRequest) Reset() {
	*x = ListChangesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChangesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChangesRequest) ProtoMessage() {}

func (x *ListChangesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChangesRequest.ProtoReflect.Descriptor instead.
func (*ListChangesRequest) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{3}
}

func (x *ListChangesRequest) GetSinceVersion() int64 {
	if x != nil {
		return x.SinceVersion
	}
	return 0
}

func (x *ListChangesRequest) GetRequestTime() int64 {
	if x != nil {
		return x.RequestTime
	}
	return 0
}

func (x *ListChangesRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type ListChangesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*Record `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *ListChangesReply) Reset() {
	*x = ListChangesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChangesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChangesReply) ProtoMessage() {}

func (x *ListChangesReply) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChangesReply.ProtoReflect.Descriptor instead.
func (*ListChangesReply) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{4}
}

func (x *ListChangesReply) GetChanges() []*Record {
	if x != nil {
		return x.Changes
	}
	return nil
}

var File_sync_proto protoreflect.FileDescriptor

var file_sync_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x79,
	0x6e, 0x63, 0x22, 0x4c, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x79, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x60, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x6e, 0x65, 0x77, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x7a, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x69, 0x6e, 0x63,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x3a, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x2a, 0x2c, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43,
	0x54, 0x10, 0x01, 0x32, 0x88, 0x01, 0x0a, 0x06, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x3b,
	0x0a, 0x09, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x65,
	0x65, 0x7a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_proto_rawDescOnce sync.Once
	file_sync_proto_rawDescData = file_sync_proto_rawDesc
)

func file_sync_proto_rawDescGZIP() []byte {
	file_sync_proto_rawDescOnce.Do(func() {
		file_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_proto_rawDescData)
	})
	return file_sync_proto_rawDescData
}

var file_sync_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sync_proto_goTypes = []any{
	(SetRecordStatus)(0),       // 0: sync.SetRecordStatus
	(*Record)(nil),             // 1: sync.Record
	(*SetRecordRequest)(nil),   // 2: sync.SetRecordRequest
	(*SetRecordReply)(nil),     // 3: sync.SetRecordReply
	(*ListChangesRequest)(nil), // 4: sync.ListChangesRequest
	(*ListChangesReply)(nil),   // 5: sync.ListChangesReply
}
var file_sync_proto_depIdxs = []int32{
	1, // 0: sync.SetRecordRequest.record:type_name -> sync.Record
	0, // 1: sync.SetRecordReply.status:type_name -> sync.SetRecordStatus
	1, // 2: sync.ListChangesReply.changes:type_name -> sync.Record
	2, // 3: sync.Syncer.SetRecord:input_type -> sync.SetRecordRequest
	4, // 4: sync.Syncer.ListChanges:input_type -> sync.ListChangesRequest
	3, // 5: sync.Syncer.SetRecord:output_type -> sync.SetRecordReply
	5, // 6: sync.Syncer.ListChanges:output_type -> sync.ListChangesReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sync_proto_init() }
func file_sync_proto_init() {
	if File_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Record); i {
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
		file_sync_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetRecordRequest); i {
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
		file_sync_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SetRecordReply); i {
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
		file_sync_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListChangesRequest); i {
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
		file_sync_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListChangesReply); i {
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
			RawDescriptor: file_sync_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sync_proto_goTypes,
		DependencyIndexes: file_sync_proto_depIdxs,
		EnumInfos:         file_sync_proto_enumTypes,
		MessageInfos:      file_sync_proto_msgTypes,
	}.Build()
	File_sync_proto = out.File
	file_sync_proto_rawDesc = nil
	file_sync_proto_goTypes = nil
	file_sync_proto_depIdxs = nil
}
