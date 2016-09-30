package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	beego.Info("Get ------------------------------------start")
	c.Data["Website"] = "b2c.cmcmall.com.cn"
	c.Data["Email"] = "majiango@hotmail.com"
	// c.TplName = "model.html"
	c.TplName = "index.html"
	beego.Info("Get -------------------------------------end")
}
