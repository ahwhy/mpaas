package task

import (
	"fmt"
	"time"

	"github.com/infraboard/mpaas/apps/job"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
	"github.com/rs/xid"
)

func NewTaskSet() *JobTaskSet {
	return &JobTaskSet{
		Items: []*JobTask{},
	}
}

func (s *JobTaskSet) Add(task *JobTask) {
	s.Items = append(s.Items, task)
}

func NewDefaultJobTask() *JobTask {
	req := pipeline.NewRunJobRequest("")
	return NewJobTask(req)
}

func NewJobTask(req *pipeline.RunJobRequest) *JobTask {
	return &JobTask{
		Id:       xid.New().String(),
		CreateAt: time.Now().Unix(),
		Spec:     req,
		Job:      nil,
		Status:   NewJobTaskStatus(),
	}
}

func (t *JobTask) Update(job *job.Job, status *JobTaskStatus) {
	t.Job = job
	t.Status = status
}

func (s *JobTask) ShowTitle() string {
	return fmt.Sprintf("任务[%s]当前状态: %s", s.Spec.Job, s.Status.Stage.String())
}

func NewJobTaskStatus() *JobTaskStatus {
	return &JobTaskStatus{
		StartAt: time.Now().Unix(),
	}
}

var (
	COMPLETE_STGGE = []STAGE{STAGE_FAILED, STAGE_SUCCEEDED}
)

func (t *JobTaskStatus) IsComplete() bool {
	for _, s := range COMPLETE_STGGE {
		if t.Stage.Equal(s) {
			return true
		}
	}

	return false
}

func (t *JobTaskStatus) Update(req *UpdateJobTaskStatusRequest) {
	t.Stage = req.Stage
	t.Message = req.Message
	t.Detail = req.Detail

	// 结束时标记结束时间
	if req.Stage.IsIn(COMPLETE_STGGE...) {
		t.EndAt = time.Now().Unix()
	}
}
