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

var count int = 0

func (this *AuthorizationController) Post() {
	fmt.Println("authorizing: %d", time.Now())
	// fmt.Println(this.Ctx.Input.RequestBody)

	// u := MyStruct{}
	// t := this.ParseForm(&u)
	// fmt.Println(t)

	jsoninfo := this.GetString("search_series")
	fmt.Println(jsoninfo)

	// Respond 200
	count++

	// if count < 2 {
	// 	fmt.Println("sending 400", count)
	this.Abort("400")
	// } else {
	// 	fmt.Println("sending 200", count)
	// 	this.Ctx.WriteString("done!")
	// }
}
