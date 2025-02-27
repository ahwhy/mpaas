package impl_test

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl pipeline.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(pipeline.AppName).(pipeline.Service)
}
