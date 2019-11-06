package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/phantacix/go-admin-panel/core/config"
	"github.com/phantacix/go-admin-panel/core/logger"
)

// key:groupId,value:{key:id,value:*gorm.Db}
var ormMap = make(map[string]map[string]*gorm.DB)

var isInit = false

// Init
func Init() {
	if isInit {
		logger.Sugar.Warnf("db is init")
		return
	}

	var mysql = config.Get().MySql
	for _, v := range mysql {
		if !v.Enable {
			continue
		}

		logger.Sugar.Infow("init mysql db.",
			"groupId", v.GroupId,
			"id", v.Id,
			"dbName", v.DbName,
		)

		db, err := createORM(v)
		if err != nil {
			panic(err)
		}

		dbs := ormMap[v.GroupId]
		if dbs == nil {
			dbs = make(map[string]*gorm.DB)
			ormMap[v.GroupId] = dbs
		}
		dbs[v.Id] = db
	}
}

func GetDb(id string) *gorm.DB {
	for _, group := range ormMap {
		for k, v := range group {
			if k == id {
				return v
			}
		}
	}
	return nil
}

func DBWithGroupId(groupId string) map[string]*gorm.DB {
	return ormMap[groupId]
}

func AdminDB() *gorm.DB {
	return GetDb("admin")
}

const (
	connectFormat = "%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

func createORM(cfg *config.MySqlConfig) (*gorm.DB, error) {
	conn := fmt.Sprintf(connectFormat, cfg.UserName, cfg.Password, cfg.Host, cfg.DbName)
	db, err := gorm.Open("mysql", conn)

	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(cfg.MaxIdleConnect)
	db.DB().SetMaxOpenConns(cfg.MaXOpenConnect)
	db.LogMode(cfg.LogMode)

	return db, nil
}
