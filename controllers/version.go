package controllers

import (
	"github.com/astaxie/beego"
)

type VersionController struct {
	beego.Controller
}

func (this *VersionController) Get() {
	this.Ctx.WriteString("1.0")
}
