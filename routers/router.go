package routers

import (
	"github.com/astaxie/beego"
	"github.com/buxiaomo/dop/controllers"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/cluster", &controllers.ClusterController{})
	beego.Router("/project", &controllers.ProjectController{})
	beego.Router("/login", &controllers.UserController{})
}