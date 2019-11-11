package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/gernate_qr/controllers:Istest"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/gernate_qr/controllers:Istest"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/gernate_qr/controllers:QrControllers"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/gernate_qr/controllers:QrControllers"],
        beego.ControllerComments{
            Method: "CreateQr",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
