// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: apps/artifact/pb/artifact.proto

package artifact

import (
	meta "github.com/infraboard/mpaas/common/meta"
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

type TYPE int32

const (
	// 容器镜像
	TYPE_IMAGE TYPE = 0
	// 安装包
	TYPE_PKG TYPE = 1
)

// Enum value maps for TYPE.
var (
	TYPE_name = map[int32]string{
		0: "IMAGE",
		1: "PKG",
	}
	TYPE_value = map[string]int32{
		"IMAGE": 0,
		"PKG":   1,
	}
)

func (x TYPE) Enum() *TYPE {
	p := new(TYPE)
	*p = x
	return p
}

func (x TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_artifact_pb_artifact_proto_enumTypes[0].Descriptor()
}

func (TYPE) Type() protoreflect.EnumType {
	return &file_apps_artifact_pb_artifact_proto_enumTypes[0]
}

func (x TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TYPE.Descriptor instead.
func (TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{0}
}

// 架构
type ARCH int32

const (
	ARCH_AMD64 ARCH = 0
	ARCH_ARM64 ARCH = 1
)

// Enum value maps for ARCH.
var (
	ARCH_name = map[int32]string{
		0: "AMD64",
		1: "ARM64",
	}
	ARCH_value = map[string]int32{
		"AMD64": 0,
		"ARM64": 1,
	}
)

func (x ARCH) Enum() *ARCH {
	p := new(ARCH)
	*p = x
	return p
}

func (x ARCH) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ARCH) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_artifact_pb_artifact_proto_enumTypes[1].Descriptor()
}

func (ARCH) Type() protoreflect.EnumType {
	return &file_apps_artifact_pb_artifact_proto_enumTypes[1]
}

func (x ARCH) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ARCH.Descriptor instead.
func (ARCH) EnumDescriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{1}
}

// 系统
type OS int32

const (
	OS_LINUX   OS = 0
	OS_WINDOWS OS = 1
	OS_MACOS   OS = 2
)

// Enum value maps for OS.
var (
	OS_name = map[int32]string{
		0: "LINUX",
		1: "WINDOWS",
		2: "MACOS",
	}
	OS_value = map[string]int32{
		"LINUX":   0,
		"WINDOWS": 1,
		"MACOS":   2,
	}
)

func (x OS) Enum() *OS {
	p := new(OS)
	*p = x
	return p
}

func (x OS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OS) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_artifact_pb_artifact_proto_enumTypes[2].Descriptor()
}

func (OS) Type() protoreflect.EnumType {
	return &file_apps_artifact_pb_artifact_proto_enumTypes[2]
}

func (x OS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OS.Descriptor instead.
func (OS) EnumDescriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{2}
}

type Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// 创建信息
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateArtifactRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// 产物质量
	// @gotags: bson:"quality" json:"quality"
	Quality *Quality `protobuf:"bytes,3,opt,name=quality,proto3" json:"quality" bson:"quality"`
}

func (x *Artifact) Reset() {
	*x = Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_artifact_pb_artifact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifact) ProtoMessage() {}

func (x *Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_apps_artifact_pb_artifact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifact.ProtoReflect.Descriptor instead.
func (*Artifact) Descriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *Artifact) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Artifact) GetSpec() *CreateArtifactRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Artifact) GetQuality() *Quality {
	if x != nil {
		return x.Quality
	}
	return nil
}

type CreateArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 那个服务的制品
	// @gotags: bson:"service_id" json:"service_id" validate:"required"
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id" bson:"service_id" validate:"required"`
	// 产物地址
	// @gotags: bson:"address" json:"address"
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address" bson:"address"`
	// 产物版本
	// @gotags: bson:"version" json:"version"
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version" bson:"version"`
	// 那个任务构建出的产物
	// @gotags: bson:"task_id" json:"task_id"
	TaskId string `protobuf:"bytes,4,opt,name=task_id,json=taskId,proto3" json:"task_id" bson:"task_id"`
	// 产物类型
	// @gotags: bson:"type" json:"type"
	Type TYPE `protobuf:"varint,5,opt,name=type,proto3,enum=infraboard.mpaas.artifact.TYPE" json:"type" bson:"type"`
	// 产物运行的cpu架构
	// @gotags: bson:"architecture" json:"architecture"
	Architecture ARCH `protobuf:"varint,6,opt,name=architecture,proto3,enum=infraboard.mpaas.artifact.ARCH" json:"architecture" bson:"architecture"`
	// 产物运行的操作系统
	// @gotags: bson:"os" json:"os"
	Os OS `protobuf:"varint,7,opt,name=os,proto3,enum=infraboard.mpaas.artifact.OS" json:"os" bson:"os"`
	// 产物内容Hash
	// @gotags: bson:"digest" json:"digest"
	Digest string `protobuf:"bytes,8,opt,name=digest,proto3" json:"digest" bson:"digest"`
	// 产物大小
	// @gotags: bson:"size" json:"size"
	Size int64 `protobuf:"varint,9,opt,name=size,proto3" json:"size" bson:"size"`
	// 标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *CreateArtifactRequest) Reset() {
	*x = CreateArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_artifact_pb_artifact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArtifactRequest) ProtoMessage() {}

