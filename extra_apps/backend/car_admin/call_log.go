package car_admin

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCallLogTable() table.Table {

    callLogTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := callLogTable.GetInfo()
	
	info.AddField("Id","id", db.Int)
	info.AddField("拨打时间","call_date", db.Datetime)
	info.AddField("拨打状态","call_status", db.Int)
	info.AddField("二维码id","qr_id", db.Int)
	info.AddField("拨打","call_from", db.Varchar)
	
	info.SetTable("call_log").SetTitle("拨号日志").SetDescription("拨号日志")

	formList := callLogTable.GetForm()
	
	formList.AddField("Id","id",db.Int,form.Default).FieldNotAllowAdd()
	formList.AddField("Call_date","call_date",db.Datetime,form.Datetime)
	formList.AddField("Call_status","call_status",db.Int,form.Number)
	formList.AddField("Qr_id","qr_id",db.Int,form.Number)
	formList.AddField("Call_from","call_from",db.Varchar,form.Text)
	
	formList.SetTable("call_log").SetTitle("Call_log").SetDescription("Call_log")

	return callLogTable
}