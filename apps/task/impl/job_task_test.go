package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestQueryJobTask(t *testing.T) {
	req := task.NewQueryTaskRequest()
	req.PipelineTaskId = conf.C.PIPELINE_TASK_ID
	set, err := impl.QueryJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tools.MustToJson(set))
}

func TestRunBuildJob(t *testing.T) {
	req := pipeline.NewRunJobRequest("docker_build@default.default")
	// 添加飞书通知的Webhook
	req.AddWebhook(pipeline.NewWebHook(conf.C.FEISHU_BOT_URL))
	// 添加任务执行成功提醒
	req.AddMentionUser(task.NewMentionUser("admin", notify.NOTIFY_TYPE_IM))
	// 添加参数
	version := job.NewVersionedRunParam(job.LATEST_VERSION_NAME)
	version.Params = job.NewRunParamWithKVPaire(
		"GIT_SSH_URL", "git@github.com:infraboard/mcenter.git",
		"GIT_BRANCH", "master",
		"GIT_COMMIT_ID", "06ca94c2c535d6a25643f40cacfaa8a76abf1592",
		build.SYSTEM_VARIABLE_IMAGE_REPOSITORY, "registry.cn-hangzhou.aliyuncs.com/infraboard/mcenter",
		build.SYSTEM_VARIABLE_APP_VERSION, "v0.0.6",
	)
	req.RunParams = version
	ins, err := impl.RunJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestRunDeployJob(t *testing.T) {
	req := pipeline.NewRunJobRequest("docker_deploy@default.default")
	version := job.NewVersionedRunParam(job.LATEST_VERSION_NAME)
	version.Params = job.NewRunParamWithKVPaire(
		job.SYSTEM_VARIABLE_DEPLOY_ID, conf.C.DEPLOY_ID,
		build.SYSTEM_VARIABLE_APP_VERSION, "1.30",
	)
	req.RunParams = version
	req.DryRun = false

	ins, err := impl.RunJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins.Status.Detail))
}

func TestUpdateJobTaskOutput(t *testing.T) {
	req := task.NewUpdateJobTaskOutputRequest(conf.C.MCENTER_BUILD_TASK_ID)
	req.UpdateToken = conf.C.MCENTER_BUILD_TASK_ID
	req.AddRuntimeEnv(build.SYSTEM_VARIABLE_APP_VERSION, "v0.0.5")
	req.MarkdownOutput = "构建产物描述信息"
	ins, err := impl.UpdateJobTaskOutput(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestUpdateJobTaskStatus(t *testing.T) {
	req := task.NewUpdateJobTaskStatusRequest(conf.C.MCENTER_BUILD_TASK_ID)
	req.Stage = task.STAGE_SUCCEEDED
	req.Message = "执行成功"
	req.ForceUpdateStatus = true
	req.ForceTriggerPipeline = true
	req.UpdateToken = conf.C.MCENTER_BUILD_TASK_TOKEN
	req.Detail = tools.MustReadContentFile("test/k8s_build_job.yml")
	ins, err := impl.UpdateJobTaskStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestDescribeJobTask(t *testing.T) {
	req := task.NewDescribeJobTaskRequest(conf.C.MCENTER_BUILD_TASK_ID)
	ins, err := impl.DescribeJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDeleteJobTask(t *testing.T) {
	req := task.NewDeleteJobTaskRequest(conf.C.MCENTER_BUILD_TASK_ID)
	req.Force = true
	set, err := impl.DeleteJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
