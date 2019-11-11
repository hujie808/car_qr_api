package backend_amin

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/beego"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	_ "github.com/GoAdminGroup/themes/adminlte" // 必须引入，不然报错
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"web_bee/car_qr_api/extra_apps/backend/car_admin"
)

var AdminPath  string ="static/admin"

func init() {

	app := beego.BeeApp

	eng := engine.Default()
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       "39.96.160.74",
				Port:       "3306",
				User:       "ccl123bat23E",
				Pwd:        "cclbat32123EQ*",
				Name:       "car_card",
				MaxIdleCon: 30,
				MaxOpenCon: 150,
				Driver:     config.DriverMysql,
			},
		},
		UrlPrefix: "admin",
		Debug:     true,
		Store: config.Store{
			Path:   "./"+AdminPath+"/",
			Prefix: AdminPath,
		},
		Language: language.CN,
	}
	adminPlugin := admin.NewAdmin(car_admin.Generators)

	adminPlugin.AddGenerator("ad_lists", car_admin.GetQrInfoTable)

	template.AddComp(chartjs.NewChart())
	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(app); err != nil {
		panic(err)
	}
	app.Handlers.Get("/admin", func(ctx *context.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return datamodel.GetContent()
		})
	})

	//beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	//beego.BConfig.Listen.HTTPPort = 8081
	//app.Run()
}
