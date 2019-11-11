package car_admin

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUserInfoTable() table.Table {

    userInfoTable := table.NewDefaultTable(table.Config{
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
	info := userInfoTable.GetInfo()
	
	info.AddField("Id","id", db.Int)
	info.AddField("手机号","telphone", db.Varchar)
	info.AddField("车牌号","car_num", db.Varchar)
	info.AddField("验证时间","verify_date", db.Datetime)
	info.AddField("更新时间","update_date", db.Datetime)
	
	info.SetTable("user_info").SetTitle("二维码用户信息").SetDescription("二维码用户信息")

	formList := userInfoTable.GetForm()
	
	formList.AddField("Id","id",db.Int,form.Default).FieldNotAllowAdd()
	formList.AddField("Telphone","telphone",db.Varchar,form.Text)
	formList.AddField("Car_num","car_num",db.Varchar,form.Text)
	formList.AddField("Verify_date","verify_date",db.Datetime,form.Datetime)
	formList.AddField("Update_date","update_date",db.Datetime,form.Datetime)
	
	formList.SetTable("user_info").SetTitle("二维码用户信息").SetDescription("二维码用户信息")

	return userInfoTable
}