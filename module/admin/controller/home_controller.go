package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phantacix/go-admin-panel/core/server"
)

type HomeController struct {
}

func (h *HomeController) Init(s *gin.Engine) {
	adminGroup := s.Group("/home")
	adminGroup.GET("/", server.BuildHandle(h.Console))
	adminGroup.GET("/console", server.BuildHandle(h.Console))
	adminGroup.GET("/tree", server.BuildHandle(h.Console))
}

func (h *HomeController) Console(ctx *server.GinContext) {
	ctx.HTML("home/console.html", "")
}

func (h *HomeController) tree(ctx *server.GinContext) {
	ctx.HTML("home/tree_menu.html", "")
}
