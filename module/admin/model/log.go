package model

import (
	"github.com/phantacix/go-admin-panel/core/db"
	"github.com/phantacix/go-admin-panel/core/logger"
	"github.com/phantacix/go-admin-panel/core/utils"
)

type Log struct {
	Id          int64  `gorm:"column:id;primary_key;auto_increment;comment:'日志自增id'" json:"id"`
	AccountId   int64  `gorm:"column:account_id;comment:'用户id'" json:"accountId"`
	AccountName string `gorm:"column:account_name;index;size:60;comment:'帐号名'" json:"accountName"`
	Method      string `gorm:"column:method;ize:255;comment:'执行方法'" json:"method"`
	Path        string `gorm:"column:path;ize:512;comment:'执行路径'" json:"path"`
	Params      string `gorm:"column:params;size:512;comment:'请求参数'" json:"params"`
	Ip          string `gorm:"column:ip;comment:'访问ip'" json:"ip"`
	CreateTime  int    `gorm:"column:create_time;comment:'创建时间'" json:"createTime"`
}

func (*Log) TableName() string {
	return "sys_log"
}

func LogCreate(accountId int64, accountName string, method string, path string, params string, ip string) {
	log := &Log{
		AccountId:   accountId,
		AccountName: accountName,
		Method:      method,
		Path:        path,
		Params:      params,
		Ip:          ip,
		CreateTime:  utils.NowMillisecond(),
	}
	db.AdminDB().Create(log)
}

func LogPagination(accountName string, page int, pageSize int) ([]*Log, int) {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	var logs []*Log
	var count int

	if len(accountName) > 0 {
		db.AdminDB().Model(&Log{}).Where("account_name=?", accountName).Count(&count)
	} else {
		db.AdminDB().Model(&Log{}).Count(&count)
	}

	if count > 0 {
		logs = make([]*Log, pageSize)

		s := db.AdminDB().Limit(pageSize).Offset((page - 1) * pageSize)

		if len(accountName) > 0 {
			s.Where("account_name=?", accountName)
		}

		if err := s.Find(&logs).Error; err != nil {
			logger.Sugar.Error(err)
		}
	}

	return logs, count
}
