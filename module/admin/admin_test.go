package admin

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/phantacix/go-admin-panel/core"
	"github.com/phantacix/go-admin-panel/core/db"
	"github.com/phantacix/go-admin-panel/module/admin/model"
	"testing"
)

func TestAdminDBMigrate(t *testing.T) {

	core.Init("../../../config/env/local/admin.json")

	orm := db.AdminDB()

	dept, log, menu, role, account := &model.Dept{}, &model.Log{}, &model.Menu{}, &model.Role{}, &model.Account{}
	orm.DropTableIfExists(dept, log, menu, role, account)

	if err := orm.AutoMigrate(dept, log, menu, role, account).Error; err != nil {
		fmt.Print(err)
	}

}
