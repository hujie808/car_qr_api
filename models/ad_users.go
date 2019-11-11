package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type AdUsers struct {
	Id            int       `orm:"column(id);auto" description:"广告客户id"`
	User          string    `orm:"column(user);size(64)" description:"广告客户账号"`
	Username      string    `orm:"column(username);size(64)" description:"广告客户显示名称"`
	Adpass        string    `orm:"column(adpass);size(64)" description:"客户密码"`
	ExpireDate    time.Time `orm:"column(expire_date);type(datetime)" description:"账户过期时间"`
	AddDate       time.Time `orm:"column(add_date);type(datetime)" description:"账号添加时间"`
	LastLogin     time.Time `orm:"column(last_login);type(datetime);null" description:"最后访问时间"`
	AdCount       int       `orm:"column(ad_count)" description:"设置广告数量"`
	DefaultAdList *AdList   `orm:"column(default_ad_list);rel(fk)" description:"默认广告id"`
}

func (t *AdUsers) TableName() string {
	return "ad_users"
}

func init() {
	orm.RegisterModel(new(AdUsers))
}

// AddAdUsers insert a new AdUsers into database and returns
// last inserted Id on success.
func AddAdUsers(m *AdUsers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdUsersById retrieves AdUsers by Id. Returns error if
// Id doesn't exist
func GetAdUsersById(id int) (v *AdUsers, err error) {
	o := orm.NewOrm()
	v = &AdUsers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAdUsers retrieves all AdUsers matches certain condition. Returns empty list if
// no records exist
func GetAllAdUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AdUsers))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []AdUsers
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAdUsers updates AdUsers by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdUsersById(m *AdUsers) (err error) {
	o := orm.NewOrm()
	v := AdUsers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdUsers deletes AdUsers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdUsers(id int) (err error) {
	o := orm.NewOrm()
	v := AdUsers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AdUsers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
