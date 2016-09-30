package main

import (
	"github.com/astaxie/beego"
	_ "redisRegionWebServer/routers"
)

func main() {
	beego.Run()
}
