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
	logGroup.GET("/list", server.BuildHandle(l.listPage))
	logGroup.GET("/data", server.BuildHandle(l.data))
}

func (l *LogController) listPage(ctx *server.GinContext) {
	ctx.HTML("log/list.html", "")
}

func (l *LogController) data(ctx *server.GinContext) {
	page := ctx.QueryInt("page", 0)
	limit := ctx.QueryInt("limit", 10)
	accountName := strings.TrimSpace(ctx.Query("accountName"))

	logs, total := model.LogPagination(accountName, page, limit)
	ctx.Pagination(total, logs)
}
