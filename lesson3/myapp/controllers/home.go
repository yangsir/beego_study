package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["isActive"] = "home"
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}
