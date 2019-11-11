package qr_utils

import (
	beeLogger "github.com/beego/bee/logger"
	uuid "github.com/satori/go.uuid"
	"time"
	"web_bee/car_qr_api/models"
	"web_bee/car_qr_api/utils/bee_orm"
)


//新建 二维码info
func QrAddRandom(number, adId int, ) []int{
	//number 生成二维码个数
	//adId 生成的广告商id
	var qr_info_list []int
	for i := 1; i <= number; i++ {
		u1 := uuid.Must(uuid.NewV4())
		var qrInfo = &models.QrInfo{
			Uuid:       u1.String(),
			Inuse:      0,
			UserId:     nil,
			StartDate:  time.Time{},//启动时间
			VisitDate:  time.Time{},
			VisitTimes: 0,
		}
		if adId>0{
			qrInfo.AdId =&models.AdList{Id:adId}
		}
		qr_info_id,err :=bee_orm.O.Insert(qrInfo)
		if err !=nil{
			beeLogger.Log.Errorf("qr_info_add.go QrAddRandom Insert qr_info_id err=",err)
		}
		qr_info_list=append(qr_info_list, int(qr_info_id))
	}
	return qr_info_list
}
