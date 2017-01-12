package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"insta-profile-downloader/api"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	l := api.GetPhotos()
	fmt.Println(l)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
