package car_admin

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"strconv"
	"time"
)

func GetAdListTable() table.Table {

	adListTable := table.NewDefaultTable(table.Config{
		Driver:     db.DriverMysql,
		CanAdd:     true, // 是否可以新增
		Editable:   true, // 是否可以编辑
		Deletable:  true, // 是否可以删除
		Exportable: true, // 是否可以导出为excel
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := adListTable.GetInfo()

	info.AddField("Id", "id", db.Int)
	info.AddField("广告绑定用户id", "user", db.Varchar).FieldJoin(types.Join{//要显示关联表的字段
		Table:     "ad_users",    // 连表的表名
		Field:     "ad_user_id",      // 本表的字段
		JoinField: "id", // 连表的表的字段
	})
	info.AddField("添加时间", "add_date", db.Datetime)
	info.AddField("修改时间", "update_date", db.Datetime)
	info.AddField("图片广告的路片路径", "pic", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return fmt.Sprintf(`<img src="%s" width="175" height="47" />`,value.Value)
	})
	info.AddField("广告图片的url跳转链接", "url", db.Varchar)
	//info.AddField("广告内容", "content", db.Text)
	info.AddField("广告模板，预留接口", "template", db.Text)
	info.AddField("默认广告", "defualt_flag", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "1" {
			return "🔴"
		} else if model.Value == "0" {
			return "⭕"
		}
		return "unknown"
	})
	info.SetTable("ad_list").SetTitle("广告列表").SetDescription("广告列表")


	var roles,defualt_flag []map[string]string
	rolesModel, _ := db.Table("ad_users").Select("id", "user").Where("id", ">", 0).All()//找到ad_users所有字段
	for _, v := range rolesModel {
		roles = append(roles, map[string]string{
			"field": v["user"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}
	defualt_flag=append(defualt_flag, map[string]string{"field": "默认", "value": "1",},map[string]string{"field": "非默认", "value": "0",})

	formList := adListTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("广告绑定用户id", "ad_user_id", db.Int, form.SelectSingle).FieldOptions(roles)
	formList.AddField("图片广告的路片路径", "pic", db.Varchar, form.File).FieldPostFilterFn(func(model types.PostFieldModel) string {
			return "/static/admin/"+model.Value.Value()
		})
	formList.AddField("广告图片的url跳转链接", "url", db.Varchar, form.Text)
	formList.AddField("广告内容", "content", db.Text, form.RichText)
	formList.AddField("广告模板，预留接口", "template", db.Text, form.Text)
	formList.AddField("默认广告", "defualt_flag", db.Int, form.SelectSingle).FieldDefaultOptionDelimiter("0").FieldOptions(defualt_flag)
	formList.AddField("添加时间", "add_date", db.Datetime, form.Datetime).FieldDefault(time.Now().String())
	formList.AddField("修改时间", "update_date", db.Datetime, form.Datetime).FieldDefault(time.Now().String())

	formList.SetTable("ad_list").SetTitle("广告列表").SetDescription("广告列表")

	return adListTable
}
