package main

import (
	"controllers/controllers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/api/version/", &controllers.VersionController{})
	beego.Router("/api/authorizations/", &controllers.AuthorizationController{})

	// beego.BConfig.CopyRequestBody = true

	beego.Run()
}
