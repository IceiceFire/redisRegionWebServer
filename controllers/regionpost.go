package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"redisRegionWebServer/models"
)

type PostController struct {
	beego.Controller
}

func (p *PostController) Post() {
	beego.Info("Post ---------------------------------------------------start")

	req := struct{ Region, Redisip, Keyprefix, Expire, Starting string }{}
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &req); err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("empty title"))
		return
	}

	jsonval := "{ \"Redisip\": " + "\"" + req.Redisip + "\", " +
		"\"Keyprefix\": " + "\"" + req.Keyprefix + "\", " +
		"\"Expire\": " + "\"" + req.Expire + "\", " +
		"\"Starting\": " + "\"" + req.Starting + "\" " +
		"}"

	models.Regionredis.Setkey(req.Region, jsonval)

	p.TplName = "index.html"
	beego.Info("Post ---------------------------------------------------end")
}

func (p *PostController) Get() {

	p.TplName = "login.html"
}
