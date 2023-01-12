package batch

import (
	"context"

	"github.com/infraboard/mpaas/provider/k8s/meta"
	v1 "k8s.io/api/batch/v1"
)

func (b *Batch) ListCronJob(ctx context.Context, req *meta.ListRequest) (*v1.CronJobList, error) {
	return b.batchV1.CronJobs(req.Namespace).List(ctx, req.Opts)
}
