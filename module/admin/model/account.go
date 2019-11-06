package model

import (
	"github.com/phantacix/go-admin-panel/core/db"
	"github.com/phantacix/go-admin-panel/core/logger"
)

type Account struct {
	AccountId   int64  `gorm:"column:account_id;primary_key;auto_increment;comment:'帐号id'" json:"accountId"`
	AccountName string `gorm:"column:account_name;comment:'帐号名'" json:"accountName"`
	Name        string `gorm:"column:name;comment:'姓名'" json:"name"`
	Password    string `gorm:"column:password;comment:'密码'" json:"-"`
	DeptId      int64  `gorm:"column:dept_id;comment:'部门id'" json:"deptId"`
	RoleList    string `gorm:"column:role_list;comment:'帐号所属角色id.逗号分隔多个id'" json:"roleList"`
	Status      int    `gorm:"column:status;comment:'帐号状态'" json:"status"`
}

func (*Account) TableName() string {
	return "sys_account"
}

func (a *Account) ISDisabled() bool {
	return a.Status == 1
}

func (a *Account) Save() {
	db.AdminDB().Save(a)
}

func (a *Account) Remove() {
	db.AdminDB().Delete(a)
}

func AccountWithName(accountName string) *Account {
	var account Account
	db.AdminDB().Where("account_name=?", accountName).First(&account)
	return &account
}

func AccountWithId(accountId int64) *Account {
	var account Account
	db.AdminDB().Where("account_id=?", accountId).First(&account)
	return &account
}

func AccountsWithDeptId(deptId int64, page, pageSize int) ([]*Account, int) {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	var accounts []*Account
	var count int

	db.AdminDB().Model(&Account{}).Where("dept_id=?", deptId).Count(&count)

	if count > 0 {
		accounts = make([]*Account, pageSize)

		s := db.AdminDB().Limit(pageSize).Offset((page - 1) * pageSize)
		if err := s.Find(&accounts).Error; err != nil {
			logger.Sugar.Error(err)
		}
	}

	return accounts, count
}

func AccountCount(deptId int64) int {
	var count int
	db.AdminDB().Model(&Account{}).Where("dept_id=?", deptId).Count(&count)
	return count
}
