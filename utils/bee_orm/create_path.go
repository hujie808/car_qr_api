package bee_orm

import (
	"github.com/astaxie/beego"
	"log"
	"os"
)

func init() {
	ad_static := beego.AppConfig.String("img::ad_static")
	qr_image_path := beego.AppConfig.String("img::qr_image_path")
	qr_reuslt_path:= beego.AppConfig.String("img::qr_reuslt_path")
	qr_model_path := beego.AppConfig.String("img::qr_model_path")
	pathList := []string{ad_static, qr_image_path, qr_reuslt_path, qr_model_path}
	for _, path := range pathList {
		_, err := os.Stat(path)
		if err != nil {
			if os.IsExist(err) {
				continue
			}
			if os.IsNotExist(err) {
				log.Println("文件目录不存在 正在创建:",path)
				os.MkdirAll(path, os.ModePerm)
				continue
			}
			os.MkdirAll(path, os.ModePerm)
			continue
		}
	}
}
