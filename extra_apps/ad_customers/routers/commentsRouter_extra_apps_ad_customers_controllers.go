package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"],
        beego.ControllerComments{
            Method: "ImgEdit",
            Router: `/ad_img`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"],
        beego.ControllerComments{
            Method: "GetListAdd",
            Router: `/ad_list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"],
        beego.ControllerComments{
            Method: "PutListAdd",
            Router: `/ad_list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/ad_list/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/ad_list/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"],
        beego.ControllerComments{
            Method: "GetListUpdate",
            Router: `/ad_list/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"],
        beego.ControllerComments{
            Method: "GetIndex",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:Index"],
        beego.ControllerComments{
            Method: "PostIndex",
            Router: `/index`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:LoginController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:LoginController"],
        beego.ControllerComments{
            Method: "GetLogin",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:LoginController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/extra_apps/ad_customers/controllers:LoginController"],
        beego.ControllerComments{
            Method: "PostLogin",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
