package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/frontend/controllers:Istest"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/frontend/controllers:Istest"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/frontend/controllers:QrController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/frontend/controllers:QrController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/index/:uuid/:ad/:ad_customer`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/frontend/controllers:QrController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/frontend/controllers:QrController"],
        beego.ControllerComments{
            Method: "Reg",
            Router: `/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
