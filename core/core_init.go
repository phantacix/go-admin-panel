package core

import (
	"github.com/phantacix/go-admin-panel/core/config"
	"github.com/phantacix/go-admin-panel/core/db"
	"github.com/phantacix/go-admin-panel/core/logger"
)

func Init(cfgPath string) {
	config.LoadConfig(cfgPath)
	logger.NewZap(config.Get().Logger)
	db.Init()

	logger.Sugar.Info("config & logger & db is init complete!")
}
