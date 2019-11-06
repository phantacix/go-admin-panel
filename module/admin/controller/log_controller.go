package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phantacix/go-admin-panel/core/server"
	"github.com/phantacix/go-admin-panel/module/admin/model"
	"strconv"
	"strings"
)

type LogController struct {
}

func (l *LogController) Init(s *gin.Engine) {
	logGroup := s.Group("/log")
	logGroup.GET("/list", server.BuildHandle(l.List))
	logGroup.GET("/data", server.BuildHandle(l.Data))
	logGroup.GET("/delete", server.BuildHandle(l.Delete))
	logGroup.GET("/list/accountName", server.BuildHandle(l.LogWithAccountName))

}

func (l *LogController) List(ctx *server.GinContext) {
	ctx.HTML("log/list.html", "")
}

func (l *LogController) Delete(ctx *server.GinContext) {
	ids := ctx.DefaultPostForm("ids", "")
	idSplit := strings.Split(ids, ",")
	idList := make([]int, len(idSplit))

	for i := 0; i < len(idSplit); i++ {
		id, err := strconv.Atoi(idSplit[i])
		if err != nil {
			ctx.Error("ID格式错误")
			return
		}
		idList[i] = id
	}

	if len(idList) <= 0 {
		ctx.Error("所选择ID不能为空")
		return
	}

	for i := 0; i < len(idList); i++ {
		model.LogDelete(idList[i])
	}
	ctx.Ok()
}

func (l *LogController) Data(ctx *server.GinContext) {
	page := ctx.QueryInt("page", 0)
	limit := ctx.QueryInt("limit", 10)

	logs, total := model.LogPagination(page, limit)

	ctx.Pagination(total, logs)
}

func (l *LogController) LogWithAccountName(ctx *server.GinContext) {
	accountName := ctx.DefaultPostForm("accountName", "")
	if accountName == "" {
		return
	}
	page := ctx.QueryInt("page", 0)
	limit := ctx.QueryInt("limit", 10)

	logs, total := model.LogWithAccountName(accountName, page, limit)

	ctx.Pagination(total, logs)
}
