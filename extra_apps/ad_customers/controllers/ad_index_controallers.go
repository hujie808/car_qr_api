package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"web_bee/car_qr_api/extra_apps/ad_customers/ad_valid"
	"web_bee/car_qr_api/models"
	"web_bee/car_qr_api/utils"
	"web_bee/car_qr_api/utils/bee_orm"
)

type Index struct {
	beego.Controller
	user *models.AdUsers
	r    utils.Result
}

func (c *Index) URLMapping() {
	c.Mapping("GetIndex", c.GetIndex)
	c.Mapping("GetListAdd", c.GetListAdd)
	c.Mapping("ImgEdit", c.ImgEdit)
	c.Mapping("AdDelete", c.Delete)
	c.Mapping("AdPut", c.Put)
}

func (c *Index) Prepare() {
	if c.Ctx.Request.Method == "GET" {
		return
	}

	var user_id int
	tmp := c.GetSession("user_id")
	if tmp != nil {
		user_id = tmp.(int)
	}
	var user models.AdUsers
	if user_id > 0 {
		user.Id = user_id
		user_err := bee_orm.O.Read(&user)
		if user_err != nil {
			c.Data["Title"] = "您的用户不正确"
			c.Data["Msg"] = "请重新登录"
			c.TplName = "faild.html"
			c.StopRun()
		}
	}
	c.user = &user
	//验证用户是否过期
	//postIndex := beego.URLFor("AdIndex.PostIndex")
	//re, _ := regexp.Compile(fmt.Sprintf("^%s.?$", postIndex))
	//conBool := re.MatchString(c.Ctx.Request.RequestURI)
	//if !conBool {

	if user.ExpireDate.Unix() <= time.Now().Unix() {
		if strings.Contains(c.Ctx.Request.URL.RequestURI(), "index") {
			return
		}
		c.r.Msg = "您的账号已过期 禁止操作"
		c.Data["json"]=c.r
		c.TplName = "faild.html"
		c.Ctx.Output.Status = 400
		c.ServeJSON()
		c.StopRun()
	}
	//}
}

//用户登录
// @router /index [get]
func (c *Index) GetIndex() {
	c.Layout = "ad_index/base.html"
	c.TplName = "ad_index/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "ad_index/index_scripts.html"
	c.LayoutSections["Body"] = "ad_index/index_body.html"
}

//首页信息
// @router /index [post]
func (c *Index) PostIndex() {
	data := make(map[string]interface{})
	user := c.user
	qr_qs := bee_orm.O.QueryTable(new(models.QrInfo))
	ad_qs := bee_orm.O.QueryTable(new(models.AdList))
	user_all_qr := qr_qs.Filter("ad_users_id__id", user.Id)
	user_qr_all_count, _ := user_all_qr.Count()
	data["qr_all_count"] = user_qr_all_count //用户全部二维码
	qr_inuse_count, _ := user_all_qr.Filter("inuse", 1).Count()
	data["qr_inuse_count"] = qr_inuse_count //用户正在使用二维码
	all_ad := ad_qs.Filter("ad_user_id__id", user.Id)
	ad_all_count, _ := all_ad.Count()
	data["ad_all_count"] = ad_all_count //用户的广告个数
	var mo []*models.AdList
	num, err := all_ad.All(&mo)
	if err!=nil{

	}
	//for _,m :=range mo{
	//	m.Pic=beego.AppConfig.String("domainName")+m.Pic
	//	log.Println(m.Pic)
	//}
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	data["ad_all"] = mo                          //用户所有的广告
	data["ad_count"] = user.AdCount              //用户可设置的最大数量
	var defAdList models.AdList
	bee_orm.O.QueryTable("ad_list").Filter("ad_user_id__id", user.Id).Filter("defualt_flag",1).One(&defAdList)
	data["default_ad_list"] =defAdList.Id  //默认广告id
	data["expire_date"] = user.ExpireDate        //账号过期时间
	c.r.Data = data
	//c.r.Code=0
	c.Data["json"] = c.r
	c.ServeJSON()
}

