package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type AuthorizationController struct {
	beego.Controller
}

type MyStruct struct {
	payload string
}

func (this *AuthorizationController) Post() {
	fmt.Println("authorizing: %d", time.Now())
	// fmt.Println(this.Ctx.Input.RequestBody)

	// u := MyStruct{}
	// t := this.ParseForm(&u)
	// fmt.Println(t)

	jsoninfo := this.GetString("search_series")
	fmt.Println(jsoninfo)

	// this.Ctx.WriteString("done!")

	this.Abort("400")
}
