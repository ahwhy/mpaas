// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: apps/approval/pb/approval.proto

package approval

import (
	job "github.com/infraboard/mpaas/apps/job"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
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

type STAGE int32

const (
	// 草稿
	STAGE_DRAFT STAGE = 0
	// 待审核
	STAGE_PENDDING STAGE = 1
	// 已过期, 如果一直没审核, 多久后会过期
	STAGE_EXPIRED STAGE = 2
	// 审核拒绝
	STAGE_DENY STAGE = 3
	// 审核通过
	STAGE_PASSED STAGE = 4
	// 关闭, 执行成功后,验证没问题, 流程结束
	STAGE_CLOSED STAGE = 15
)

// Enum value maps for STAGE.
var (
	STAGE_name = map[int32]string{
		0:  "DRAFT",
		1:  "PENDDING",
		2:  "EXPIRED",
		3:  "DENY",
		4:  "PASSED",
		15: "CLOSED",
	}
	STAGE_value = map[string]int32{
		"DRAFT":    0,
		"PENDDING": 1,
		"EXPIRED":  2,
		"DENY":     3,
		"PASSED":   4,
		"CLOSED":   15,
	}
)

func (x STAGE) Enum() *STAGE {
	p := new(STAGE)
	*p = x
	return p
}

func (x STAGE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STAGE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_approval_pb_approval_proto_enumTypes[0].Descriptor()
}

func (STAGE) Type() protoreflect.EnumType {
	return &file_apps_approval_pb_approval_proto_enumTypes[0]
}

func (x STAGE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STAGE.Descriptor instead.
func (STAGE) EnumDescriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{0}
}

type ApprovalSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 列表
	// @gotags: bson:"items" json:"items"
	Items []*Approval `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *ApprovalSet) Reset() {
	*x = ApprovalSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_approval_pb_approval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalSet) ProtoMessage() {}

func (x *ApprovalSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_approval_pb_approval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalSet.ProtoReflect.Descriptor instead.
func (*ApprovalSet) Descriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{0}
}

func (x *ApprovalSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ApprovalSet) GetItems() []*Approval {
	if x != nil {
		return x.Items
	}
	return nil
}

// 发布申请单
type Approval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// 创建信息
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateApprovalRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// 部署流水线配置
	// @gotags: bson:"-" json:"pipeline"
	Pipeline *pipeline.Pipeline `protobuf:"bytes,7,opt,name=pipeline,proto3" json:"pipeline" bson:"-"`
	// 发布单当前状态
	// @gotags: bson:"status" json:"status"
	Status *Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status" bson:"status"`
}

func (x *Approval) Reset() {
	*x = Approval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_approval_pb_approval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Approval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Approval) ProtoMessage() {}

func (x *Approval) ProtoReflect() protoreflect.Message {
	mi := &file_apps_approval_pb_approval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Approval.ProtoReflect.Descriptor instead.
func (*Approval) Descriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{1}
}

func (x *Approval) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Approval) GetSpec() *CreateApprovalRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Approval) GetPipeline() *pipeline.Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

