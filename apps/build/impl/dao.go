package impl

import (
	"github.com/infraboard/mpaas/apps/build"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *build.QueryBuildConfigRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*build.QueryBuildConfigRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
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

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Domain != "" {
		filter["domain"] = r.Domain
	}
	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}

	if len(r.ServiceIds) > 0 {
		filter["service_id"] = bson.M{"$in": r.ServiceIds}
	}
	if r.Event != "" {
		filter["condition.events"] = r.Event
	}
	if r.Branche != "" {
		filter["condition.branches"] = r.Branche
	}

	if r.Enabled != nil {
		filter["enabled"] = *r.Enabled
	}

	return filter
}
