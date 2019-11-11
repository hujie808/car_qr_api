package routers

import (
	"github.com/astaxie/beego"
	"web_bee/car_qr_api/extra_apps/gernate_qr/controllers"
)

func init() {

	ns := beego.NewNamespace("/v2",

		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.Istest{},
			),
		),
	)
	beego.AddNamespace(ns)
}
