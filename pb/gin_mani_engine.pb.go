// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: gin_mani_engine.proto

package pb_mani

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FileAction int32

const (
	FileAction_unknown_file_action FileAction = 0
)

// Enum value maps for FileAction.
var (
	FileAction_name = map[int32]string{
		0: "unknown_file_action",
	}
	FileAction_value = map[string]int32{
		"unknown_file_action": 0,
	}
)

func (x FileAction) Enum() *FileAction {
	p := new(FileAction)
	*p = x
	return p
}

func (x FileAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileAction) Descriptor() protoreflect.EnumDescriptor {
	return file_gin_mani_engine_proto_enumTypes[0].Descriptor()
}

func (FileAction) Type() protoreflect.EnumType {
	return &file_gin_mani_engine_proto_enumTypes[0]
}

func (x FileAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileAction.Descriptor instead.
func (FileAction) EnumDescriptor() ([]byte, []int) {
	return file_gin_mani_engine_proto_rawDescGZIP(), []int{0}
}

type FileUriToServerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName   string     `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Files      []byte     `protobuf:"bytes,2,opt,name=files,proto3" json:"files,omitempty"`
	FileAction FileAction `protobuf:"varint,3,opt,name=file_action,json=fileAction,proto3,enum=pb_mani.FileAction" json:"file_action,omitempty"`
	Base       *Base      `protobuf:"bytes,255,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *FileUriToServerReq) Reset() {
	*x = FileUriToServerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gin_mani_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUriToServerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUriToServerReq) ProtoMessage() {}

func (x *FileUriToServerReq) ProtoReflect() protoreflect.Message {
	mi := &file_gin_mani_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUriToServerReq.ProtoReflect.Descriptor instead.
func (*FileUriToServerReq) Descriptor() ([]byte, []int) {
	return file_gin_mani_engine_proto_rawDescGZIP(), []int{0}
}

func (x *FileUriToServerReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileUriToServerReq) GetFiles() []byte {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *FileUriToServerReq) GetFileAction() FileAction {
	if x != nil {
		return x.FileAction
	}
	return FileAction_unknown_file_action
}

func (x *FileUriToServerReq) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type FileUriToServerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaveUrl  string    `protobuf:"bytes,1,opt,name=save_url,json=saveUrl,proto3" json:"save_url,omitempty"`
	BaseResp *BaseResp `protobuf:"bytes,255,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *FileUriToServerResp) Reset() {
	*x = FileUriToServerResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gin_mani_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUriToServerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUriToServerResp) ProtoMessage() {}

func (x *FileUriToServerResp) ProtoReflect() protoreflect.Message {
	mi := &file_gin_mani_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUriToServerResp.ProtoReflect.Descriptor instead.
func (*FileUriToServerResp) Descriptor() ([]byte, []int) {
	return file_gin_mani_engine_proto_rawDescGZIP(), []int{1}
}

func (x *FileUriToServerResp) GetSaveUrl() string {
	if x != nil {
		return x.SaveUrl
	}
	return ""
}

func (x *FileUriToServerResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type FileUriToCrmReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaveUrl    string     `protobuf:"bytes,1,opt,name=save_url,json=saveUrl,proto3" json:"save_url,omitempty"`
	FileAction FileAction `protobuf:"varint,3,opt,name=file_action,json=fileAction,proto3,enum=pb_mani.FileAction" json:"file_action,omitempty"`
	Base       *Base      `protobuf:"bytes,255,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *FileUriToCrmReq) Reset() {
	*x = FileUriToCrmReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gin_mani_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUriToCrmReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUriToCrmReq) ProtoMessage() {}

