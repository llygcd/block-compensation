package dto

import (
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type Block struct {
	Height   int64  `bson:"height"`
	Hash     string `bson:"hash"`
	Txn      int64  `bson:"txn"`
	Time     int64  `bson:"time"`
	Proposer string `bson:"proposer"`
}

func (b Block) CollectionName() string {
	return "sync_block"
}

func (b Block) Indexes() (indexes []options.IndexModel) {
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-height"},
		Unique:     true,
		Background: true,
	})
	return
}

func (b Block) PkKvPair() map[string]interface{} {
	return bson.M{"height": b.Height}
}
