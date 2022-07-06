package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/llygcd/block-compensation/config"
	"github.com/llygcd/block-compensation/handlers"
	"github.com/llygcd/block-compensation/internal/initialization"
	"github.com/llygcd/block-compensation/pkg/opb_client"
	"github.com/llygcd/block-compensation/pkg/pool"
	"github.com/sirupsen/logrus"
)

func Serve(cfg *config.Config) {

	mongoDb := initialization.NewQMgo(cfg.DataBase, context.Background())

	/*	if mongoDb != nil {
		initialization.MgoCollections(mongoDb)
	}*/

	pool.Init(cfg)

	opb_client.Conn = opb_client.Init(cfg)

	handlers.InitRouter(cfg)

	poolClient := pool.GetClient()

	repositories := initialization.NewRepositories(mongoDb, cfg.DataBase.DbName)

	services := initialization.NewServices(repositories, poolClient)
	controllers := initialization.NewControllers(services)

	r := gin.Default()
	initialization.Routers(r, controllers)
	logrus.Fatal(r.Run(cfg.Server.App))
}
