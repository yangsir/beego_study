package routers

import (
	"github.com/astaxie/beego"
	"myapp/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}
