package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/task"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryPipelineTaskRequest(r *task.QueryPipelineTaskRequest) *queryPipelineTaskRequest {
	return &queryPipelineTaskRequest{
		r,
	}
}

type queryPipelineTaskRequest struct {
	*task.QueryPipelineTaskRequest
}

func (r *queryPipelineTaskRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort: bson.D{
			{Key: "create_at", Value: -1},
		},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryPipelineTaskRequest) FindFilter() bson.M {
	filter := bson.M{}

	if len(r.Ids) > 0 {
		filter["_id"] = bson.M{"$in": r.Ids}
	}

	if r.PipelineId != "" {
		filter["pipeline._id"] = r.PipelineId
	}

	if len(r.Stages) > 0 {
		filter["status.stage"] = bson.M{"$in": r.Stages}
	}

	return filter
}

func (i *impl) deletecluster(ctx context.Context, ins *task.PipelineTask) error {
	if ins == nil || ins.Meta.Id == "" {
		return fmt.Errorf("pipeline is nil")
	}

	result, err := i.pcol.DeleteOne(ctx, bson.M{"_id": ins.Meta.Id})
	if err != nil {
		return exception.NewInternalServerError("delete pipeline task(%s) error, %s", ins.Meta.Id, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("pipeline task %s not found", ins.Meta.Id)
	}

	return nil
}
