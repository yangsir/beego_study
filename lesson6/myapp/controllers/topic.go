package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["isActive"] = "topic"
	this.TplName = "topic.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	category := this.Input().Get("category")
	lable := this.Input().Get("lable")
	content := this.Input().Get("content")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, lable, content)
	} else {
		err = models.ModifyTopic(tid, title, category, lable, content)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.Data["isActive"] = ""
	this.TplName = "topic_add.html"
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"
	this.Data["isActive"] = ""

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	this.Data["isActive"] = ""

	tid := this.Ctx.Input.Params()["0"]
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Lables"] = strings.Split(topic.Lables, " ")

	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)

}
