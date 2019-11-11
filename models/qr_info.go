package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type QrInfo struct {
	Id         int       `orm:"column(id);auto" description:"id"`
	Uuid       string    `orm:"column(uuid);size(64)" description:"uuid"`
	Inuse      int       `orm:"column(inuse)" description:"0,1当前是否使用"`
	UserId     *UserInfo `orm:"column(user_id);rel(fk);null" description:"关联用户表的用户id "`
	AdId       *AdList   `orm:"column(ad_id);rel(fk);null" description:"关联的广告id "`
	StartDate  time.Time `orm:"column(start_date);type(datetime);null" description:"启用时间"`
	VisitDate  time.Time `orm:"column(visit_date);type(datetime);null" description:"最新访问时间 "`
	VisitTimes int       `orm:"column(visit_times)" description:"访问次数"`
	AdUsersId  *AdUsers  `orm:"column(ad_users_id);rel(fk);null" description:"所属广告商"`
}

func (t *QrInfo) TableName() string {
	return "qr_info"
}

func init() {
	orm.RegisterModel(new(QrInfo))
}

// AddQrInfo insert a new QrInfo into database and returns
// last inserted Id on success.
func AddQrInfo(m *QrInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetQrInfoById retrieves QrInfo by Id. Returns error if
// Id doesn't exist
func GetQrInfoById(id int) (v *QrInfo, err error) {
	o := orm.NewOrm()
	v = &QrInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllQrInfo retrieves all QrInfo matches certain condition. Returns empty list if
// no records exist
func GetAllQrInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(QrInfo))
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

	var l []QrInfo
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

// UpdateQrInfo updates QrInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateQrInfoById(m *QrInfo) (err error) {
	o := orm.NewOrm()
	v := QrInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteQrInfo deletes QrInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteQrInfo(id int) (err error) {
	o := orm.NewOrm()
	v := QrInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&QrInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
