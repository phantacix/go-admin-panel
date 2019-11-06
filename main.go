package main

import (
	"flag"
	"github.com/phantacix/go-admin-panel/core/config"
	"github.com/phantacix/go-admin-panel/core/db"
	"github.com/phantacix/go-admin-panel/core/logger"
	"github.com/phantacix/go-admin-panel/core/server"
	"github.com/phantacix/go-admin-panel/module/admin/controller"
)

var (
	env string
)

func init() {
	flag.StringVar(&env, "env", "local", "-env=local set running environment(test/prod)")
	flag.Parse()
}

func main() {
	config.Parse(env)

	logger.NewZap(config.Get().Logger)

	db.Init()

	server.Web = server.New()
	server.Web.AddController(
		&controller.DeptController{},
		&controller.HomeController{},
		&controller.IndexController{},
		&controller.LogController{},
		&controller.AccountController{},
	)
	server.Web.Run()
}