func (x *Approval) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 创建发布申请
type CreateApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属域
	// @gotags: bson:"domain" json:"domain" validate:"required"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain" validate:"required"`
	// 所属空间
	// @gotags: bson:"namespace" json:"namespace" validate:"required"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace" validate:"required"`
	// 是否是模版
	// @gotags: bson:"is_template" json:"is_template"
	IsTemplate bool `protobuf:"varint,3,opt,name=is_template,json=isTemplate,proto3" json:"is_template" bson:"is_template"`
	// 申请人
	// @gotags: bson:"create_by" json:"create_by" validate:"required"
	CreateBy string `protobuf:"bytes,4,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by" validate:"required"`
	// 执行人列表, 申请通过后, 执行人负责流水线的执行
	// @gotags: bson:"operators" json:"operators"
	Operators []string `protobuf:"bytes,5,rep,name=operators,proto3" json:"operators" bson:"operators"`
	// 审核人列表
	// @gotags: bson:"auditors" json:"auditors"
	Auditors []string `protobuf:"bytes,6,rep,name=auditors,proto3" json:"auditors" bson:"auditors"`
	// 申请单标题
	// @gotags: bson:"title" json:"title" validate:"required"
	Title string `protobuf:"bytes,7,opt,name=title,proto3" json:"title" bson:"title" validate:"required"`
	// 申请说明, 支持Markdown
	// @gotags: bson:"describe" json:"describe" validate:"required"
	Describe string `protobuf:"bytes,8,opt,name=describe,proto3" json:"describe" bson:"describe" validate:"required"`
	// 审核通过后, 是否自动执行流水线, 默认审核通过后手动执行
	// @gotags: bson:"auto_run" json:"auto_run"
	AutoRun bool `protobuf:"varint,9,opt,name=auto_run,json=autoRun,proto3" json:"auto_run" bson:"auto_run"`
	// 流水线配置
	// @gotags: bson:"-" json:"pipeline_spec,omitempty"
	PipelineSpec *pipeline.CreatePipelineRequest `protobuf:"bytes,10,opt,name=pipeline_spec,json=pipelineSpec,proto3" json:"pipeline_spec,omitempty" bson:"-"`
	// 流水线Id
	// @gotags: bson:"pipeline_id" json:"pipeline_id"
	PipelineId string `protobuf:"bytes,11,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id" bson:"pipeline_id"`
	// 运行时需要传递的参数
	// @gotags: bson:"run_params" json:"run_params"
	RunParams []*job.RunParam `protobuf:"bytes,12,rep,name=run_params,json=runParams,proto3" json:"run_params" bson:"run_params"`
	// 申请单标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *CreateApprovalRequest) Reset() {
	*x = CreateApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_approval_pb_approval_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApprovalRequest) ProtoMessage() {}

func (x *CreateApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_approval_pb_approval_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApprovalRequest.ProtoReflect.Descriptor instead.
func (*CreateApprovalRequest) Descriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{2}
}

func (x *CreateApprovalRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateApprovalRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateApprovalRequest) GetIsTemplate() bool {
	if x != nil {
		return x.IsTemplate
	}
	return false
}

func (x *CreateApprovalRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *CreateApprovalRequest) GetOperators() []string {
	if x != nil {
		return x.Operators
	}
	return nil
}

func (x *CreateApprovalRequest) GetAuditors() []string {
	if x != nil {
		return x.Auditors
	}
	return nil
}

func (x *CreateApprovalRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateApprovalRequest) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *CreateApprovalRequest) GetAutoRun() bool {
	if x != nil {
		return x.AutoRun
	}
	return false
}

func (x *CreateApprovalRequest) GetPipelineSpec() *pipeline.CreatePipelineRequest {
	if x != nil {
		return x.PipelineSpec
	}
	return nil
}

func (x *CreateApprovalRequest) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

func (x *CreateApprovalRequest) GetRunParams() []*job.RunParam {
	if x != nil {
		return x.RunParams
	}
	return nil
}

