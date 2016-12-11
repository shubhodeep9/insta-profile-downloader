package routers

import (
	"insta-profile-downloader/insta-profile-downloader/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
