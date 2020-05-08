package routers

import (
	"github.com/astaxie/beego"
	"github.com/buxiaomo/dop/controllers"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{})

	beego.Router("/cluster", &controllers.ClusterController{})
	beego.Router("/cluster/list", &controllers.ClusterController{}, "get:List")
	beego.Router("/cluster/delete/:id:int", &controllers.ClusterController{}, "get:Delete")
	//beego.Router("/cluster/info/:id:int", &controllers.ClusterController{}, "get:Info")

	beego.Router("/project", &controllers.ProjectController{})
	beego.Router("/application", &controllers.ApplicationController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/login", &controllers.AuthController{})
	beego.Router("/logout", &controllers.AuthController{},"get:Exit")
}