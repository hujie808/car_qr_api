package routers

import (
	"github.com/astaxie/beego"
	"web_bee/car_qr_api/extra_apps/ad_customers/controllers"
)

func init() {

	ns := beego.NewNamespace("/ad",

		beego.NSInclude(
			&controllers.LoginController{},
			&controllers.Index{},
		),

	)
	beego.AddNamespace(ns)
}
