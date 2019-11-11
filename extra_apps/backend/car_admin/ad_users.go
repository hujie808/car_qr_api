package car_admin

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAdUsersTable() table.Table {

    adUsersTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := adUsersTable.GetInfo()
	
	info.AddField("Id","id", db.Int)
	info.AddField("账号","user", db.Varchar)
	info.AddField("名称","username", db.Varchar)
	info.AddField("密码","adpass", db.Varchar)
	info.AddField("过期时间","expire_date", db.Datetime)
	info.AddField("添加时间","add_date", db.Datetime)
	info.AddField("最后访问时间","last_login", db.Datetime)
	info.AddField("最大广告数量","ad_count", db.Int)
	
	info.SetTable("ad_users").SetTitle("商户").SetDescription("商户")

	formList := adUsersTable.GetForm()
	
	formList.AddField("Id","id",db.Int,form.Default).FieldNotAllowAdd()
	formList.AddField("账号","user",db.Varchar,form.Text)
	formList.AddField("名称","username",db.Varchar,form.Text)
	formList.AddField("密码","adpass",db.Varchar,form.Password)
	formList.AddField("过期时间","expire_date",db.Datetime,form.Datetime)
	formList.AddField("添加时间","add_date",db.Datetime,form.Datetime)
	formList.AddField("最后访问时间","last_login",db.Datetime,form.Datetime)
	formList.AddField("最大广告数量","ad_count",db.Int,form.Number)
	
	formList.SetTable("ad_users").SetTitle("Ad_users").SetDescription("Ad_users")

	return adUsersTable
}