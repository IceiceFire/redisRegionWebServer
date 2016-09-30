package routers

import (
	"github.com/astaxie/beego"
	"redisRegionWebServer/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/region/", &controllers.PostController{})
}
