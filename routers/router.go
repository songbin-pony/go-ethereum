package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Index")
	beego.Router("/register", &controllers.MainController{},"get:Register")
	beego.Router("/get_accounts", &controllers.MainController{},"get:GetAccounts")
	beego.Router("/submit", &controllers.MainController{},"post:Submit")
	beego.Router("/registerSubmit", &controllers.MainController{},"post:RegisterSubmit")

}
