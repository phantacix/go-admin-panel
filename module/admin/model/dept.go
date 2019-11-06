package model

import "github.com/phantacix/go-admin-panel/core/db"

type Dept struct {
	DeptId   int64  `gorm:"column:dept_id;primary_key;auto_increment;comment:'部门id'"`
	ParentId int64  `gorm:"column:parent_id;comment:'部门父id'"`
	Name     string `gorm:"column:name;comment:'部门名称'"`
	OrderNum int    `gorm:"column:order_num;comment:'排序id'"`
	DelFlag  int8   `gorm:"column:del_flag;comment:'删除标识'"`
}

func (*Dept) TableName() string {
	return "sys_dept"
}

func (d *Dept) Get(deptId int64) {
	db.AdminDB().Where("dept_id=?", deptId).First(d)
}

func (d *Dept) Save() {
	db.AdminDB().Save(d)
}

func (d *Dept) Update() {
	db.AdminDB().Update(d)
}

func (d *Dept) Remove() {
	db.AdminDB().Delete(d)
}

func (d *Dept) Insert() {
	db.AdminDB().Create(d)
}

func DeptList() []*Dept {
	var dept []*Dept
	db.AdminDB().Find(&dept)
	return dept
}
