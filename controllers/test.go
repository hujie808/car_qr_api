package controllers

import (
	"github.com/astaxie/beego"
	beeLogger "github.com/beego/bee/logger"
)

type TestController struct {
	beego.Controller

}

// URLMapping ...
func (c *TestController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Get", c.Get)
}

// @router / [get]
func (c *TestController) GetAll() {
	username := c.GetString("usernmaes")
	g_username := c.Input().Get("username")
	beeLogger.Log.Info(username + g_username)
	beeLogger.Log.Info(c.Ctx.Input.Domain())
	c.Data["json"] = 11111
	flash := beego.NewFlash()
	flash.Data["notice"]="www.baidu.com"
	flash.Notice("Settings saved!")
	c.ServeJSON()
}

// @router /flash [get]
func (c *TestController) Get() {
	a:=new(AdUsersController)
	beeLogger.Log.Info(a.URLFor("AdUsersController.GetAll","name","list","password","listpass"))
	c.Data["json"] ="1111"

	c.ServeJSON()
}
