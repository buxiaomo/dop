package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.Data["Title"] = "首页"
	c.Data["Username"] = c.GetSession("Username")
	c.Layout = "layout/layout.html"
	c.TplName = "index.html"
}