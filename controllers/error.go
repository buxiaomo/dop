package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["Title"] = "page not found"
	c.Layout = "layout/layout.html"
	c.TplName = "error/404.html"
}

func (c *ErrorController) Error501() {
	c.Data["Title"] = "server error"
	c.Layout = "layout/layout.tpl"
	c.TplName = "error/501.tpl"
}