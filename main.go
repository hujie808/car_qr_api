package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "web_bee/car_qr_api/extra_apps/backend/backend_admin"
	_ "web_bee/car_qr_api/routers"
	//_ "web_bee/car_qr_api/utils/bee_cors" //API 增加 CORS 支持
	_ "web_bee/car_qr_api/utils/bee_orm"
	_ "web_bee/car_qr_api/utils/bee_redis" //API redis
	_ "web_bee/car_qr_api/utils/bee_session"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	//
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}

	beego.Run()
}
