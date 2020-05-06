package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/buxiaomo/dop/models"
)

type AuthController struct {
	beego.Controller
}

type user struct {
	Id    		int    		`form:"-"`
	Username	string 		`form:"username"`
	Password   	string      `form:"password"`
}

func (c *AuthController) Get() {
	c.Data["Title"] = "登录"
	c.TplName = "login.html"
}

func (c *AuthController) Post() {
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		// handle error
	}
	var User models.User
	o := orm.NewOrm()
	if err := o.QueryTable(new(models.User)).Filter("username", u.Username).One(&User); err !=nil{
		fmt.Println(err.Error())
	}
	if User.Password == u.Password {
		//c.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
		c.SetSession("User", map[string]interface{}{
			"Id": User.Id,
			"Username": User.Username,
			"IsAdmin": User.IsAdmin,
			"Email": User.Email,
		})
		c.SetSession("Username", User.Username)
		c.Ctx.Redirect(302,"/index")
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
		c.Ctx.Redirect(302,"/login")
	}
}

func (c *AuthController) Exit() {
	c.Data["Title"] = "登录"
	c.DelSession("User")
	c.Ctx.Redirect(302,"/login")
}