//广告表单提交
// @router /ad_list [get]
func (c *Index) GetListAdd() {
	c.Layout = "ad_index/base.html"
	c.TplName = "ad_index/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "ad_index/list_add_scripts.html"
	c.LayoutSections["Body"] = "ad_index/list_add_body.html"
}

//广告表单提交 post
// @router /ad_list [post]
func (c *Index) PutListAdd() {
	log.Println(string(c.Ctx.Input.RequestBody))
	c.adListValid()
	o := bee_orm.O
	//过滤无效用户id
	//过滤最大广告数量
	ad_list_all := o.QueryTable("ad_list").Filter("ad_user_id__id", c.user.Id)
	ad_list_count, _ := ad_list_all.Count()
	if int(ad_list_count) >= c.user.AdCount {
		c.r.Msg = "您的广告模板不能再多了,亲~"
		c.Data["json"] = c.r
		c.Ctx.Output.Status = 400
		c.ServeJSON()
		c.StopRun()
	}
	var ad_list models.AdList
	if err := c.ParseForm(&ad_list); err == nil {
		//ad_list.Pic = path
		ad_list.AddDate = time.Now()
		ad_list.UpdateDate = time.Now()
		ad_list.AdUserId = c.user

		if ad_list_count == 0 {
			ad_list.DefualtFlag = 1
		} else {
			if ad_list.DefualtFlag == 1 {
				_, err := bee_orm.O.QueryTable("ad_list").Filter("ad_user_id__id", c.user.Id).Update(orm.Params{"defualt_flag": 0,})
				if err != nil {
					log.Println("index.go Put() Update err=", err)
				}
			}
		}

		if _, err := models.AddAdList(&ad_list); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = ad_list
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

//用户删除广告
// Delete ...
// @Title Delete
// @Description delete the AdList
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /ad_list/:id [delete]
func (c *Index) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var d_AdList models.AdList
	d_AdList.Id = id
	d_err := bee_orm.O.Read(&d_AdList, "id")
	if d_err != nil {
		log.Println("index.go delete() Read err=", d_err)
	}
	if err := models.DeleteAdList(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}

	if d_AdList.DefualtFlag == 1 {
		var one_ad models.AdList
		err := bee_orm.O.QueryTable("ad_list").Filter("ad_user_id__id", c.user.Id).One(&one_ad)
		log.Println(one_ad)
		if err != nil {
			log.Println("index.go Delete() One err=", err)
		}
		if one_ad.Id > 0 {
			one_ad.DefualtFlag = 1
			_, err = bee_orm.O.Update(&one_ad, "DefualtFlag")
			if err != nil {
				log.Println("index.go Delete() Update err=", err)
			}
		}
	}
	c.ServeJSON()
}

//广告表单修改
// @router /ad_list/:id [get]
func (c *Index) GetListUpdate() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var d_AdList models.AdList
	d_AdList.Id = id
	o := bee_orm.O
	err := o.Read(&d_AdList)
	if err != nil {
		log.Println("ad_index_controallers.go GetListUpdate Read err=", err)
	}
	c.Data["ad_list"] = d_AdList
	c.Layout = "ad_index/base.html"
	c.TplName = "ad_index/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "ad_index/list_put_scripts.html"
	c.LayoutSections["Body"] = "ad_index/list_add_body.html"
}

