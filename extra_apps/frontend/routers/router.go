package routers

import (
	"github.com/astaxie/beego"
)

func init(){

	ns := beego.NewNamespace("/frontend",


			//beego.NSInclude(
			//	&controllers.QrController{},
			//),

	)
	beego.AddNamespace(ns)
}