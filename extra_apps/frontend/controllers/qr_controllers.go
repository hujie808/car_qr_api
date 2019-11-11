package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"web_bee/car_qr_api/models"
	"strconv"
	"web_bee/car_qr_api/utils/bee_redis"
	"encoding/json"
	"time"
	"web_bee/car_qr_api/utils/bee_orm"
	"fmt"
)
 type QrController struct {
 	beego.Controller
 }

 type AdListRes struct {
 	Url string `json:"url"`
 	Pic string `json:"pic"`
 	Content string `json:"content"`
 }
func (c *QrController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Reg", c.Reg)
}

/*
	method get
    use for user qrcode login
    Params
    uuid: 		 qr's uuid for unique
    ad:			 the ad contents id
    ad_customer: the ad customer id
	Edit by Hyw 2019.10.22
*/
// @router /index/:uuid/:ad/:ad_customer [get]
func (this *QrController) Index()  {
	uuid:=this.Ctx.Input.Param(":uuid")// 二维码id
	ad:=this.Ctx.Input.Param(":ad")//广告id  预留字段
	ad_id,_ := strconv.Atoi(ad)//广告id string 2 intger
	ad_customer:=this.Ctx.Input.Param(":ad_customer")//绑定的广告商id
	ad_customer_id,_ := strconv.Atoi(ad_customer)//广告商id  string 2 intger

	//判断是否有该uuid
	qr := models.QrInfo{Uuid:uuid}
	err:= bee_orm.O.Read(&qr,"Uuid")
	if err == orm.ErrNoRows{
		this.Abort("401")//没有该二维码
	}
	var adurl string  //广告跳转链接  可以为空
	var adpic string  //广告弹窗图片
	var adcontent string //广告文字内容 可以为空
	var adRes string  //广告数据   json字符串
	ad_id_flag := 0

	//判断该二维码是否绑定了广告客户
	if ad_customer_id >0{
		adCustomer := models.AdUsers{Id:ad_customer_id}
		err:=bee_orm.O.Read(&adCustomer)
		if err == orm.ErrNoRows{
			//不存在该广告用户
		}else{
			//存在该广告用户  1.判断广告用户是否过期   2.如果广告用户没有过期  判断广告用户广告id是否存在 如果存在 使用该id进行展示
			currentTime:=time.Now()
			if adCustomer.ExpireDate.Before(currentTime){
				//用户已过期
			}else{
				var adList models.AdList
				err := bee_orm.O.QueryTable("ad_list").Filter("ad_user_id",adCustomer.Id).Filter("default_flag",1).One(&adList,"Url","Pic","Content")
				if (err == orm.ErrMultiRows)||(err == orm.ErrNoRows) {
					// 数据有误、有多条记录或者没有记录
					fmt.Printf("Returned Multi Rows Not One")
				}else{
					//有该广告数据  就不需要讨论二维码绑定id数据了
					ad_id_flag = 1
					adurl = adList.Url
					adpic = adList.Pic
					adcontent = adList.Content
				}
			}
		}
	}

	//判断该二维码是否已设置了广告id
	if (ad_id_flag==0)&&(ad_id>0){
		adContent := models.AdList{Id:ad_id}//  查找数据库中是否有该广告数据
		err:= bee_orm.O.Read(&adContent)
		if err == orm.ErrNoRows{
			//没有广告内容  选择系统设置的广告信息进行展示
			ad_id_flag = 2
			adRes :=this.GetAd()
		}else{
			//
			adurl = adContent.Url
			adpic = adContent.Pic
			adcontent = adContent.Content
		}
	}else if ad_id_flag==0 {
		//判断该广告是否存在，如果到期选择系统广告进行展示
		ad_id_flag = 2
		adRes :=this.GetAd()
	}
	if ad_id_flag == 2{
		var ad_res AdListRes
		//格式化adRes
		json.Unmarshal([]byte(adRes), &ad_res)
		adurl = ad_res.Url
		adpic = ad_res.Pic
		adcontent = ad_res.Content
	}
	return
}

// @router /reg [post]
func (this *QrController) Reg(){

}


func (this *QrController) GetAd() string{
	//判断 默认广告是否已经写入redis
	if bee_redis.Cache.IsExist("defaultAd"){
		return string(bee_redis.Cache.Get("defaultAd").([]byte))
	}else{
		var adRes AdListRes
		qb, _ := orm.NewQueryBuilder("mysql")

		// 构建查询对象
		qb.Select("url",
			"pic","content").
			From("ad_list").
			Where("default_flag > ?").
			OrderBy("id").Desc().
			Limit(1).Offset(0)

		// 导出 SQL 语句
		sql := qb.String()

		// 执行 SQL 语句

		bee_orm.O.Raw(sql, 1).QueryRows(&adRes)
		result,_ := json.Marshal(adRes)
		bee_redis.Cache.Put("defaultAd",adRes,60*60*24*time.Second)
		return string(result)
	}
}