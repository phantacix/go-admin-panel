package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//web server instance
var Web *WebServer

type IController interface {
	Init(*gin.Engine)
}

// 包装为，自定义的GinContext
func BuildHandle(handler func(ctx *GinContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(GinContext)
		context.Context = c
		context.AuthCookie = &AuthCookie{}
		context.AuthCookie.Parse(context)
		handler(context)
	}
}

type GinContext struct {
	*gin.Context
	AuthCookie *AuthCookie
}

func (c *GinContext) QueryInt(name string, defaultValue int) int {
	if value, ok := c.GetQuery(name); ok {
		if v, e := strconv.Atoi(value); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) QueryInt64(name string, defaultValue int64) int64 {
	if value, ok := c.GetQuery(name); ok {
		if v, e := strconv.ParseInt(value, 10, 64); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) PostInt(name string, defaultValue int) int {
	if value, ok := c.GetPostForm(name); ok {
		if v, e := strconv.Atoi(value); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) PostInt64(name string, defaultValue int64) int64 {
	if value, ok := c.GetPostForm(name); ok {
		if v, e := strconv.ParseInt(value, 10, 64); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) SetCookie(name, value string) {
	c.Context.SetCookie(name, value, -1, "", "", false, false)
}

func (c *GinContext) CookieInt(name string, defaultValue int) int {
	if value, err := c.Context.Cookie(name); err == nil {
		if v, e := strconv.Atoi(value); e == nil {
			return v
		}
	}
	return defaultValue
}

func (c *GinContext) CookieInt64(name string, defaultValue int64) int64 {
	if value, err := c.Context.Cookie(name); err == nil {
		if v, e := strconv.ParseInt(value, 10, 64); e == nil {
			return v
		}
	}

	return defaultValue
}

func (c *GinContext) Error(msg string) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  msg,
	})
}

func (c *GinContext) Ok() {
	c.JSON("")
}

func (c *GinContext) JSON(value interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  value,
	})
}

func (c *GinContext) Pagination(total int, rows interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "",
		"count": total,
		"data":  rows,
	})
}

func (c *GinContext) PaginationMsg(total int, rows interface{}, msg string) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   msg,
		"total": total,
		"rows":  rows,
	})
}

func (c *GinContext) Object(obj interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": obj,
	})
}

func (c *GinContext) HTML(name string, obj interface{}) {
	c.Context.HTML(http.StatusOK, name, obj)
}
