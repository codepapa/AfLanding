package routers

import (
	"github.com/astaxie/beego"
	"github.com/codepapa/AfLanding/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