func (x *CreateArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_artifact_pb_artifact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArtifactRequest.ProtoReflect.Descriptor instead.
func (*CreateArtifactRequest) Descriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArtifactRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *CreateArtifactRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateArtifactRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CreateArtifactRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *CreateArtifactRequest) GetType() TYPE {
	if x != nil {
		return x.Type
	}
	return TYPE_IMAGE
}

func (x *CreateArtifactRequest) GetArchitecture() ARCH {
	if x != nil {
		return x.Architecture
	}
	return ARCH_AMD64
}

func (x *CreateArtifactRequest) GetOs() OS {
	if x != nil {
		return x.Os
	}
	return OS_LINUX
}

func (x *CreateArtifactRequest) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *CreateArtifactRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CreateArtifactRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// 制品质量
type Quality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 单元测试质量
	// @gotags: bson:"unit_test" json:"unit_test"
	UnitTest *UnitTest `protobuf:"bytes,1,opt,name=unit_test,json=unitTest,proto3" json:"unit_test" bson:"unit_test"`
	// 接口测试质量
	// @gotags: bson:"api_test" json:"api_test"
	ApiTest *ApiTest `protobuf:"bytes,2,opt,name=api_test,json=apiTest,proto3" json:"api_test" bson:"api_test"`
	// 场景测试质量
	// @gotags: bson:"scenario_test" json:"scenario_test"
	ScenarioTest *ScenarioTest `protobuf:"bytes,3,opt,name=scenario_test,json=scenarioTest,proto3" json:"scenario_test" bson:"scenario_test"`
	// 安全测试
	// @gotags: bson:"security_test" json:"security_test"
	SecurityTest *SecurityTest `protobuf:"bytes,4,opt,name=security_test,json=securityTest,proto3" json:"security_test" bson:"security_test"`
	// 性能测试
	// @gotags: bson:"benchmark_test" json:"benchmark_test"
	BenchmarkTest *BarchmarkTest `protobuf:"bytes,5,opt,name=benchmark_test,json=benchmarkTest,proto3" json:"benchmark_test" bson:"benchmark_test"`
}

func (x *Quality) Reset() {
	*x = Quality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_artifact_pb_artifact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quality) ProtoMessage() {}

func (x *Quality) ProtoReflect() protoreflect.Message {
	mi := &file_apps_artifact_pb_artifact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quality.ProtoReflect.Descriptor instead.
func (*Quality) Descriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{2}
}

func (x *Quality) GetUnitTest() *UnitTest {
	if x != nil {
		return x.UnitTest
	}
	return nil
}

func (x *Quality) GetApiTest() *ApiTest {
	if x != nil {
		return x.ApiTest
	}
	return nil
}

func (x *Quality) GetScenarioTest() *ScenarioTest {
	if x != nil {
		return x.ScenarioTest
	}
	return nil
}

func (x *Quality) GetSecurityTest() *SecurityTest {
	if x != nil {
		return x.SecurityTest
	}
	return nil
}

func (x *Quality) GetBenchmarkTest() *BarchmarkTest {
	if x != nil {
		return x.BenchmarkTest
	}
	return nil
}

type UnitTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 代码覆盖率
	// @gotags: bson:"coverage" json:"coverage"
	Coverage float32 `protobuf:"fixed32,1,opt,name=coverage,proto3" json:"coverage" bson:"coverage"`
}

func (x *UnitTest) Reset() {
	*x = UnitTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_artifact_pb_artifact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitTest) ProtoMessage() {}

func (x *UnitTest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_artifact_pb_artifact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitTest.ProtoReflect.Descriptor instead.
func (*UnitTest) Descriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{3}
}

