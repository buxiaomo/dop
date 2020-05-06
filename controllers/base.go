package controllers

import (
	"github.com/astaxie/beego"
	"github.com/buxiaomo/dop/models"
	"time"
)

type BaseController struct {
	beego.Controller
	IsAdmin   bool
	UserName  string
	URL       string
	LastLogin	time.Time
	User      models.User
}

func (c *BaseController) Prepare() {
	// flash := beego.NewFlash()
	// Setting properties.
	c.StartSession()
	c.Data["PageStartTime"] = time.Now()
	user := c.GetSession("User")
	if user == nil {
		c.Ctx.Redirect(302,"/login")
	}
}