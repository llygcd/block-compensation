package dto

import (
	"github.com/kaifei-bianjie/msg-parser/types"
	"github.com/llygcd/block-compensation/internal/global"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	Tx struct {
		Time          int64         `bson:"time"`
		Height        int64         `bson:"height"`
		TxHash        string        `bson:"tx_hash"`
		Type          string        `bson:"type"` // parse from first msg
		Memo          string        `bson:"memo"`
		Status        uint32        `bson:"status"`
		Log           string        `bson:"log"`
		Fee           *types.Fee    `bson:"fee"`
		GasUsed       int64         `bson:"gas_used"`
		Types         []string      `bson:"types"`
		EventsNew     []EventNew    `bson:"events_new"`
		Signers       []string      `bson:"signers"`
		DocTxMsgs     []types.TxMsg `bson:"msgs"`
		Addrs         []string      `bson:"addrs"`
		ContractAddrs []string      `bson:"contract_addrs"`
		TxIndex       uint32        `bson:"tx_index"`
		Ext           interface{}   `bson:"ext"`
	}

	Event struct {
		Type       string   `bson:"type"`
		Attributes []KvPair `bson:"attributes"`
	}

	KvPair struct {
		Key   string `bson:"key"`
		Value string `bson:"value"`
	}

	EventNew struct {
		MsgIndex uint32  `bson:"msg_index" json:"msg_index"`
		Events   []Event `bson:"events"`
	}
)

func (t Tx) CollectionName() string {
	return global.GetServerConf().ChainId + "_tx"
}

func (t Tx) Indexes() (indexes []options.IndexModel) {
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-tx_hash"},
		Unique:     true,
		Background: true,
	})

	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-height", "-msgs.type", "-status"},
		Unique:     false,
		Background: true,
	})

	return
}

func (t Tx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": t.TxHash}
}