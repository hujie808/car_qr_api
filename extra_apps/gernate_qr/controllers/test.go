package controllers

import "github.com/astaxie/beego"

type Istest struct {
	beego.Controller
}

func (c *Istest) URLMapping() {
	c.Mapping("Get",c.Get)
}

// @router / [get]
func(c *Istest)Get(){
	c.Data["json"]="ok start"
	c.ServeJSON()
}