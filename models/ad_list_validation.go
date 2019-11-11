package models

import (
	"github.com/astaxie/beego/validation"
	"strings"
)

func (a *AdList) Valid(v *validation.Validation) {
	if strings.Index(a.Template, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}
