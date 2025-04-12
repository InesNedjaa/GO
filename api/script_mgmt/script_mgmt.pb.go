// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: api/script_mgmt/script_mgmt.proto

package script_mgmt

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

type ScheduleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Time          string                 `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"` // "seconds minutes hour day_of_month month day_of_week " exemple "*/5 * * * * *"
	ScriptPath    string                 `protobuf:"bytes,2,opt,name=script_path,json=scriptPath,proto3" json:"script_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScheduleRequest) Reset() {
	*x = ScheduleRequest{}
	mi := &file_api_script_mgmt_script_mgmt_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequest) ProtoMessage() {}

func (x *ScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_script_mgmt_script_mgmt_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequest.ProtoReflect.Descriptor instead.
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return file_api_script_mgmt_script_mgmt_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *ScheduleRequest) GetScriptPath() string {
	if x != nil {
		return x.ScriptPath
	}
	return ""
}

type ScheduleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Id            int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScheduleResponse) Reset() {
	*x = ScheduleResponse{}
	mi := &file_api_script_mgmt_script_mgmt_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleResponse) ProtoMessage() {}

func (x *ScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_script_mgmt_script_mgmt_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleResponse.ProtoReflect.Descriptor instead.
func (*ScheduleResponse) Descriptor() ([]byte, []int) {
	return file_api_script_mgmt_script_mgmt_proto_rawDescGZIP(), []int{1}
}

func (x *ScheduleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ScheduleResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_script_mgmt_script_mgmt_proto protoreflect.FileDescriptor

const file_api_script_mgmt_script_mgmt_proto_rawDesc = "" +
	"\n" +
	"!api/script_mgmt/script_mgmt.proto\"F\n" +
	"\x0fScheduleRequest\x12\x12\n" +
	"\x04time\x18\x01 \x01(\tR\x04time\x12\x1f\n" +
	"\vscript_path\x18\x02 \x01(\tR\n" +
	"scriptPath\"<\n" +
	"\x10ScheduleResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\x05R\x02id2F\n" +
	"\vScript_Mgmt\x127\n" +
	"\x0eScheduleScript\x12\x10.ScheduleRequest\x1a\x11.ScheduleResponse\"\x00B\x11Z\x0fapi/script_mgmtb\x06proto3"

var (
	file_api_script_mgmt_script_mgmt_proto_rawDescOnce sync.Once
	file_api_script_mgmt_script_mgmt_proto_rawDescData []byte
)

func file_api_script_mgmt_script_mgmt_proto_rawDescGZIP() []byte {
	file_api_script_mgmt_script_mgmt_proto_rawDescOnce.Do(func() {
		file_api_script_mgmt_script_mgmt_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_script_mgmt_script_mgmt_proto_rawDesc), len(file_api_script_mgmt_script_mgmt_proto_rawDesc)))
	})
	return file_api_script_mgmt_script_mgmt_proto_rawDescData
}

var file_api_script_mgmt_script_mgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_script_mgmt_script_mgmt_proto_goTypes = []any{
	(*ScheduleRequest)(nil),  // 0: ScheduleRequest
	(*ScheduleResponse)(nil), // 1: ScheduleResponse
}
var file_api_script_mgmt_script_mgmt_proto_depIdxs = []int32{
	0, // 0: Script_Mgmt.ScheduleScript:input_type -> ScheduleRequest
	1, // 1: Script_Mgmt.ScheduleScript:output_type -> ScheduleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_script_mgmt_script_mgmt_proto_init() }
func file_api_script_mgmt_script_mgmt_proto_init() {
	if File_api_script_mgmt_script_mgmt_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_script_mgmt_script_mgmt_proto_rawDesc), len(file_api_script_mgmt_script_mgmt_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_script_mgmt_script_mgmt_proto_goTypes,
		DependencyIndexes: file_api_script_mgmt_script_mgmt_proto_depIdxs,
		MessageInfos:      file_api_script_mgmt_script_mgmt_proto_msgTypes,
	}.Build()
	File_api_script_mgmt_script_mgmt_proto = out.File
	file_api_script_mgmt_script_mgmt_proto_goTypes = nil
	file_api_script_mgmt_script_mgmt_proto_depIdxs = nil
}
