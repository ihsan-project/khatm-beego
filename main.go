package main

import (
	"controllers/controllers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/api/version/", &controllers.VersionController{})
	beego.Run()
}