func (x *CreateApprovalRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前状态
	// @gotags: bson:"stage" json:"stage"
	Stage STAGE `protobuf:"varint,1,opt,name=stage,proto3,enum=infraboard.mpaas.approval.STAGE" json:"stage" bson:"stage"`
	// 审核时间
	// @gotags: bson:"audit_at" json:"audit_at"
	AuditAt int64 `protobuf:"varint,2,opt,name=audit_at,json=auditAt,proto3" json:"audit_at" bson:"audit_at"`
	// 审核意见
	// @gotags: bson:"audit_comment" json:"audit_comment"
	AuditComment string `protobuf:"bytes,3,opt,name=audit_comment,json=auditComment,proto3" json:"audit_comment" bson:"audit_comment"`
	// 发布关闭的时间
	// @gotags: bson:"close_at" json:"close_at"
	CloseAt int64 `protobuf:"varint,12,opt,name=close_at,json=closeAt,proto3" json:"close_at" bson:"close_at"`
	// 关闭备注
	// @gotags: bson:"close_comment" json:"close_comment"
	CloseComment string `protobuf:"bytes,13,opt,name=close_comment,json=closeComment,proto3" json:"close_comment" bson:"close_comment"`
	// 通知记录
	// @gotags: bson:"notify_records" json:"notify_records"
	NotifyRecords []*NotifyRecord `protobuf:"bytes,14,rep,name=notify_records,json=notifyRecords,proto3" json:"notify_records" bson:"notify_records"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_approval_pb_approval_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_apps_approval_pb_approval_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetStage() STAGE {
	if x != nil {
		return x.Stage
	}
	return STAGE_DRAFT
}

func (x *Status) GetAuditAt() int64 {
	if x != nil {
		return x.AuditAt
	}
	return 0
}

func (x *Status) GetAuditComment() string {
	if x != nil {
		return x.AuditComment
	}
	return ""
}

func (x *Status) GetCloseAt() int64 {
	if x != nil {
		return x.CloseAt
	}
	return 0
}

func (x *Status) GetCloseComment() string {
	if x != nil {
		return x.CloseComment
	}
	return ""
}

func (x *Status) GetNotifyRecords() []*NotifyRecord {
	if x != nil {
		return x.NotifyRecords
	}
	return nil
}

type NotifyRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 那种状态的通知
	// @gotags: bson:"stage" json:"stage"
	Stage STAGE `protobuf:"varint,1,opt,name=stage,proto3,enum=infraboard.mpaas.approval.STAGE" json:"stage" bson:"stage"`
	// 通知时间
	// @gotags: bson:"notify_at" json:"notify_at"
	NotifyAt int64 `protobuf:"varint,2,opt,name=notify_at,json=notifyAt,proto3" json:"notify_at" bson:"notify_at"`
	// 是否通知成功
	// @gotags: bson:"is_success" json:"is_success"
	IsSuccess bool `protobuf:"varint,3,opt,name=is_success,json=isSuccess,proto3" json:"is_success" bson:"is_success"`
	// 通知失败的原因
	// @gotags: bson:"message" json:"message"
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message" bson:"message"`
	// 具体详情, 通过成功后的响应数据, Debug使用
	// @gotags: bson:"detail" json:"detail"
	Detail string `protobuf:"bytes,5,opt,name=detail,proto3" json:"detail" bson:"detail"`
}

func (x *NotifyRecord) Reset() {
	*x = NotifyRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_approval_pb_approval_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRecord) ProtoMessage() {}

func (x *NotifyRecord) ProtoReflect() protoreflect.Message {
	mi := &file_apps_approval_pb_approval_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRecord.ProtoReflect.Descriptor instead.
func (*NotifyRecord) Descriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{4}
}

func (x *NotifyRecord) GetStage() STAGE {
	if x != nil {
		return x.Stage
	}
	return STAGE_DRAFT
}

func (x *NotifyRecord) GetNotifyAt() int64 {
	if x != nil {
		return x.NotifyAt
	}
	return 0
}

func (x *NotifyRecord) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *NotifyRecord) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NotifyRecord) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

var File_apps_approval_pb_approval_proto protoreflect.FileDescriptor

var file_apps_approval_pb_approval_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2f,
	0x70, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x1a, 0x16, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2f,
	0x70, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0b,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x84, 0x02, 0x0a,
	0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3f, 0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0xda, 0x04, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x75, 0x6e, 0x12, 0x55, 0x0a, 0x0d, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x54, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x90, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x53, 0x54, 0x41, 0x47, 0x45, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x75, 0x64, 0x69, 0x74, 0x41, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e,
	0x53, 0x54, 0x41, 0x47, 0x45, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2a, 0x4f, 0x0a, 0x05, 0x53, 0x54,
	0x41, 0x47, 0x45, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x45, 0x4e, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e,
	0x59, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x0f, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_approval_pb_approval_proto_rawDescOnce sync.Once
	file_apps_approval_pb_approval_proto_rawDescData = file_apps_approval_pb_approval_proto_rawDesc
)

func file_apps_approval_pb_approval_proto_rawDescGZIP() []byte {
	file_apps_approval_pb_approval_proto_rawDescOnce.Do(func() {
		file_apps_approval_pb_approval_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_approval_pb_approval_proto_rawDescData)
	})
	return file_apps_approval_pb_approval_proto_rawDescData
}

