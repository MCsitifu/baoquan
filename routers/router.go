package routers

import (
	"DateCertproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user_register",&controllers.RegisterController{})

	beego.Router("/login.html",&controllers.LoginController{})

	beego.Router( "/user_login",&controllers.LoginController{})
}
