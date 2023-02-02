// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/pipeline/pb/rpc.proto

package pipeline

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

type RunPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *RunPipelineRequest) Reset() {
	*x = RunPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pipeline_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunPipelineRequest) ProtoMessage() {}

func (x *RunPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pipeline_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunPipelineRequest.ProtoReflect.Descriptor instead.
func (*RunPipelineRequest) Descriptor() ([]byte, []int) {
	return file_apps_pipeline_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RunPipelineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type QueryPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryPipelineRequest) Reset() {
	*x = QueryPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pipeline_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPipelineRequest) ProtoMessage() {}

func (x *QueryPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pipeline_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPipelineRequest.ProtoReflect.Descriptor instead.
func (*QueryPipelineRequest) Descriptor() ([]byte, []int) {
	return file_apps_pipeline_pb_rpc_proto_rawDescGZIP(), []int{1}
}

var File_apps_pipeline_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_pipeline_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x75, 0x6e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16,
	0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xbb, 0x02, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x61,
	0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2d, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x68, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x2f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x74, 0x12, 0x67, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x30, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_pipeline_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_pipeline_pb_rpc_proto_rawDescData = file_apps_pipeline_pb_rpc_proto_rawDesc
)

func file_apps_pipeline_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_pipeline_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_pipeline_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_pipeline_pb_rpc_proto_rawDescData)
	})
	return file_apps_pipeline_pb_rpc_proto_rawDescData
}

var file_apps_pipeline_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_pipeline_pb_rpc_proto_goTypes = []interface{}{
	(*RunPipelineRequest)(nil),    // 0: infraboard.mpaas.pipeline.RunPipelineRequest
	(*QueryPipelineRequest)(nil),  // 1: infraboard.mpaas.pipeline.QueryPipelineRequest
	(*CreatePipelineRequest)(nil), // 2: infraboard.mpaas.pipeline.CreatePipelineRequest
	(*Pipeline)(nil),              // 3: infraboard.mpaas.pipeline.Pipeline
	(*PipelineSet)(nil),           // 4: infraboard.mpaas.pipeline.PipelineSet
}
var file_apps_pipeline_pb_rpc_proto_depIdxs = []int32{
	0, // 0: infraboard.mpaas.pipeline.RPC.RunPipeline:input_type -> infraboard.mpaas.pipeline.RunPipelineRequest
	1, // 1: infraboard.mpaas.pipeline.RPC.QueryPipeline:input_type -> infraboard.mpaas.pipeline.QueryPipelineRequest
	2, // 2: infraboard.mpaas.pipeline.RPC.CreatePipeline:input_type -> infraboard.mpaas.pipeline.CreatePipelineRequest
	3, // 3: infraboard.mpaas.pipeline.RPC.RunPipeline:output_type -> infraboard.mpaas.pipeline.Pipeline
	4, // 4: infraboard.mpaas.pipeline.RPC.QueryPipeline:output_type -> infraboard.mpaas.pipeline.PipelineSet
	3, // 5: infraboard.mpaas.pipeline.RPC.CreatePipeline:output_type -> infraboard.mpaas.pipeline.Pipeline
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_pipeline_pb_rpc_proto_init() }
func file_apps_pipeline_pb_rpc_proto_init() {
	if File_apps_pipeline_pb_rpc_proto != nil {
		return
	}
	file_apps_pipeline_pb_pipeline_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_pipeline_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunPipelineRequest); i {
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
		file_apps_pipeline_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPipelineRequest); i {
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
			RawDescriptor: file_apps_pipeline_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_pipeline_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_pipeline_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_pipeline_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_pipeline_pb_rpc_proto = out.File
	file_apps_pipeline_pb_rpc_proto_rawDesc = nil
	file_apps_pipeline_pb_rpc_proto_goTypes = nil
	file_apps_pipeline_pb_rpc_proto_depIdxs = nil
}
