package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type AdList struct {
	Id          int       `orm:"column(id);auto" description:"广告内容id"`
	AdUserId    *AdUsers  `orm:"column(ad_user_id);rel(fk);null" description:"广告绑定用户id"`
	AddDate     time.Time `orm:"column(add_date);type(datetime)" description:"添加时间"`
	UpdateDate  time.Time `orm:"column(update_date);type(datetime)" description:"修改时间"`
	Pic         string    `orm:"column(pic);size(200)" description:"图片广告的路片路径"`
	Url         string    `orm:"column(url);size(200);null" description:"广告图片的url跳转链接"`
	Content     string    `orm:"column(content);null" description:"广告内容"`
	Template    string    `orm:"column(template);null" description:"广告模板，为后来预留的数据接口,Tp1、tp2 "`
	DefualtFlag int       `orm:"column(defualt_flag)" description:"0，非默认展示广告  1，默认展示广告"`
}

func (t *AdList) TableName() string {
	return "ad_list"
}

func init() {
	orm.RegisterModel(new(AdList))
}

// AddAdList insert a new AdList into database and returns
// last inserted Id on success.
func AddAdList(m *AdList) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdListById retrieves AdList by Id. Returns error if
// Id doesn't exist
func GetAdListById(id int) (v *AdList, err error) {
	o := orm.NewOrm()
	v = &AdList{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAdList retrieves all AdList matches certain condition. Returns empty list if
// no records exist
func GetAllAdList(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AdList))
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

	var l []AdList
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

// UpdateAdList updates AdList by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdListById(m *AdList) (err error) {
	o := orm.NewOrm()
	v := AdList{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdList deletes AdList by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdList(id int) (err error) {
	o := orm.NewOrm()
	v := AdList{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AdList{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
