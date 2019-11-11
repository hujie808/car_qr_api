package ad_valid

import (
	"github.com/astaxie/beego/validation"
	"time"
	"web_bee/car_qr_api/models"
)

type AdListValid struct {
	Id          int		`null;valid:"Min(1)"`
	AdUserId    *models.AdUsers
	AddDate     time.Time
	UpdateDate  time.Time
	Pic         string	  `null;valid:"Match(/.*static.*/)"`
	Url         string    `null;valid:"MinSize(7);MaxSize(255)"`
	Content     string    `valid:"MinSize(1);MaxSize(9999)"`
	Template    string
	DefualtFlag int       `valid:"Max(1)"`
}

func (u *AdListValid) Valid(v *validation.Validation) {
	//log.Println(u.AdUserId)
    //if strings.Index(u.Pic, "admin") != -1 {
    //    // 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
    //    v.SetError("Name", "名称里不能含有 admin")
    //}
}
