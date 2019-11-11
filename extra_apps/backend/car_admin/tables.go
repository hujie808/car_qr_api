package car_admin

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "ad_list" => http://localhost:9033/admin/info/ad_list
// "ad_log" => http://localhost:9033/admin/info/ad_log
// "ad_users" => http://localhost:9033/admin/info/ad_users
// "call_log" => http://localhost:9033/admin/info/call_log
// "qr_info" => http://localhost:9033/admin/info/qr_info
// "user_info" => http://localhost:9033/admin/info/user_info
//
var Generators = map[string]table.Generator{
	"ad_list": GetAdListTable,
	"ad_log": GetAdLogTable,
	"ad_users": GetAdUsersTable,
	"call_log": GetCallLogTable,
	"qr_info": GetQrInfoTable,
	"user_info": GetUserInfoTable,
}
