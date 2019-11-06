package model

import (
	"fmt"
	"github.com/phantacix/go-admin-panel/core"
	"github.com/phantacix/go-admin-panel/core/db"
	"testing"
)

func init() {
	core.Init("../../../config/env/local/admin.json")
}

func TestAccountWithName(t *testing.T) {
	var account = AccountWithName("su")
	if account == nil {
		t.Error("not found account")
	}
	fmt.Print(account)
}

func TestAccountWithId(t *testing.T) {
	var account = AccountWithId(1)
	if account == nil {
		t.Error("not found account")
	}
	fmt.Print(account)
}

func TestAccountsWithDeptId(t *testing.T) {
	accounts, total := AccountsWithDeptId(1, 1, 10)
	if accounts == nil {
		t.Error("not found account")
	}
	fmt.Print(accounts, total)
}

func TestAccountCount(t *testing.T) {
	var count = AccountCount(1)
	fmt.Print(count)
}

func TestAccount_Save1(t *testing.T) {
	var account = AccountWithName("su")
	fmt.Print(account)
	account.Status = 2
	account.Save()

	account = AccountWithName("su")
	fmt.Print(account)
}

func TestAccount_Save(t *testing.T) {
	account := &Account{
		AccountName: "su1",
		Name:        "ryan su",
		Password:    "password",
		DeptId:      1,
		Status:      0,
	}

	account.Save()
	fmt.Print(account)
}

func TestAccountOrm(t *testing.T) {
	first := &Account{}
	db.AdminDB().Select("*").Where("account_name=?", "su1").Find(first)
	fmt.Print(first)

	m := &Account{}
	db.AdminDB().Model(m).Where("account_id > 0").Update("account_name", "su")
	fmt.Print(m)

	db.AdminDB().Table("sys_account").Where("account_id > 0").Update("account_name", "su")
}