func (x *UnitTest) GetCoverage() float32 {
	if x != nil {
		return x.Coverage
	}
	return 0
}

type ApiTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 代码覆盖率
	// @gotags: bson:"coverage" json:"coverage"
	Coverage float32 `protobuf:"fixed32,1,opt,name=coverage,proto3" json:"coverage" bson:"coverage"`
}

func (x *ApiTest) Reset() {
	*x = ApiTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_artifact_pb_artifact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiTest) ProtoMessage() {}

func (x *ApiTest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_artifact_pb_artifact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiTest.ProtoReflect.Descriptor instead.
func (*ApiTest) Descriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{4}
}

func (x *ApiTest) GetCoverage() float32 {
	if x != nil {
		return x.Coverage
	}
	return 0
}

type ScenarioTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 代码覆盖率
	// @gotags: bson:"coverage" json:"coverage"
	Coverage float32 `protobuf:"fixed32,1,opt,name=coverage,proto3" json:"coverage" bson:"coverage"`
}

func (x *ScenarioTest) Reset() {
	*x = ScenarioTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_artifact_pb_artifact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScenarioTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScenarioTest) ProtoMessage() {}

func (x *ScenarioTest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_artifact_pb_artifact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScenarioTest.ProtoReflect.Descriptor instead.
func (*ScenarioTest) Descriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{5}
}

func (x *ScenarioTest) GetCoverage() float32 {
	if x != nil {
		return x.Coverage
	}
	return 0
}

type SecurityTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SecurityTest) Reset() {
	*x = SecurityTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_artifact_pb_artifact_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityTest) ProtoMessage() {}

func (x *SecurityTest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_artifact_pb_artifact_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityTest.ProtoReflect.Descriptor instead.
func (*SecurityTest) Descriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{6}
}

type BarchmarkTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BarchmarkTest) Reset() {
	*x = BarchmarkTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_artifact_pb_artifact_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarchmarkTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarchmarkTest) ProtoMessage() {}

func (x *BarchmarkTest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_artifact_pb_artifact_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarchmarkTest.ProtoReflect.Descriptor instead.
func (*BarchmarkTest) Descriptor() ([]byte, []int) {
	return file_apps_artifact_pb_artifact_proto_rawDescGZIP(), []int{7}
}

var File_apps_artifact_pb_artifact_proto protoreflect.FileDescriptor

var file_apps_artifact_pb_artifact_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f,
	0x70, 0x62, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x1a, 0x16, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x12, 0x36, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x3c, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xe9, 0x03,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x54, 0x59,
	0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x41, 0x52, 0x43, 0x48, 0x52,
	0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2d, 0x0a,
	0x02, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x4f, 0x53, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf7, 0x02, 0x0a, 0x07, 0x51, 0x75,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x08, 0x75,
	0x6e, 0x69, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x41, 0x70, 0x69, 0x54, 0x65, 0x73, 0x74, 0x52, 0x07, 0x61,
	0x70, 0x69, 0x54, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x54, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x42, 0x61, 0x72, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x54,
	0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x07, 0x41,
	0x70, 0x69, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x0e,
	0x0a, 0x0c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x65, 0x73, 0x74, 0x22, 0x0f,
	0x0a, 0x0d, 0x42, 0x61, 0x72, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x2a,
	0x1a, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4b, 0x47, 0x10, 0x01, 0x2a, 0x1c, 0x0a, 0x04, 0x41,
	0x52, 0x43, 0x48, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4d, 0x44, 0x36, 0x34, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x41, 0x52, 0x4d, 0x36, 0x34, 0x10, 0x01, 0x2a, 0x27, 0x0a, 0x02, 0x4f, 0x53, 0x12,
	0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x49,
	0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x43, 0x4f, 0x53,
	0x10, 0x02, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_artifact_pb_artifact_proto_rawDescOnce sync.Once
	file_apps_artifact_pb_artifact_proto_rawDescData = file_apps_artifact_pb_artifact_proto_rawDesc
)

func file_apps_artifact_pb_artifact_proto_rawDescGZIP() []byte {
	file_apps_artifact_pb_artifact_proto_rawDescOnce.Do(func() {
		file_apps_artifact_pb_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_artifact_pb_artifact_proto_rawDescData)
	})
	return file_apps_artifact_pb_artifact_proto_rawDescData
}

