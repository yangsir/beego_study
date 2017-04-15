package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["isActive"] = "home"
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(this.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Categories"] = categories
}
