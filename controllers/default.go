package controllers

import (
	"github.com/astaxie/beego"
	"insta-profile-downloader/api"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	l := api.GetPhotos("https://www.instagram.com/shubhothegreat")
	c.Data["Count"] = l.Count
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
