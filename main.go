package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/buxiaomo/dop/controllers"
	_ "github.com/buxiaomo/dop/models"
	_ "github.com/buxiaomo/dop/routers"
	"os"
)

func init(){
	if err := os.MkdirAll("kubeconfig", 0777); err != nil {
		panic("failed to create Session directory" + err.Error())
	}
}
func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}