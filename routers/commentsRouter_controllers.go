package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdListController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdLogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:AdUsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:CallLogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:QrInfoController"],
        beego.ControllerComments{
            Method: "GetQr",
            Router: `/add/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:TestController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:TestController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:TestController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:TestController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/flash`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["web_bee/car_qr_api/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
