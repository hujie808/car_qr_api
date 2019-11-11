package car_admin

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetQrInfoTable() table.Table {

    qrInfoTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := qrInfoTable.GetInfo()
	
	info.AddField("Id","id", db.Int)
	info.AddField("Uuid","uuid", db.Varchar)
	info.AddField("是否使用","inuse", db.Int)
	info.AddField("关联用户id","user_id", db.Int)
	info.AddField("关联广告id","ad_id", db.Int)
	info.AddField("启动时间","start_date", db.Datetime)
	info.AddField("最新访问时间","visit_date", db.Datetime)
	info.AddField("访问次数","visit_times", db.Int)
	info.AddField("所属广告商","user", db.Varchar).FieldJoin(types.Join{//要显示关联表的字段
		Table:     "ad_users",    // 连表的表名
		Field:     "ad_users_id",      // 本表的字段
		JoinField: "id", // 连表的表的字段
	})
	
	info.SetTable("qr_info").SetTitle("Qr_info").SetDescription("Qr_info")

	formList := qrInfoTable.GetForm()
	
	formList.AddField("Id","id",db.Int,form.Default).FieldNotAllowAdd()
	formList.AddField("Uuid","uuid",db.Varchar,form.Text)
	formList.AddField("Inuse","inuse",db.Int,form.Number)
	formList.AddField("User_id","user_id",db.Int,form.Number)
	formList.AddField("Ad_id","ad_id",db.Int,form.Number)
	formList.AddField("Start_date","start_date",db.Datetime,form.Datetime)
	formList.AddField("Visit_date","visit_date",db.Datetime,form.Datetime)
	formList.AddField("Visit_times","visit_times",db.Int,form.Number)
	formList.AddField("Ad_users_id","ad_users_id",db.Int,form.Number)
	
	formList.SetTable("qr_info").SetTitle("二维码信息").SetDescription("二维码信息")

	return qrInfoTable
}