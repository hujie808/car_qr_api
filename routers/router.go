// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"web_bee/car_qr_api/controllers"
	"github.com/astaxie/beego"
	_ "web_bee/car_qr_api/extra_apps/gernate_qr/routers"
	_ "web_bee/car_qr_api/extra_apps/backend/routers"
	_ "web_bee/car_qr_api/extra_apps/frontend/routers"
	_ "web_bee/car_qr_api/extra_apps/ad_customers/routers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/ad_list",
			beego.NSInclude(
				&controllers.AdListController{},
			),
		),

		beego.NSNamespace("/ad_log",
			beego.NSInclude(
				&controllers.AdLogController{},
			),
		),

		beego.NSNamespace("/ad_users",
			beego.NSInclude(
				&controllers.AdUsersController{},
			),
		),

		beego.NSNamespace("/call_log",
			beego.NSInclude(
				&controllers.CallLogController{},
			),
		),

		beego.NSNamespace("/qr_info",
			beego.NSInclude(
				&controllers.QrInfoController{},
			),
		),

		beego.NSNamespace("/user_info",
			beego.NSInclude(
				&controllers.UserInfoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
