package initialization

import (
	"context"
	"github.com/llygcd/block-compensation/config"
	"github.com/llygcd/block-compensation/internal/global"
	"github.com/llygcd/block-compensation/internal/model"
	"github.com/llygcd/block-compensation/internal/model/dto"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"github.com/sirupsen/logrus"
)

func NewQMgo(cfg config.DataBaseConf, ctx context.Context) *qmgo.Client {
	var maxPoolSize uint64 = 4096
	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri:         cfg.Uri,
		MaxPoolSize: &maxPoolSize,
	})
	if err != nil {
		logrus.Fatalf("connect mongo failed, uri: %s, err:%s", cfg.Uri, err.Error())
	}
	return client
}

// MgoCollections create Collections and indexes
func MgoCollections(cli *qmgo.Client) {
	collections := []model.Docs{
		dto.Block{},
		dto.Denom{},
		dto.Nft{},
		dto.Tx{},
	}

	for _, v := range collections {
		createIndexes(cli, v.CollectionName(), v.Indexes())
	}
}

func createIndexes(cli *qmgo.Client, collectionName string, indexes []options.IndexModel) {
	c := cli.Database(global.Config.DataBase.DbName).Collection(collectionName)
	if len(indexes) > 0 {
		for _, v := range indexes {
			if err := c.CreateOneIndex(context.Background(), v); err != nil {
				logrus.Warnf(err.Error())
			}
		}
	}
}
