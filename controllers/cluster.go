package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/buxiaomo/dop/models"
)

type ClusterController struct {
	BaseController
}

func (c *ClusterController) Get() {
	c.Data["Title"] = "集群管理"
	var clusters []models.Cluster
	orm.NewOrm().QueryTable(new(models.Cluster)).All(&clusters)
	c.Data["data"] = clusters
	c.Layout = "layout/layout.html"
	c.TplName = "cluster.html"
}