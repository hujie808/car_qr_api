package main

import (
	"github.com/astaxie/beego"
	"testing"
	"web_bee/car_qr_api/extra_apps/gernate_qr/qr_utils"
)

func TestM(t *testing.T) {
	url := beego.AppConfig.String("img::qr_url")
	var qrTest = qr_utils.QrCrate{
		Uuid:         "weebksdjfsadj",
		Url:          url,
		ImgX:         800,
		ImgY:         800,
		QrBoundsX:    200,
		QrBoundsY:    150,
		ImgModelPath: "/Users/a/go/src/web_bee/car_qr_api/static/img/qr_model/qrmodels.jpeg",
		QrMin:        1.5,
		QrWideX:      10,
		QrHighY:      10,
		BankX:        30,
		BankY:        30,
		AdUserId:         1,
	}
	qrTest.Start()
}
