package global

import (
	"github.com/llygcd/block-compensation/config"
	"github.com/sirupsen/logrus"
)

var (
	Config *config.Config
)

func GetServerConf() *config.ServerConf {
	if Config == nil {
		logrus.Error("db.Init not work")
	}
	return &Config.Server
}
