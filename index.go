package main

import (
	"github.com/astaxie/beego"
)

type VersionController struct {
	beego.Controller
}

func (this *VersionController) Get() {
	this.Ctx.WriteString("1.0")
}

// func (this *MainController) Get() {
// 	this.Ctx.WriteString("bye!")
// }

func main() {
	beego.Router("/api/version/", &VersionController{})
	beego.Run()
}
