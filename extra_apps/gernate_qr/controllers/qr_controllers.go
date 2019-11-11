package controllers

import (
	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
	"web_bee/car_qr_api/extra_apps/gernate_qr/qr_utils"
	"web_bee/car_qr_api/utils"
)

type QrControllers struct {
	beego.Controller
	r utils.Result
}

func (c *QrControllers) Prepare() {
	secret_key := beego.AppConfig.String("secret_key")
	get_secret_key := gjson.Get(string(c.Ctx.Input.RequestBody), "secret_key").String()
	if secret_key != get_secret_key {
		c.r.Msg = "key错误,请勿非法访问"
		c.Data["json"] = c.r
		c.Ctx.Output.Status = 401
		c.ServeJSON()
		c.StopRun()
	}
}

//用户
// @router / [post]
func (c *QrControllers) CreateQr() {
	gjson.Get(string(c.Ctx.Input.RequestBody), "secret_key")
	ad_user_id := int(gjson.Get(string(c.Ctx.Input.RequestBody), "ad_user_id").Int())
	ad_list_count := int(gjson.Get(string(c.Ctx.Input.RequestBody), "ad_list_count").Int())
	url := beego.AppConfig.String("img::qr_url")
	var qrStruct = qr_utils.QrCrate{
		Uuid:         "weebksdjfsadj",
		Url:          url,
		ImgX:         800,
		ImgY:         800,
		QrBoundsX:    200,
		QrBoundsY:    150,
		ImgModelPath: "/Users/a/go/src/web_bee/car_qr_api/static/img/qr_model/qrmodels.jpeg",
		QrMin:        1.5,
		QrWideX:      3,
		QrHighY:      3,
		BankX:        30,
		BankY:        30,
	}
	if ad_user_id>0{
		qrStruct.AdUserId=ad_user_id
	}
	qrOneCount := qrStruct.QrWideX * qrStruct.QrHighY
	qr_for_Count := (ad_list_count / qrOneCount) + (ad_list_count % qrOneCount)
	for i := 0; i < qr_for_Count; i++ {
		qrStruct.Start()
	}

}
