package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"myapp/controllers"
	"myapp/models"
	_ "myapp/routers"
	"os"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 注册 beego 路由
	//beego.Router("/", &controllers.HomeController{})

	// 附件处理
	os.Mkdir("attachment", os.ModePerm)

	// 启动 beego
	beego.Run()
}
