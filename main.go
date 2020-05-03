package main

import (
	"github.com/astaxie/beego"
	"github.com/buxiaomo/dop/controllers"
	_ "github.com/buxiaomo/dop/models"
	_ "github.com/buxiaomo/dop/routers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}