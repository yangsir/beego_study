package routers

import (
	"github.com/astaxie/beego"
	"myapp/controllers"
	"os"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

	// 附件处理
	os.Mkdir("attachment", os.ModePerm)
	//方法一：
	//beego.SetStaticPath("/attachment", "attachment")

	//方法二
	beego.Router("/attachment/:all", &controllers.AttachController{})
}
