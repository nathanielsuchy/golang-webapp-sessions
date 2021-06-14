package routers

import (
	"webapp/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sign-in", &controllers.SessionController{})
	beego.Router("/sign-out", &controllers.SignoutController{})
}
