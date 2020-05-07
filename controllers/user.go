package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/buxiaomo/dop/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	c.Data["Title"] = "用户管理"
	var clusters []models.Cluster
	orm.NewOrm().QueryTable(new(models.Cluster)).All(&clusters)
	c.Data["data"] = clusters
	c.Layout = "layout/layout.html"
	c.TplName = "user.html"
}