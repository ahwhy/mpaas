package apps

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/infraboard/mpaas/apps/book/http"
	_ "github.com/infraboard/mpaas/apps/cluster/http"
)