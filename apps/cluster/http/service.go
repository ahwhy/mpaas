package http

import (
	"io"

	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/provider/k8s"
	"sigs.k8s.io/yaml"

	v1 "k8s.io/api/core/v1"
)

func (h *handler) registryServiceHandler(ws *restful.WebService) {
	tags := []string{"服务管理"}

	ws.Route(ws.POST("/{id}/services").To(h.CreateService).
		Doc("创建服务").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(response.NewData(v1.Service{})).
		Returns(200, "OK", v1.Service{}))

	ws.Route(ws.GET("/{id}/services").To(h.QueryService).
		Doc("查询服务列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(response.NewData(v1.ServiceList{})).
		Returns(200, "OK", v1.ServiceList{}))

	ws.Route(ws.GET("/{id}/services/{name}").To(h.GetService).
		Doc("查询服务详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(cluster.QueryClusterRequest{}).
		Writes(response.NewData(v1.Service{})).
		Returns(200, "OK", v1.Service{}))
}

func (h *handler) CreateService(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := &v1.Service{}

	data, err := io.ReadAll(r.Request.Body)
	if err != nil {
		response.Failed(w, err)
		return
	}
	defer r.Request.Body.Close()

	if err := yaml.Unmarshal(data, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := client.CreateService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryService(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := k8s.NewListRequestFromHttp(r.Request)
	ins, err := client.ListService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) GetService(r *restful.Request, w *restful.Response) {
	client, err := h.GetClient(r.Request.Context(), r.PathParameter("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := k8s.NewGetRequestFromHttp(r.Request)
	req.Name = r.PathParameter("name")
	ins, err := client.GetService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}