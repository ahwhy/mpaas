package impl

import (
	"context"
	"time"

	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"github.com/infraboard/mpaas/apps/approval"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 创建发布申请
func (i *impl) CreateApproval(ctx context.Context, in *approval.CreateApprovalRequest) (
	*approval.Approval, error) {
	ins, err := approval.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 补充Pipeline创建
	if in.DeployPipelineSpec != nil {
		in.DeployPipelineSpec.AddLabel(approval.APPROVAL_LABEL_KEY, ins.Meta.Id)
		p, err := i.pipeline.CreatePipeline(ctx, in.DeployPipelineSpec)
		if err != nil {
			return nil, err
		}

		ins.Spec.DeployPipelineId = p.Meta.Id
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a approval document error, %s", err)
	}
	return nil, nil
}

// 查询发布申请列表
func (i *impl) QueryApproval(ctx context.Context, in *approval.QueryApprovalRequest) (
	*approval.ApprovalSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find deploy error, error is %s", err)
	}

	set := approval.NewApprovalSet()
	// 循环
	for resp.Next(ctx) {
		ins := approval.NewDefaultApproval()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode deploy error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get deploy count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

// 查询发布申请详情
func (i *impl) DescribeApproval(ctx context.Context, in *approval.DescribeApprovalRequest) (
	*approval.Approval, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := approval.NewDefaultApproval()
	if err := i.col.FindOne(ctx, bson.M{"_id": in.Id}).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("approval %s not found", in)
		}

		return nil, exception.NewInternalServerError("find approval %s error, %s", in.Id, err)
	}

	return ins, nil
}

// 编辑发布申请
func (i *impl) EditApproval(ctx context.Context, in *approval.EditApprovalRequest) (
	*approval.Approval, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := approval.NewDescribeApprovalRequest(in.Id)
	ins, err := i.DescribeApproval(ctx, req)
	if err != nil {
		return nil, err
	}

	// 1. 只有处于草稿状态的申请才允许编辑
	if !ins.Status.Stage.Equal(approval.STAGE_DRAFT) {
		if err != nil {
			return nil, exception.NewBadRequest("只有处于草稿状态的发布申请才能编辑")
		}
	}

	switch in.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Spec = in.Spec
	case request.UpdateMode_PATCH:
		if err := mergo.MergeWithOverwrite(ins.Spec, in.Spec); err != nil {
			return nil, err
		}
		if err := ins.Spec.Validate(); err != nil {
			return nil, err
		}
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", in.UpdateMode)
	}

	// 校验更新后请求合法性
	if err := ins.Spec.Validate(); err != nil {
		return nil, err
	}

	ins.Meta.UpdateAt = time.Now().Unix()
	_, err = i.col.UpdateOne(ctx, bson.M{"_id": ins.Meta.Id}, bson.M{"$set": ins})
	if err != nil {
		return nil, exception.NewInternalServerError("update approval(%s) error, %s", ins.Meta.Id, err)
	}

	return ins, nil
}

// 更新发布申请状态
func (i *impl) UpdateApprovalStatus(ctx context.Context, in *approval.UpdateApprovalStatusRequest) (
	*approval.Approval, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := approval.NewDescribeApprovalRequest(in.Id)
	ins, err := i.DescribeApproval(ctx, req)
	if err != nil {
		return nil, err
	}

	// 1. 只有处于草稿状态的申请才允许编辑
	if !ins.Status.Stage.Equal(approval.STAGE_CLOSED) {
		if err != nil {
			return nil, exception.NewBadRequest("发布申请已关闭, 禁止更新状态")
		}
	}

	// 3. 保存更新
	ins.Status.Update(in.Status.Stage)
	_, err = i.col.UpdateOne(ctx, bson.M{"_id": ins.Meta.Id}, bson.M{"$set": bson.M{"status": in.Status}})
	if err != nil {
		return nil, exception.NewInternalServerError("update approval(%s) error, %s", ins.Meta.Id, err)
	}
	return ins, nil
}