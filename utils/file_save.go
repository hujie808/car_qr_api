package utils

import "net/http"

//import (
//	"fmt"
//	"github.com/astaxie/beego"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//	"path/filepath"
//	"strconv"
//	"strings"
//)
//
//
func RenderError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-type", "application/text")
	w.WriteHeader(code)
	w.Write([]byte(msg))

}
//
//func FileSave(Ctx *beego.Controller, fileKey, path string, maxSize int) string {
//	r := Ctx.Ctx.Request
//
//	w := Ctx.Ctx.ResponseWriter
//	maxUploadSize := int64(maxSize * 1024 * 1024)
//	uploadPath := bootstrap.StaticAssets + path
//	//验证大小
//	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
//	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
//		RenderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
//		return ""
//	}
//	//读取文件
//	file, o, err := r.FormFile(fileKey)
//	if err != nil {
//		log.Println("file_save.go FileSave FormFile err=", err)
//		return ""
//	}
//	if file == nil {
//		return ""
//	}
//
//	if o.Size > maxUploadSize {
//		RenderError(w, "INVALID_FILE", http.StatusBadRequest)
//		return ""
//	}
//	if err != nil {
//		RenderError(w, "INVALID_FILE", http.StatusBadRequest)
//		return ""
//	}
//	defer file.Close()
//	fileBytes, err := ioutil.ReadAll(file)
//	if err != nil {
//		RenderError(w, "INVALID_FILE", http.StatusBadRequest)
//		return ""
//	}
//	//判断文件类型
//	filetype := http.DetectContentType(fileBytes)
//	if filetype != "image/jpeg" && filetype != "image/jpg" &&
//		filetype != "image/gif" && filetype != "image/png" &&
//		filetype != "application/pdf" {
//		RenderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
//		return ""
//	}
//	//给文件命名
//	fileName := strconv.Itoa(comm.NowUnix()) + strconv.Itoa(comm.Random(100))
//	filName := strings.Split(o.Filename, ".")
//	fileType := "." + filName[len(filName)-1]
//	if err != nil {
//		RenderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
//		return ""
//	}
//	newPath := filepath.Join(uploadPath, fileName+fileType)
//	fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)
//	//返回消息
//	newFile, err := os.Create(newPath)
//	if err != nil {
//		RenderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
//		return ""
//	}
//	defer newFile.Close()
//	if _, err := newFile.Write(fileBytes); err != nil {
//		RenderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
//		return ""
//	}
//	//w.Write([]byte("SUCCESS"))
//	return "/" + newPath
//}