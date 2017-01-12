package routers

import (
	"github.com/astaxie/beego"
	"insta-profile-downloader/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