//用户更改ad_LIst
// Put ...
// @Title Put
// @Description update the AdUsers
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.AdUsers	true		"body for AdUsers content"
// @Success 200 {object} models.AdUsers
// @Failure 403 :id is not int
// @router /ad_list/:id [put]
func (c *Index) Put() {
	c.adListValid()
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.AdList{Id: id}
	err := bee_orm.O.Read(&v)
	if err != nil && v.AdUserId == nil {
		c.r.Msg = "无此用户"
		c.Data["json"] = c.r
		c.ServeJSON()
		c.Ctx.Output.Status = 400
		return
	}
	if err := c.ParseForm(&v); err == nil { //TODO 验证
		path, code := c.Upload("myfile")
		if code == 200 {
			v.Pic = path
		} else if code == 400 {
			c.r.Msg = path
			c.Data["json"] = c.r
			c.ServeJSON()
			c.Ctx.Output.Status = 400
			return
		}
		v.UpdateDate = time.Now()
		if v.DefualtFlag == 1 {
			_, err := bee_orm.O.QueryTable("ad_list").Filter("ad_user_id__id", c.user.Id).Update(orm.Params{"defualt_flag": 0,})
			if err != nil {
				log.Println("index.go Put() Update err=", err)
			}
		}
		v.AdUserId = c.user
		//v.DefualtFlag=1
		if err := models.UpdateAdListById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

//富文本上传图片接口
// @router /ad_img [post]
func (c *Index) ImgEdit() {
	path, code := c.Upload("file")
	if code == 200 {
		c.r.Code = 0
		paths := path[1:]
		rDict := make(map[string]string)
		rDict["src"] = paths
		log.Println(paths)
		log.Println(rDict["src"])
		pathList := strings.Split(paths, "/")
		title := pathList[len(pathList)-1]
		rDict["title"] = title
		c.r.Data = rDict
		c.Data["json"] = c.r
		i := 0
		for i == 1 {
			_, err := os.Stat(path)
			if os.IsExist(err) {
				i = 1
			}
		}
	} else {
		c.Data["json"] = c.r
		c.Ctx.ResponseWriter.Status = 400
	}
	c.ServeJSON()
}

//参数效验
func (c *Index) adListValid() {
	//效验参数
	var ad_v ad_valid.AdListValid
	valid := validation.Validation{}
	c.ParseForm(&ad_v)
	if err := c.ParseForm(&ad_v); err != nil {
		log.Println("ad_index_controallers.go adListValid err=", err)
	}
	b, err := valid.Valid(&ad_v)
	if err != nil {
		log.Println("index.go valid.Valid err =", err)
	}
	ad_v.AdUserId = &models.AdUsers{Id: c.user.Id}
	if ad_v.AdUserId.Id < 0 {
		c.r.Code = 2000
		c.r.Msg = "您登录的用户与提交用户不符"
		c.Ctx.Output.Status = 400
		c.Data["json"] = c.r
		c.ServeJSON()
		c.StopRun()
	}

	if !b { //不通过则返回
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			c.r.Code = 2000
			c.r.Msg = err.Key + ":" + err.Message
			c.Ctx.Output.Status = 400
			c.Data["json"] = c.r
			c.ServeJSON()
			c.StopRun()
		}
	}
}

//上传文件函数
func (this *Index) Upload(myfile string) (r_str string, code int) {
	w := this.Ctx.ResponseWriter
	r := this.Ctx.Request
	//判断文件
	size, _ := beego.AppConfig.Int("img::img_size")
	f, h, err := this.GetFile(myfile) //获取上传的文件
	if f == nil {
		code = 201
		r_str = "您未上传图片"
		return
	}

	var filename string
	maxUploadSize := int64(size * 1024 * 1024)
	//验证大小
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		r_str = "文件太大超过限制 亲~"
		return
	}

	filenameList := strings.Split(h.Filename, ".") //apche 5D%漏洞

	res, _ := regexp.Compile("jpg|png|gif|jpeg|JPG|PNG")
	filebool := res.MatchString(filenameList[len(filenameList)-1]) //TODO .文件判断
	if !filebool {
		r_str = "文件格式类型暂不支持"
		return
	} else {
		filename = fmt.Sprintf("%d_%d.%s", this.user.Id, time.Now().Unix(), filenameList[len(filenameList)-1])
	}

	code = 400
	r_str = ""
	path := beego.AppConfig.String("img::ad_static") + filename //文件目录
	defer f.Close()                                             //关闭上传的文件，不然的话会出现临时文件不能清除的情况

	if err != nil {
		r_str = ""
		return
	}
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		r_str = "文件读取错误"
		return
	}

	//验证文件类型
	filetype := http.DetectContentType(fileBytes)
	if filetype != "image/jpeg" && filetype != "image/jpg" &&
		filetype != "image/gif" && filetype != "image/png" &&
		filetype != "application/pdf" {
		r_str = "文件类型错误"
		return
	}
	this.SaveToFile(myfile, path) //存文件
	r_str = path
	code = 200
	return
}