var file_apps_artifact_pb_artifact_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_apps_artifact_pb_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_artifact_pb_artifact_proto_goTypes = []interface{}{
	(TYPE)(0),                     // 0: infraboard.mpaas.artifact.TYPE
	(ARCH)(0),                     // 1: infraboard.mpaas.artifact.ARCH
	(OS)(0),                       // 2: infraboard.mpaas.artifact.OS
	(*Artifact)(nil),              // 3: infraboard.mpaas.artifact.Artifact
	(*CreateArtifactRequest)(nil), // 4: infraboard.mpaas.artifact.CreateArtifactRequest
	(*Quality)(nil),               // 5: infraboard.mpaas.artifact.Quality
	(*UnitTest)(nil),              // 6: infraboard.mpaas.artifact.UnitTest
	(*ApiTest)(nil),               // 7: infraboard.mpaas.artifact.ApiTest
	(*ScenarioTest)(nil),          // 8: infraboard.mpaas.artifact.ScenarioTest
	(*SecurityTest)(nil),          // 9: infraboard.mpaas.artifact.SecurityTest
	(*BarchmarkTest)(nil),         // 10: infraboard.mpaas.artifact.BarchmarkTest
	nil,                           // 11: infraboard.mpaas.artifact.CreateArtifactRequest.LabelsEntry
	(*meta.Meta)(nil),             // 12: infraboard.mpaas.common.meta.Meta
}
var file_apps_artifact_pb_artifact_proto_depIdxs = []int32{
	12, // 0: infraboard.mpaas.artifact.Artifact.meta:type_name -> infraboard.mpaas.common.meta.Meta
	4,  // 1: infraboard.mpaas.artifact.Artifact.spec:type_name -> infraboard.mpaas.artifact.CreateArtifactRequest
	5,  // 2: infraboard.mpaas.artifact.Artifact.quality:type_name -> infraboard.mpaas.artifact.Quality
	0,  // 3: infraboard.mpaas.artifact.CreateArtifactRequest.type:type_name -> infraboard.mpaas.artifact.TYPE
	1,  // 4: infraboard.mpaas.artifact.CreateArtifactRequest.architecture:type_name -> infraboard.mpaas.artifact.ARCH
	2,  // 5: infraboard.mpaas.artifact.CreateArtifactRequest.os:type_name -> infraboard.mpaas.artifact.OS
	11, // 6: infraboard.mpaas.artifact.CreateArtifactRequest.labels:type_name -> infraboard.mpaas.artifact.CreateArtifactRequest.LabelsEntry
	6,  // 7: infraboard.mpaas.artifact.Quality.unit_test:type_name -> infraboard.mpaas.artifact.UnitTest
	7,  // 8: infraboard.mpaas.artifact.Quality.api_test:type_name -> infraboard.mpaas.artifact.ApiTest
	8,  // 9: infraboard.mpaas.artifact.Quality.scenario_test:type_name -> infraboard.mpaas.artifact.ScenarioTest
	9,  // 10: infraboard.mpaas.artifact.Quality.security_test:type_name -> infraboard.mpaas.artifact.SecurityTest
	10, // 11: infraboard.mpaas.artifact.Quality.benchmark_test:type_name -> infraboard.mpaas.artifact.BarchmarkTest
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_apps_artifact_pb_artifact_proto_init() }
func file_apps_artifact_pb_artifact_proto_init() {
	if File_apps_artifact_pb_artifact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_artifact_pb_artifact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifact); i {
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
		file_apps_artifact_pb_artifact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArtifactRequest); i {
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
		file_apps_artifact_pb_artifact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quality); i {
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
		file_apps_artifact_pb_artifact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitTest); i {
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
		file_apps_artifact_pb_artifact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiTest); i {
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
		file_apps_artifact_pb_artifact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScenarioTest); i {
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
		file_apps_artifact_pb_artifact_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityTest); i {
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
		file_apps_artifact_pb_artifact_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarchmarkTest); i {
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
			RawDescriptor: file_apps_artifact_pb_artifact_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_artifact_pb_artifact_proto_goTypes,
		DependencyIndexes: file_apps_artifact_pb_artifact_proto_depIdxs,
		EnumInfos:         file_apps_artifact_pb_artifact_proto_enumTypes,
		MessageInfos:      file_apps_artifact_pb_artifact_proto_msgTypes,
	}.Build()
	File_apps_artifact_pb_artifact_proto = out.File
	file_apps_artifact_pb_artifact_proto_rawDesc = nil
	file_apps_artifact_pb_artifact_proto_goTypes = nil
	file_apps_artifact_pb_artifact_proto_depIdxs = nil
}
