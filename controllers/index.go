package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["Title"] = "首页"
	c.Layout = "layout/layout.html"
	c.TplName = "index.html"
}