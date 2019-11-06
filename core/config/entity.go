package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type AdminConfig struct {
	Debug  bool           `json:"debug"`
	Logger *LoggerConfig  `json:"logger"`
	Gin    *GinConfig     `json:"gin`
	MySql  []*MySqlConfig `json:"mysql"`
	Redis  []*RedisConfig `json:"redis"`
}

func (a *AdminConfig) IsDebug() bool {
	return a.Debug
}

type MySqlConfig struct {
	Enable         bool   `json:"enable"`
	GroupId        string `json:"groupId"`
	Id             string `json:"id"`
	DbName         string `json:"dbName"`
	Host           string `json:"host"`
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	MaxIdleConnect int    `json:"maxIdleConnect"`
	MaXOpenConnect int    `json:"maxOpenConnect"`
	LogMode        bool   `json:"logMode"`
}

type RedisConfig struct {
	Id             string `json:"id"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Password       string `json:"password"`
	MaxConnect     int    `json:"maxConnect"`
	MaxIdleConnect int    `json:"maxIdleConnect"`
	MinIdleConnect int    `json:"minIdleConnect"`
}

type GinConfig struct {
	View               string `json:"view"`
	StaticRelativePath string `json:"staticRelativePath"`
	StaticRootPath     string `json:"staticRootPath"`
	Favicon            string `json:"favicon"`
	FaviconPath        string `json:"faviconPath"`
	URL                string `json:"url"`
	Port               int    `json:"port"`
}

type LoggerConfig struct {
	Path         string `json:"path"`
	Suffix       string `json:"suffix"`
	Level        string `json:"level"`
	IsWriteFile  bool   `json:"isWriteFile"`
	MaxAge       int    `json:"maxAge"`
	RotationHour int    `json:"rotationHour"`
}

func (l *LoggerConfig) GetFileName() string {
	return l.Level + l.Suffix
}

func (l *LoggerConfig) GetLevel() zapcore.Level {
	switch strings.ToLower(l.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