var file_apps_approval_pb_approval_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_approval_pb_approval_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_approval_pb_approval_proto_goTypes = []interface{}{
	(STAGE)(0),                             // 0: infraboard.mpaas.approval.STAGE
	(*ApprovalSet)(nil),                    // 1: infraboard.mpaas.approval.ApprovalSet
	(*Approval)(nil),                       // 2: infraboard.mpaas.approval.Approval
	(*CreateApprovalRequest)(nil),          // 3: infraboard.mpaas.approval.CreateApprovalRequest
	(*Status)(nil),                         // 4: infraboard.mpaas.approval.Status
	(*NotifyRecord)(nil),                   // 5: infraboard.mpaas.approval.NotifyRecord
	nil,                                    // 6: infraboard.mpaas.approval.CreateApprovalRequest.LabelsEntry
	(*meta.Meta)(nil),                      // 7: infraboard.mpaas.common.meta.Meta
	(*pipeline.Pipeline)(nil),              // 8: infraboard.mpaas.pipeline.Pipeline
	(*pipeline.CreatePipelineRequest)(nil), // 9: infraboard.mpaas.pipeline.CreatePipelineRequest
	(*job.RunParam)(nil),                   // 10: infraboard.mpaas.job.RunParam
}
var file_apps_approval_pb_approval_proto_depIdxs = []int32{
	2,  // 0: infraboard.mpaas.approval.ApprovalSet.items:type_name -> infraboard.mpaas.approval.Approval
	7,  // 1: infraboard.mpaas.approval.Approval.meta:type_name -> infraboard.mpaas.common.meta.Meta
	3,  // 2: infraboard.mpaas.approval.Approval.spec:type_name -> infraboard.mpaas.approval.CreateApprovalRequest
	8,  // 3: infraboard.mpaas.approval.Approval.pipeline:type_name -> infraboard.mpaas.pipeline.Pipeline
	4,  // 4: infraboard.mpaas.approval.Approval.status:type_name -> infraboard.mpaas.approval.Status
	9,  // 5: infraboard.mpaas.approval.CreateApprovalRequest.pipeline_spec:type_name -> infraboard.mpaas.pipeline.CreatePipelineRequest
	10, // 6: infraboard.mpaas.approval.CreateApprovalRequest.run_params:type_name -> infraboard.mpaas.job.RunParam
	6,  // 7: infraboard.mpaas.approval.CreateApprovalRequest.labels:type_name -> infraboard.mpaas.approval.CreateApprovalRequest.LabelsEntry
	0,  // 8: infraboard.mpaas.approval.Status.stage:type_name -> infraboard.mpaas.approval.STAGE
	5,  // 9: infraboard.mpaas.approval.Status.notify_records:type_name -> infraboard.mpaas.approval.NotifyRecord
	0,  // 10: infraboard.mpaas.approval.NotifyRecord.stage:type_name -> infraboard.mpaas.approval.STAGE
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_apps_approval_pb_approval_proto_init() }
func file_apps_approval_pb_approval_proto_init() {
	if File_apps_approval_pb_approval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_approval_pb_approval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalSet); i {
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
		file_apps_approval_pb_approval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Approval); i {
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
		file_apps_approval_pb_approval_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApprovalRequest); i {
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
		file_apps_approval_pb_approval_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_apps_approval_pb_approval_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRecord); i {
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
			RawDescriptor: file_apps_approval_pb_approval_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_approval_pb_approval_proto_goTypes,
		DependencyIndexes: file_apps_approval_pb_approval_proto_depIdxs,
		EnumInfos:         file_apps_approval_pb_approval_proto_enumTypes,
		MessageInfos:      file_apps_approval_pb_approval_proto_msgTypes,
	}.Build()
	File_apps_approval_pb_approval_proto = out.File
	file_apps_approval_pb_approval_proto_rawDesc = nil
	file_apps_approval_pb_approval_proto_goTypes = nil
	file_apps_approval_pb_approval_proto_depIdxs = nil
}
