package rest

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mpaas/apps/cluster"
)

func (c *ClientSet) CreateCluster(ctx context.Context, req *cluster.CreateClusterRequest) (
	*cluster.Cluster, error) {
	ins := cluster.NewDefaultCluster()

	err := c.c.
		Post("clusters").
		Body(req).
		Do(ctx).
		Into(response.NewData(ins))
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (c *ClientSet) QueryCluster(ctx context.Context, req *cluster.QueryClusterRequest) (
	*cluster.ClusterSet, error) {
	set := cluster.NewClusterSet()

	err := c.c.
		Get("clusters").
		Body(req).
		Do(ctx).
		Into(response.NewData(set))
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (c *ClientSet) DescribeCluster(ctx context.Context, req *cluster.DescribeClusterRequest) (
	*cluster.Cluster, error) {
	ins := cluster.NewDefaultCluster()

	err := c.c.
		Get(fmt.Sprintf("clusters/%s", req.Id)).
		Body(req).
		Do(ctx).
		Into(response.NewData(ins))
	if err != nil {
		return nil, err
	}

	return ins, nil
}
