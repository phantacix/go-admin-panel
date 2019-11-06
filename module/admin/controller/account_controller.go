package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phantacix/go-admin-panel/core/logger"
	"github.com/phantacix/go-admin-panel/core/server"
	"github.com/phantacix/go-admin-panel/core/utils"
	"github.com/phantacix/go-admin-panel/module/admin/model"
	"net/http"
	"strconv"
)

type AccountController struct {
}

func (a *AccountController) Init(s *gin.Engine) {
	accountGroup := s.Group("/admin/account")

	accountGroup.GET("/logout", server.BuildHandle(a.logout))
	accountGroup.GET("/add", server.BuildHandle(a.add))

	accountGroup.Any("/list", server.BuildHandle(a.list))

	accountGroup.POST("/login", server.BuildHandle(a.login))
	accountGroup.POST("/save", server.BuildHandle(a.save))
	accountGroup.POST("/update", server.BuildHandle(a.update))

	s.GET("/login", server.BuildHandle(a.loginPage))
}

func (a *AccountController) loginPage(ctx *server.GinContext) {
	ctx.HTML("user/login.html", "")
}

func (a *AccountController) login(ctx *server.GinContext) {
	accountName := ctx.DefaultPostForm("accountName", "")
	password := ctx.DefaultPostForm("password", "")

	if accountName == "" || password == "" {
		ctx.Error("帐号名或密码错误")
		return
	}

	account := model.AccountWithName(accountName)
	if account == nil {
		ctx.Error("用户名或密码错误")
		return
	}

	hash := utils.MD5(accountName + password)
	if hash != account.Password {
		ctx.Error("用户名或密码错误")
		return
	}

	if account.ISDisabled() {
		ctx.Error("账号已禁用")
		return
	}

	ctx.SetCookie("token", hash)
	ctx.SetCookie("accountId", strconv.FormatInt(account.AccountId, 10))
	ctx.SetCookie("accountName", accountName)
	ctx.SetCookie("name", utils.EncodeBase64(account.Name))

	ctx.Ok()
}

func (a *AccountController) logout(ctx *server.GinContext) {
	ctx.SetCookie("token", "")
	ctx.SetCookie("accountId", "")
	ctx.SetCookie("accountName", "")
	ctx.SetCookie("name", "")

	ctx.Redirect(http.StatusOK, "/")
}

func (a *AccountController) list(ctx *server.GinContext) {
	deptId := ctx.QueryInt64("deptId", 0)
	offset := ctx.QueryInt("offset", 0)
	limit := ctx.QueryInt("limit", 0)

	if deptId < 1 || offset < 1 || limit < 1 {
		ctx.Error("参数错误")
		return
	}

	accounts, total := model.AccountsWithDeptId(deptId, offset, limit)

	logger.Sugar.Debugf("deptId:%s, offset:%s, limit:%s size:%s", deptId, offset, limit, total)

	ctx.Pagination(total, accounts)
}

func (a *AccountController) add(ctx *server.GinContext) {
	ctx.HTML("home/account/add.html", "")
}

func (a *AccountController) save(ctx *server.GinContext) {
	account := &model.Account{
		Name:        ctx.PostForm("name"),
		AccountName: ctx.PostForm("accountName"),
		Password:    ctx.PostForm("password"),
		DeptId:      ctx.PostInt64("deptId", 0),
		Status:      ctx.PostInt("status", 0),
		RoleList:    ctx.PostForm("roleIds"),
	}

	if account.DeptId < 1 {
		ctx.Error("请选择部门")
		return
	}

	account.Save()

	if account.AccountId < 1 {
		ctx.Error("保存失败")
		return
	}

	ctx.Ok()
}

func (a *AccountController) update(ctx *server.GinContext) {
	accountId := ctx.PostInt64("accountId", 0)

	if accountId < 1 {
		ctx.Error("用户不存在")
		return
	}

}
