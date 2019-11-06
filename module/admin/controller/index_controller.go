package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phantacix/go-admin-panel/core/server"
)

type IndexController struct {
}

func (i *IndexController) Init(s *gin.Engine) {
	s.GET("/", server.BuildHandle(i.IndexMain))
	s.GET("/main", server.BuildHandle(i.IndexMain))
}

func (i *IndexController) IndexMain(ctx *server.GinContext) {
	ctx.HTML("home/main.html", "")
}
