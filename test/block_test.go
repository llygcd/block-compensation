package test

import (
	"context"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/utils"
	"github.com/llygcd/block-compensation/config"
	"github.com/llygcd/block-compensation/handlers"
	"github.com/llygcd/block-compensation/internal/initialization"
	"github.com/llygcd/block-compensation/pkg/pool"
	"io/ioutil"
	"testing"
)

func TestGetBlock(t *testing.T) {
	data, err := ioutil.ReadFile("/home/lly/go/src/block-compensation/config/cfg.toml")
	if err != nil {
		panic(err)
	}
	config, err := config.ReadConfig(data)
	pool.Init(config)

	var height = int64(5536555)
	client := pool.GetClient()

	handlers.InitRouter(config)
	blockDoc, txDocs, err := handlers.ParseBlockAndTxs(height, client)

	mgoSupTech := initialization.NewQMgo(config.DataBase, context.Background())
	repositories := initialization.NewRepositories(mgoSupTech, config.DataBase.DbName)

	repositories.BolckRepo.Save(blockDoc)
	repositories.TxRepo.Save(txDocs)
	fmt.Println(utils.MarshalJsonIgnoreErr(blockDoc))
	fmt.Println(utils.MarshalJsonIgnoreErr(txDocs))
}
