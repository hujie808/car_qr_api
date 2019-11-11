package car_admin

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAdLogTable() table.Table {

    adLogTable := table.NewDefaultTable(table.Config{
		Driver:     db.DriverMysql,
		//CanAdd:     true, // 是否可以新增
		//Editable:   true, // 是否可以编辑
		//Deletable:  true, // 是否可以删除
		Exportable: true, // 是否可以导出为excel
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})
	info := adLogTable.GetInfo()
	
	info.AddField("Id","id", db.Int)
	info.AddField("二维码uuid","qr_id", db.Int)
	info.AddField("展示广告id","ad_id", db.Int)
	info.AddField("展示时间","show_datetime", db.Datetime)
	info.AddField("是否跳转","click_flag", db.Int)
	
	info.SetTable("ad_log").SetTitle("广告日志").SetDescription("广告日志")

	formList := adLogTable.GetForm()
	
	formList.AddField("Id","id",db.Int,form.Default).FieldNotAllowAdd()
	formList.AddField("Qr_id","qr_id",db.Int,form.Number)
	formList.AddField("Ad_id","ad_id",db.Int,form.Number)
	formList.AddField("Show_datetime","show_datetime",db.Datetime,form.Datetime)
	formList.AddField("Click_flag","click_flag",db.Int,form.Number)
	
	formList.SetTable("ad_log").SetTitle("Ad_log").SetDescription("Ad_log")

	return adLogTable
}