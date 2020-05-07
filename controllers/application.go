package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/buxiaomo/dop/models"
)

type ApplicationController struct {
	BaseController
}

func (c *ApplicationController) Get() {
	c.Data["Title"] = "应用管理"
	var clusters []models.Cluster
	orm.NewOrm().QueryTable(new(models.Cluster)).All(&clusters)
	c.Data["data"] = clusters
	c.Layout = "layout/layout.html"
	c.TplName = "application.html"
}