func (x *FileUriToCrmReq) ProtoReflect() protoreflect.Message {
	mi := &file_gin_mani_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUriToCrmReq.ProtoReflect.Descriptor instead.
func (*FileUriToCrmReq) Descriptor() ([]byte, []int) {
	return file_gin_mani_engine_proto_rawDescGZIP(), []int{2}
}

func (x *FileUriToCrmReq) GetSaveUrl() string {
	if x != nil {
		return x.SaveUrl
	}
	return ""
}

func (x *FileUriToCrmReq) GetFileAction() FileAction {
	if x != nil {
		return x.FileAction
	}
	return FileAction_unknown_file_action
}

func (x *FileUriToCrmReq) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type FileUriToCrmResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string    `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	BaseResp *BaseResp `protobuf:"bytes,255,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *FileUriToCrmResp) Reset() {
	*x = FileUriToCrmResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gin_mani_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUriToCrmResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUriToCrmResp) ProtoMessage() {}

func (x *FileUriToCrmResp) ProtoReflect() protoreflect.Message {
	mi := &file_gin_mani_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUriToCrmResp.ProtoReflect.Descriptor instead.
func (*FileUriToCrmResp) Descriptor() ([]byte, []int) {
	return file_gin_mani_engine_proto_rawDescGZIP(), []int{3}
}

func (x *FileUriToCrmResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FileUriToCrmResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_gin_mani_engine_proto protoreflect.FileDescriptor

var file_gin_mani_engine_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x5f, 0x6d, 0x61, 0x6e, 0x69,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1,
	0x01, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x70, 0x62, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x62, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x22, 0x61, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x54, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x76,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x76,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x61,
	0x6e, 0x69, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x86, 0x01, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72,
	0x69, 0x54, 0x6f, 0x43, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x76,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x76,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x5f, 0x6d,
	0x61, 0x6e, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x66, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x5f, 0x6d,
	0x61, 0x6e, 0x69, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x55,
	0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x54, 0x6f, 0x43, 0x72, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x61,
	0x6e, 0x69, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x2a, 0x25, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x32, 0xa5, 0x01, 0x0a,
	0x10, 0x47, 0x69, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x54, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x72, 0x69, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x43, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x54, 0x6f, 0x43, 0x72, 0x6d, 0x12,
	0x18, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72,
	0x69, 0x54, 0x6f, 0x43, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x5f, 0x6d,
	0x61, 0x6e, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x54, 0x6f, 0x43, 0x72, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gin_mani_engine_proto_rawDescOnce sync.Once
	file_gin_mani_engine_proto_rawDescData = file_gin_mani_engine_proto_rawDesc
)

func file_gin_mani_engine_proto_rawDescGZIP() []byte {
	file_gin_mani_engine_proto_rawDescOnce.Do(func() {
		file_gin_mani_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_gin_mani_engine_proto_rawDescData)
	})
	return file_gin_mani_engine_proto_rawDescData
}

var file_gin_mani_engine_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gin_mani_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gin_mani_engine_proto_goTypes = []interface{}{
	(FileAction)(0),             // 0: pb_mani.FileAction
	(*FileUriToServerReq)(nil),  // 1: pb_mani.FileUriToServerReq
	(*FileUriToServerResp)(nil), // 2: pb_mani.FileUriToServerResp
	(*FileUriToCrmReq)(nil),     // 3: pb_mani.FileUriToCrmReq
	(*FileUriToCrmResp)(nil),    // 4: pb_mani.FileUriToCrmResp
	(*Base)(nil),                // 5: pb_mani.Base
	(*BaseResp)(nil),            // 6: pb_mani.BaseResp
}
var file_gin_mani_engine_proto_depIdxs = []int32{
	0, // 0: pb_mani.FileUriToServerReq.file_action:type_name -> pb_mani.FileAction
	5, // 1: pb_mani.FileUriToServerReq.base:type_name -> pb_mani.Base
	6, // 2: pb_mani.FileUriToServerResp.base_resp:type_name -> pb_mani.BaseResp
	0, // 3: pb_mani.FileUriToCrmReq.file_action:type_name -> pb_mani.FileAction
	5, // 4: pb_mani.FileUriToCrmReq.base:type_name -> pb_mani.Base
	6, // 5: pb_mani.FileUriToCrmResp.base_resp:type_name -> pb_mani.BaseResp
	1, // 6: pb_mani.GinEngineService.FileUriToServer:input_type -> pb_mani.FileUriToServerReq
	3, // 7: pb_mani.GinEngineService.FileUriToCrm:input_type -> pb_mani.FileUriToCrmReq
	2, // 8: pb_mani.GinEngineService.FileUriToServer:output_type -> pb_mani.FileUriToServerResp
	4, // 9: pb_mani.GinEngineService.FileUriToCrm:output_type -> pb_mani.FileUriToCrmResp
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_gin_mani_engine_proto_init() }
func file_gin_mani_engine_proto_init() {
	if File_gin_mani_engine_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gin_mani_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUriToServerReq); i {
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
		file_gin_mani_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUriToServerResp); i {
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
		file_gin_mani_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUriToCrmReq); i {
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
		file_gin_mani_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUriToCrmResp); i {
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
			RawDescriptor: file_gin_mani_engine_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gin_mani_engine_proto_goTypes,
		DependencyIndexes: file_gin_mani_engine_proto_depIdxs,
		EnumInfos:         file_gin_mani_engine_proto_enumTypes,
		MessageInfos:      file_gin_mani_engine_proto_msgTypes,
	}.Build()
	File_gin_mani_engine_proto = out.File
	file_gin_mani_engine_proto_rawDesc = nil
	file_gin_mani_engine_proto_goTypes = nil
	file_gin_mani_engine_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GinEngineServiceClient is the client API for GinEngineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GinEngineServiceClient interface {
	FileUriToServer(ctx context.Context, in *FileUriToServerReq, opts ...grpc.CallOption) (*FileUriToServerResp, error)
	FileUriToCrm(ctx context.Context, in *FileUriToCrmReq, opts ...grpc.CallOption) (*FileUriToCrmResp, error)
}

type ginEngineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGinEngineServiceClient(cc grpc.ClientConnInterface) GinEngineServiceClient {
	return &ginEngineServiceClient{cc}
}

func (c *ginEngineServiceClient) FileUriToServer(ctx context.Context, in *FileUriToServerReq, opts ...grpc.CallOption) (*FileUriToServerResp, error) {
	out := new(FileUriToServerResp)
	err := c.cc.Invoke(ctx, "/pb_mani.GinEngineService/FileUriToServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ginEngineServiceClient) FileUriToCrm(ctx context.Context, in *FileUriToCrmReq, opts ...grpc.CallOption) (*FileUriToCrmResp, error) {
	out := new(FileUriToCrmResp)
	err := c.cc.Invoke(ctx, "/pb_mani.GinEngineService/FileUriToCrm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GinEngineServiceServer is the server API for GinEngineService service.
type GinEngineServiceServer interface {
	FileUriToServer(context.Context, *FileUriToServerReq) (*FileUriToServerResp, error)
	FileUriToCrm(context.Context, *FileUriToCrmReq) (*FileUriToCrmResp, error)
}

// UnimplementedGinEngineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGinEngineServiceServer struct {
}

func (*UnimplementedGinEngineServiceServer) FileUriToServer(context.Context, *FileUriToServerReq) (*FileUriToServerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileUriToServer not implemented")
}
func (*UnimplementedGinEngineServiceServer) FileUriToCrm(context.Context, *FileUriToCrmReq) (*FileUriToCrmResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileUriToCrm not implemented")
}

func RegisterGinEngineServiceServer(s *grpc.Server, srv GinEngineServiceServer) {
	s.RegisterService(&_GinEngineService_serviceDesc, srv)
}

func _GinEngineService_FileUriToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUriToServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GinEngineServiceServer).FileUriToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_mani.GinEngineService/FileUriToServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GinEngineServiceServer).FileUriToServer(ctx, req.(*FileUriToServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GinEngineService_FileUriToCrm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUriToCrmReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GinEngineServiceServer).FileUriToCrm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_mani.GinEngineService/FileUriToCrm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GinEngineServiceServer).FileUriToCrm(ctx, req.(*FileUriToCrmReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GinEngineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb_mani.GinEngineService",
	HandlerType: (*GinEngineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileUriToServer",
			Handler:    _GinEngineService_FileUriToServer_Handler,
		},
		{
			MethodName: "FileUriToCrm",
			Handler:    _GinEngineService_FileUriToCrm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gin_mani_engine.proto",
}
