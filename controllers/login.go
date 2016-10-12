package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (p *LoginController) Post() {
	beego.Info("LoginController ---------------------------------------------------start")
	var uname, upassword string
	uname = p.Input().Get("uname")
	upassword = p.Input().Get("upassword")
	beego.Info(uname)
	beego.Info(upassword)

	// models.Regionredis.Close()

	if uname == beego.AppConfig.String("userName") && upassword == beego.AppConfig.String("userPassword") {
		p.Data["uname"] = uname
		p.TplName = "index.html"
	} else {
		p.Data["msg"] = "用户名密码不正确。"
		p.TplName = "login.html"
	}

	beego.Info("LoginController ---------------------------------------------------end")
}
