package routers

import "github.com/astaxie/beego/context"

func condUser(ctx *context.Context) bool {

	if ua := ctx.Request.UserAgent(); ua != "" {
		return true
	}
	return false
}
