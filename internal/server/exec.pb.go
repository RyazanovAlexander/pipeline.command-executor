// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: exec.proto

package server

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

type ExecCommands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands []string `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *ExecCommands) Reset() {
	*x = ExecCommands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecCommands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecCommands) ProtoMessage() {}

func (x *ExecCommands) ProtoReflect() protoreflect.Message {
	mi := &file_exec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecCommands.ProtoReflect.Descriptor instead.
func (*ExecCommands) Descriptor() ([]byte, []int) {
	return file_exec_proto_rawDescGZIP(), []int{0}
}

func (x *ExecCommands) GetCommands() []string {
	if x != nil {
		return x.Commands
	}
	return nil
}

type ExecResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       bool     `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	ErrorMessage string   `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Output       []string `protobuf:"bytes,4,rep,name=output,proto3" json:"output,omitempty"`
}

func (x *ExecResult) Reset() {
	*x = ExecResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecResult) ProtoMessage() {}

func (x *ExecResult) ProtoReflect() protoreflect.Message {
	mi := &file_exec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecResult.ProtoReflect.Descriptor instead.
func (*ExecResult) Descriptor() ([]byte, []int) {
	return file_exec_proto_rawDescGZIP(), []int{1}
}

func (x *ExecResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *ExecResult) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ExecResult) GetOutput() []string {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_exec_proto protoreflect.FileDescriptor

var file_exec_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x78,
	0x65, 0x63, 0x22, 0x2a, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x60,
	0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x32, 0x48, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x12, 0x12, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x10, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exec_proto_rawDescOnce sync.Once
	file_exec_proto_rawDescData = file_exec_proto_rawDesc
)

func file_exec_proto_rawDescGZIP() []byte {
	file_exec_proto_rawDescOnce.Do(func() {
		file_exec_proto_rawDescData = protoimpl.X.CompressGZIP(file_exec_proto_rawDescData)
	})
	return file_exec_proto_rawDescData
}

var file_exec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_exec_proto_goTypes = []interface{}{
	(*ExecCommands)(nil), // 0: exec.ExecCommands
	(*ExecResult)(nil),   // 1: exec.ExecResult
}
var file_exec_proto_depIdxs = []int32{
	0, // 0: exec.ExecService.ExecuteCommands:input_type -> exec.ExecCommands
	1, // 1: exec.ExecService.ExecuteCommands:output_type -> exec.ExecResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_exec_proto_init() }
func file_exec_proto_init() {
	if File_exec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecCommands); i {
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
		file_exec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecResult); i {
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
			RawDescriptor: file_exec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exec_proto_goTypes,
		DependencyIndexes: file_exec_proto_depIdxs,
		MessageInfos:      file_exec_proto_msgTypes,
	}.Build()
	File_exec_proto = out.File
	file_exec_proto_rawDesc = nil
	file_exec_proto_goTypes = nil
	file_exec_proto_depIdxs = nil
}
