package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego/cache"
	"time"
)

//发送短信验证方法
func TelSendMsg(b_cache cache.Cache, tel string) {
	code := GetRandomString(6)
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4FbhzAUvjE73Xh5VfTsJ", "omZedNC1T9dO9hfZR5aApUhrK1Hi2g")
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = tel
	request.SignName = "常亭"
	request.TemplateCode = "SMS_170425253"
	request.TemplateParam = fmt.Sprintf("{\"code\":\"%s\"}", code)
	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
		fmt.Printf("response is %#v\n", response)
		return
	}
	key := fmt.Sprintf("code_%s", tel)

	fmt.Printf("手机号是:%s,密码是:%s", tel, code)
	b_cache.Put(key, code, 30*time.Minute)
	fmt.Printf("response is %#v\n", response)

}

//验证注册码是否正确
func TelCodeVerify(cache cache.Cache, code, tel string) bool {
	key := fmt.Sprintf("code_%s", tel)

	key_byte := cache.Get(key)
	if code == string(key_byte.([]byte)) {
		return true
	} else {
		return false
	}
}
