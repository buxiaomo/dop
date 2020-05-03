package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type user struct {
	Id    		int    		`form:"-"`
	Username	string 		`form:"username"`
	Password   	string      `form:"password"`
}

func (c *UserController) Get() {
	c.Data["Title"] = "登录"

	c.TplName = "login.html"
}

func (c *UserController) Post() {
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		// handle error
	}
	fmt.Println(u)
	c.Data["Title"] = "登录"
	c.Layout = "layout/layout.html"
	c.TplName = "login.html"
}