package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type VersionController struct {
	beego.Controller
}

func (this *VersionController) Get() {
	fmt.Println("hello: %d", time.Now())
	this.Ctx.WriteString("1.0")
}
