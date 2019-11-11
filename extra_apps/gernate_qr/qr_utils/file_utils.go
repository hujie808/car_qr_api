package qr_utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func createDateDir(Path string) string {
	//判断目录是否存在
	err := os.MkdirAll(Path, os.ModePerm)
	if err != nil {
		log.Println("指定结果目录设置错误 err=", err)
		os.Exit(500)
	} else {
		log.Println("已成功创建目录:", Path)
	}
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(Path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	fmt.Println("请等待")
	return "./"+folderPath
}