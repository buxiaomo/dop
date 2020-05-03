package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/buxiaomo/dop/models"
)

type ProjectController struct {
	beego.Controller
}

func (c *ProjectController) Get() {
	c.Data["Title"] = "项目管理"
	var clusters []models.Cluster
	orm.NewOrm().QueryTable(new(models.Cluster)).All(&clusters)
	c.Data["data"] = clusters
	c.Layout = "layout/layout.html"
	c.TplName = "project.html"
}