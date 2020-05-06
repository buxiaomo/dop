package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/buxiaomo/dop/models"
)

type AuthController struct {
	beego.Controller
}

var cpt *captcha.Captcha

type Captcha struct {
	// beego cache store
	store cache.Cache

	// url prefix for captcha image
	URLPrefix string

	// specify captcha id input field name
	FieldIdName string
	// specify captcha result input field name
	FieldCaptchaName string

	// captcha image width and height
	StdWidth  int
	StdHeight int

	// captcha chars nums
	ChallengeNums int

	// captcha expiration seconds
	Expiration int64

	// cache key prefix
	CachePrefix string
}

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 34
}

func (c *AuthController) Get() {

	c.Data["Title"] = "登录"
	c.TplName = "login.html"
}

type user struct {
	Id    		int    		`form:"-"`
	Username	string 		`form:"username"`
	Password   	string      `form:"password"`
}

func (c *AuthController) Post() {
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		// handle error
	}
	fmt.Println(u)

	var User models.User
	o := orm.NewOrm()
	if err := o.QueryTable(new(models.User)).Filter("username", u.Username).One(&User); err !=nil{
		fmt.Println(err.Error())
	}

	captcha_t := cpt.VerifyReq(c.Ctx.Request)
	fmt.Println(captcha_t)

	if !captcha_t {
		c.Ctx.Redirect(302,"/login")
		return
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