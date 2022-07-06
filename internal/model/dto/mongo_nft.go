package dto

import (
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	NftDenomID = "denom_id"
	NftNftID   = "nft_id"
)

type (
	Nft struct {
		DenomID         string `bson:"denom_id"`
		NftID           string `bson:"nft_id"`
		CreateTime      int    `bson:"create_time"`
		Data            string `bson:"data"`
		DenomName       string `bson:"denom_name"`
		LastBlockHeight int    `bson:"last_block_height"`
		LastBlockTime   int    `bson:"last_block_time"`
		NftName         string `bson:"nft_name"`
		Owner           string `bson:"owner"`
		UpdateTime      int    `bson:"update_time"`
		URI             string `bson:"uri"`
	}
)

func (n Nft) CollectionName() string {
	return "ex_sync_nft"
}

func (n Nft) Indexes() (indexes []options.IndexModel) {
	indexes = append(indexes, options.IndexModel{
		Key:        []string{NftNftID, NftDenomID},
		Unique:     true,
		Background: true,
	})

	return
}

func (n Nft) PkKvPair() map[string]interface{} {
	return bson.M{
		NftDenomID: n.DenomID,
		NftNftID:   n.NftID,
	}
}
