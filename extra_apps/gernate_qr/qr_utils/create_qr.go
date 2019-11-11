package qr_utils

import (
	"fmt"
	"github.com/astaxie/beego"
	beeLogger "github.com/beego/bee/logger"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/nfnt/resize"
	uuid "github.com/satori/go.uuid"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"time"
	"web_bee/car_qr_api/models"
	"web_bee/car_qr_api/utils/bee_orm"
)

type QrCrate struct {
	Uuid         string  //生成uuid
	Url          string  //链接
	ImgX         int     //模板图片像素X
	ImgY         int     //模板图片像素Y
	QrBoundsX    int     //二维码边界X
	QrBoundsY    int     //二维码边界Y
	ImgModelPath string  //模板图片地址
	QrMin        float64 //二维码图片缩放比例
	QrWideX      int     //生成宽多少张
	QrHighY      int     //生成长多少张
	BankX        int     //留白X
	BankY        int     //留白Y
	AdUserId     int     //商户id 可为空
}

//一个带url的图片生成
func (q *QrCrate) QrOne(path string) string {
	r_path := path + q.Uuid + ".png"
	rUrl := fmt.Sprintf("%s?uuid=%s", q.Uuid, q.Uuid)
	qrCode, _ := qr.Encode(rUrl, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 600, 600)
	file, _ := os.Create(r_path)
	defer file.Close()
	png.Encode(file, qrCode)
	return r_path
}

func (q *QrCrate) fixSize(img1X, img1Y int) (new2X, new2Y int) {
	var ( //为了方便计算，将两个图片的宽转为 float64
		img1XWidth, img1YWidth = float64(img1X), float64(img1Y)
		ratio1, ratio2         float64
	)
	ratio1 = img1XWidth / q.QrMin
	ratio2 = img1YWidth / q.QrMin
	return int(ratio1), int(ratio2)
}

func (q *QrCrate) ModelsQr(qr_path, temPath string) string {
	file1, _ := os.Open(q.ImgModelPath) //打开图片1
	file2, _ := os.Open(qr_path)        //打开图片2
	defer file1.Close()
	defer file2.Close()
	var (
		img1, img2 image.Image
		err        error
	)
	if img1, _, err = image.Decode(file1); err != nil {
		log.Fatal(err)
		return ""
	}
	if img2, _, err = image.Decode(file2); err != nil {
		log.Fatal(err)
		return ""
	}

	b2 := img2.Bounds()
	img2X, img2Y := q.fixSize(b2.Max.X, b2.Max.Y)
	newImg := image.NewNRGBA(image.Rect(0, 0, q.ImgX, q.ImgY))
	// 调用resize库进行图片缩放(高度填0，resize.Resize函数中会自动计算缩放图片的宽高比)
	m1 := resize.Resize(uint(q.ImgX), uint(q.ImgY), img1, resize.Lanczos3)
	m2 := resize.Resize(uint(img2X), uint(img2Y), img2, resize.Lanczos3)
	draw.Draw(newImg, m1.Bounds(), m1, m1.Bounds().Min, draw.Src)
	r := m2.Bounds().Sub(m2.Bounds().Min).Add(image.Point{q.QrBoundsX, q.QrBoundsY})
	draw.Draw(newImg, r, m2, m2.Bounds().Min, draw.Over)
	// 保存文件
	r_path := temPath + "models_" + q.Uuid + ".png"
	imgfile, _ := os.Create(r_path)
	defer imgfile.Close()
	jpeg.Encode(imgfile, newImg, &jpeg.Options{100})

	return r_path
}

//3.生成可设定固定长 宽 模板

//1.拿到宽的个数与长的个数
//2.制造出来像素 X*宽,Y*高 的图片

func (q *QrCrate) QrResult(path_list []qrDict, resultX, resultY int, result_path string) string {
	//创建图片 背景白色
	newImg := image.NewNRGBA(image.Rect(0, 0, resultX, resultY))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	//newImg.Set(resultX, resultX, white)
	//拼接图片
	for _, path := range path_list {
		file, _ := os.Open(path.Path)
		img, _, _ := image.Decode(file)
		draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min.Sub(image.Pt(path.X, path.Y)), draw.Over)
		file.Close()
	}
	r_path := result_path + "models_" + time.Now().String()[0:16] + ".png"
	imgfile, _ := os.Create(r_path)
	defer imgfile.Close()
	jpeg.Encode(imgfile, newImg, &jpeg.Options{100})
	return r_path

}

type qrDict struct {
	X    int    //图片x位置
	Y    int    //图片y位置
	Path string //图地址
}

//从数据库新增 空二维码 并返回uuid
func (q *QrCrate) QrAddRandom() string {
	//number 生成二维码个数
	//adId 生成的广告商id

	uuidStr := uuid.Must(uuid.NewV4()).String()
	u1 := strings.Replace(uuidStr, "-", "", 4)
	var qrInfo = &models.QrInfo{
		Uuid:       u1,
		Inuse:      0,
		UserId:     nil,
		StartDate:  time.Time{}, //启动时间
		VisitDate:  time.Time{},
		VisitTimes: 0,
	}
	//商户Id可以为空
	if q.AdUserId > 0 {
		ad := new(models.AdUsers)
		//ad.Id=q.AdUserId
		ad.Id = q.AdUserId
		qrInfo.AdUsersId = ad
	}
	_, err := bee_orm.O.Insert(qrInfo)
	if err != nil {
		beeLogger.Log.Errorf("qr_info_add.go QrAddRandom Insert qr_info_id err=", err)
	}
	return u1
}

func (q *QrCrate) Start() {
	qr_reuslt_path := beego.AppConfig.String("img::qr_image_path")
	tmpDir := createDateDir(qr_reuslt_path) + "/" //二维码临时文件夹
	r_reuslt_path := beego.AppConfig.String("img::qr_reuslt_path")
	result_path := createDateDir(r_reuslt_path) + "/" //结果文件夹
	var listQrDict []qrDict
	cursorX := 0//
	cursorY := 0
	for i := 0; i < q.QrWideX; i++ { //计算出最后生成尺寸 并
		if i < q.QrWideX&&i>0 { //横向计算
			cursorX = cursorX + q.ImgX + q.BankX
		} else if i > 0 {
			cursorX = cursorX + q.ImgX
		}
		for j := 0; j < q.QrHighY; j++ {
			q.Uuid = q.QrAddRandom() //uuid生成 并插入数据库
			path := q.QrOne(tmpDir) //二维码生成方法
			rQr := q.ModelsQr(path, tmpDir) //合成模板方法
			r_Dic := qrDict{
				Path: rQr,
			}
			//右边与下面是否有图片,有图片加留白
			if j < q.QrHighY&&j>0 {
				cursorY = cursorY + q.ImgY + q.BankY
			} else if j>0{
				cursorY = cursorY + q.ImgY
			}else if j==0{
				cursorY=0
			}
			//判断左下 没有图片则不留白
			r_Dic.X = cursorX
			r_Dic.Y = cursorY
			listQrDict = append(listQrDict, r_Dic)
		}
	}

	q.QrResult(listQrDict, cursorX+q.ImgX, cursorY+q.ImgY, result_path)
}
