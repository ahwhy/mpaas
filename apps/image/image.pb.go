// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: apps/image/pb/image.proto

package image

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

// 镜像仓库
type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一ID
	// @gotags: json:"id" bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 录入时间
	// @gotags: json:"create_at" bson:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: json:"update_at" bson:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 更新人
	// @gotags: json:"update_by" bson:"update_by"
	UpdateBy string `protobuf:"bytes,4,opt,name=update_by,json=updateBy,proto3" json:"update_by" bson:"update_by"`
	// 基础信息
	// @gotags: json:"data" bson:"data"
	Data *CreateRegistry `protobuf:"bytes,9,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_image_pb_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_apps_image_pb_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_apps_image_pb_image_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Registry) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Registry) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Registry) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Registry) GetData() *CreateRegistry {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 仓库名称, 唯一
	// @gotags: json:"name" bson:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// 仓库地址
	// @gotags: json:"address" bson:"address"
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address" bson:"address"`
	// 描述
	// @gotags: json:"description" bson:"description"
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description" bson:"description"`
	// 集群标签
	// @gotags: json:"lables" form:"lables" bson:"lables"
	Lables map[string]string `protobuf:"bytes,9,rep,name=lables,proto3" json:"lables" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" form:"lables" bson:"lables"`
}

func (x *CreateRegistry) Reset() {
	*x = CreateRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_image_pb_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRegistry) ProtoMessage() {}

func (x *CreateRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_apps_image_pb_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRegistry.ProtoReflect.Descriptor instead.
func (*CreateRegistry) Descriptor() ([]byte, []int) {
	return file_apps_image_pb_image_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRegistry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRegistry) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateRegistry) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRegistry) GetLables() map[string]string {
	if x != nil {
		return x.Lables
	}
	return nil
}

var File_apps_image_pb_image_proto protoreflect.FileDescriptor

var file_apps_image_pb_image_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xe7, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4c,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_image_pb_image_proto_rawDescOnce sync.Once
	file_apps_image_pb_image_proto_rawDescData = file_apps_image_pb_image_proto_rawDesc
)

func file_apps_image_pb_image_proto_rawDescGZIP() []byte {
	file_apps_image_pb_image_proto_rawDescOnce.Do(func() {
		file_apps_image_pb_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_image_pb_image_proto_rawDescData)
	})
	return file_apps_image_pb_image_proto_rawDescData
}

var file_apps_image_pb_image_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_image_pb_image_proto_goTypes = []interface{}{
	(*Registry)(nil),       // 0: infraboard.mpaas.image.Registry
	(*CreateRegistry)(nil), // 1: infraboard.mpaas.image.CreateRegistry
	nil,                    // 2: infraboard.mpaas.image.CreateRegistry.LablesEntry
}
var file_apps_image_pb_image_proto_depIdxs = []int32{
	1, // 0: infraboard.mpaas.image.Registry.data:type_name -> infraboard.mpaas.image.CreateRegistry
	2, // 1: infraboard.mpaas.image.CreateRegistry.lables:type_name -> infraboard.mpaas.image.CreateRegistry.LablesEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apps_image_pb_image_proto_init() }
func file_apps_image_pb_image_proto_init() {
	if File_apps_image_pb_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_image_pb_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_apps_image_pb_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRegistry); i {
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
			RawDescriptor: file_apps_image_pb_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_image_pb_image_proto_goTypes,
		DependencyIndexes: file_apps_image_pb_image_proto_depIdxs,
		MessageInfos:      file_apps_image_pb_image_proto_msgTypes,
	}.Build()
	File_apps_image_pb_image_proto = out.File
	file_apps_image_pb_image_proto_rawDesc = nil
	file_apps_image_pb_image_proto_goTypes = nil
	file_apps_image_pb_image_proto_depIdxs = nil
}
