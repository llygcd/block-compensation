package dto

import (
	"github.com/llygcd/block-compensation/internal/global"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	DenomDenomID = "denom_id"
)

type Denom struct {
	Type            string `bson:"type"`
	Name            string `bson:"name"`
	DenomID         string `bson:"denom_id"`
	JsonSchema      string `bson:"json_schema"`
	Creator         string `bson:"creator"`
	Owner           string `bson:"owner"`
	Txhash          string `bson:"tx_hash"`
	Height          int64  `bson:"height"`
	Time            int64  `bson:"time"`
	CreateTime      int64  `bson:"create_time"`
	LastBlockHeight int64  `bson:"last_block_height"`
	LastBlockTime   int64  `bson:"last_block_time"`
}

func (d Denom) CollectionName() string {
	return global.GetServerConf().ChainId + "_denom"
}

func (d Denom) Indexes() (indexes []options.IndexModel) {
	indexes = append(indexes, options.IndexModel{
		Key:        []string{DenomDenomID},
		Unique:     true,
		Background: true,
	})
	return
}

func (d Denom) PkKvPair() map[string]interface{} {
	return bson.M{DenomDenomID: d.DenomID}
}

func (d Denom) SortKey() string {
	return d.DenomID
}

type Denoms []Denom
