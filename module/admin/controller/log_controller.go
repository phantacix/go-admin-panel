package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phantacix/go-admin-panel/core/server"
	"github.com/phantacix/go-admin-panel/module/admin/model"
	"strings"
)

type LogController struct {
}

func (l *LogController) Init(s *gin.Engine) {
	logGroup := s.Group("/log")
	logGroup.GET("/list", server.BuildHandle(l.List))
	logGroup.GET("/data", server.BuildHandle(l.Data))
}

func (l *LogController) List(ctx *server.GinContext) {
	ctx.HTML("log/list.html", "")
}

func (l *LogController) Data(ctx *server.GinContext) {
	page := ctx.QueryInt("page", 0)
	limit := ctx.QueryInt("limit", 10)
	accountName := strings.TrimSpace(ctx.Query("accountName"))

	var accountId int64

	if len(accountName) > 0 {
		accountId = model.AccountWithName(accountName).AccountId
	}

	logs, total := model.LogWithAccount(accountId, page, limit)
	ctx.Pagination(total, logs)
}
