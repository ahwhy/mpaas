package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mpaas/apps/deploy"
)

func (h *downloadHandler) Registry(ws *restful.WebService) {
	tags := []string{"部署配置下载"}
	ws.Route(ws.GET("/deploy_configs/{id}").To(h.DownloadDeployConfig).
		Doc("下载配置详情").
		Param(ws.PathParameter("id", "identifier of the deploy").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Writes(deploy.DeployConfig{}).
		Returns(200, "OK", deploy.DeployConfig{}).
		Returns(404, "Not Found", nil))
}

func (h *downloadHandler) DownloadDeployConfig(r *restful.Request, w *restful.Response) {
	req := deploy.NewDescribeDeployConfigRequest(r.PathParameter("id"))

	// 查询部署配置
	ins, err := h.service.DescribeDeployConfig(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 校验集群访问Token
	if ins.Spec.AuthEnabled {
		err = ins.ValidateToken(r.QueryParameter("token"))
		if err != nil {
			response.Failed(w, err)
			return
		}
	}

	switch ins.Spec.Type {
	case deploy.TYPE_HOST:
	case deploy.TYPE_KUBERNETES:
		w.Write([]byte(ins.Spec.K8STypeConfig.WorkloadConfig))
	}
}
