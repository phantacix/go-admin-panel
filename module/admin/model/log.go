package model

import "github.com/phantacix/go-admin-panel/core/db"

type Log struct {
	Id         int64  `gorm:"column:id;primary_key;auto_increment;comment:'日志自增id'" json:"id"`
	AccountId  int64  `gorm:"column:account_id;comment:'用户id'" json:"accountId"`
	Method     string `gorm:"column:method;comment:'执行方法'" json:"method"`
	Path       string `gorm:"column:path;comment:'执行路径'" json:"path"`
	Params     string `gorm:"column:params;comment:'请求参数'" json:"params"`
	Ip         string `gorm:"column:ip;comment:'访问ip'" json:"ip"`
	CreateTime int64  `gorm:"column:create_time;comment:'创建时间'" json:"createTime"`
}

func (*Log) TableName() string {
	return "sys_log"
}

func LogPagination(page int, limit int) ([]*Log, int) {
	offset := (page - 1) * limit
	logs := make([]*Log, 0)
	total := 0
	db.AdminDB().Limit(limit).Offset(offset).Find(&logs).Count(&total)
	return logs, total
}

func LogDelete(id int) {
	var log Log
	db.AdminDB().Delete(&log, "id=?", id)
}

func LogWithAccountName(accountName string, page int, limit int) ([]*Log, int) {
	offset := (page - 1) * limit
	logs := make([]*Log, 0)
	total := 0
	db.AdminDB().Where("accountName=?", accountName).Limit(limit).Offset(offset).Find(&logs).Count(&total)
	return logs, total
}
