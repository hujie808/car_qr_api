package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"web_bee/car_qr_api/models"
	"web_bee/car_qr_api/utils"
	"web_bee/car_qr_api/utils/bee_session"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) URLMapping() {
	c.Mapping("GetLogin", c.GetLogin)
	c.Mapping("PostLogin", c.PostLogin)
}

//登录请求页面
//@param username adpass
// @router /login [get]
func (c *LoginController) GetLogin() {
	c.TplName = "ad_index/login.html"
}

//用户登录post提交页面
//@param username adpass
// @router /login [post]
func (c *LoginController) PostLogin() {
	var v models.AdUsers
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v);
	if err != nil {
		c.Data["json"] = err.Error()
	}

	o := orm.NewOrm()
	qs := o.QueryTable(&models.AdUsers{})
	adUser := qs.Filter("username", v.Username).Filter("adpass", v.Adpass)
	var r utils.Result
	if adUser.Exist() {
		var user models.AdUsers
		adUser.One(&user)
		r.Data = user.Id
		r.Code = 1000
		log.Printf("%T,%s", r.Json(), r.Json())

		c.SetSession("user_id", user.Id)

		s, _ := beego.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		s.Set("user_id", user.Id)
		ss, _ := bee_session.BeeSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		ss.Set("user_id", user.Id)
		defer ss.SessionRelease(c.Ctx.ResponseWriter)
		c.Data["json"] = r
	} else {
		r.Code = 2000
		r.Msg = "用户名或者密码错误"
		c.Data["json"] = r
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}


