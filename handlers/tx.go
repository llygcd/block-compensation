package handlers

import (
	"github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/modules/evm"
	msgsdktypes "github.com/kaifei-bianjie/msg-parser/types"
	"github.com/llygcd/block-compensation/config"
	"github.com/llygcd/block-compensation/internal/model/dto"
	"github.com/llygcd/block-compensation/pkg/msgparser"
	"github.com/llygcd/block-compensation/pkg/pool"
	"github.com/llygcd/block-compensation/utils"
	"github.com/llygcd/block-compensation/utils/constant"
	"github.com/sirupsen/logrus"
	types2 "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"golang.org/x/net/context"
	"strings"
	"time"
)

var (
	_parser msgparser.MsgParser
	_conf   *config.Config
)

func InitRouter(conf *config.Config) {
	_conf = conf
	initBech32Prefix(conf)
	router := msgparser.RegisteRouter()
	if conf.Server.OnlySupportModule != "" {
		modules := strings.Split(conf.Server.OnlySupportModule, ",")
		msgRoute := msgparser.NewRouter()
		for _, one := range modules {
			fn, exist := msgparser.RouteHandlerMap[one]
			if !exist {
				logrus.Error("no support module: " + one)
			}
			msgRoute = msgRoute.AddRoute(one, fn)
		}
		if msgRoute.GetRoutesLen() > 0 {
			router = msgRoute
		}

	}
	_parser = msgparser.NewMsgParser(router)
	_conf = conf
}

func ParseBlockAndTxs(b int64, client *pool.Client) (*dto.Block, []*dto.Tx, error) {
	var (
		blockDoc dto.Block
		block    *ctypes.ResultBlock
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if v, err := client.Block(ctx, &b); err != nil {
		time.Sleep(1 * time.Second)
		if v2, err := client.Block(ctx, &b); err != nil {
			return &blockDoc, nil, utils.ConvertErr(b, "", "ParseBlock", err)
		} else {
			block = v2
		}
	} else {
		block = v
	}
	blockDoc = dto.Block{
		Height:   block.Block.Height,
		Time:     block.Block.Time.Unix(),
		Hash:     block.Block.Header.Hash().String(),
		Txn:      int64(len(block.Block.Data.Txs)),
		Proposer: block.Block.ProposerAddress.String(),
	}

	blockResults, err := client.BlockResults(context.Background(), &b)
	if err != nil {
		time.Sleep(1 * time.Second)
		blockResults, err = client.BlockResults(context.Background(), &b)
		if err != nil {
			return &blockDoc, nil, utils.ConvertErr(b, "", "ParseBlockResult", err)
		}
	}

	if len(block.Block.Txs) != len(blockResults.TxsResults) {
		return nil, nil, utils.ConvertErr(b, "", "block.Txs length not equal blockResult", nil)
	}

	txDocs := make([]*dto.Tx, 0, len(block.Block.Txs))
	if len(block.Block.Txs) > 0 {
		for i, v := range block.Block.Txs {
			txResult := blockResults.TxsResults[i]
			txDoc, err := parseTx(uint32(i), v, txResult, block.Block)
			if err != nil {
				return &blockDoc, txDocs, err
			}
			if txDoc.TxHash != "" && len(txDoc.Type) > 0 {
				txDocs = append(txDocs, &txDoc)
			}
		}
	}

	return &blockDoc, txDocs, nil
}

func parseTx(index uint32, txBytes types.Tx, txResult *types2.ResponseDeliverTx, block *types.Block) (dto.Tx, error) {
	var (
		docTx dto.Tx

		docTxMsgs []msgsdktypes.TxMsg
	)
	txHash := utils.BuildHex(txBytes.Hash())

	docTx.Time = block.Time.Unix()
	docTx.Height = block.Height
	docTx.TxHash = txHash
	docTx.Status = parseTxStatus(txResult.Code)
	if docTx.Status == constant.TxStatusFail {
		docTx.Log = txResult.Log
	}

	//docTx.Events = parseEvents(txResult.TxResult.Events)
	docTx.EventsNew = parseABCILogs(txResult.Log)
	docTx.TxIndex = index

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		return docTx, nil
	}
	docTx.GasUsed = txResult.GasUsed
	docTx.Fee = msgsdktypes.BuildFee(authTx.GetFee(), authTx.GetGas())
	docTx.Memo = authTx.GetMemo()

	msgs := authTx.GetMsgs()
	if len(msgs) == 0 {
		return docTx, nil
	}

	for i, v := range msgs {
		msgDocInfo := _parser.HandleTxMsg(v)
		if len(msgDocInfo.Addrs) == 0 {
			continue
		}
		if i == 0 {
			docTx.Type = msgDocInfo.DocTxMsg.Type
		}

		if msgDocInfo.DocTxMsg.Type == "ethereum_tx" {
			var msgEtheumTx evm.DocMsgEthereumTx
			var txData msgparser.LegacyTx
			utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(msgDocInfo.DocTxMsg.Msg), &msgEtheumTx)
			utils.UnMarshalJsonIgnoreErr(msgEtheumTx.Data, &txData)
			docTx.ContractAddrs = append(docTx.ContractAddrs, txData.To)
		}

		docTx.Signers = append(docTx.Signers, removeDuplicatesFromSlice(msgDocInfo.Signers)...)
		docTx.Addrs = append(docTx.Addrs, removeDuplicatesFromSlice(msgDocInfo.Addrs)...)
		docTxMsgs = append(docTxMsgs, msgDocInfo.DocTxMsg)
		docTx.Types = append(docTx.Types, msgDocInfo.DocTxMsg.Type)

	}

	docTx.Addrs = removeDuplicatesFromSlice(docTx.Addrs)
	docTx.Types = removeDuplicatesFromSlice(docTx.Types)
	docTx.Signers = removeDuplicatesFromSlice(docTx.Signers)
	docTx.ContractAddrs = removeDuplicatesFromSlice(docTx.ContractAddrs)
	docTx.DocTxMsgs = docTxMsgs

	// don't save txs which have not parsed
	if docTx.Type == "" {

		return dto.Tx{}, nil
	}

	return docTx, nil
}
func parseTxStatus(code uint32) uint32 {
	if code == 0 {
		return constant.TxStatusSuccess
	} else {
		return constant.TxStatusFail
	}
}

// parseABCILogs attempts to parse a stringified ABCI tx log into a slice of
// EventNe types. It ignore error upon JSON decoding failure.
func parseABCILogs(logs string) []dto.EventNew {
	var res []dto.EventNew
	utils.UnMarshalJsonIgnoreErr(logs, &res)
	return res
}

func removeDuplicatesFromSlice(data []string) (result []string) {
	tempSet := make(map[string]string, len(data))
	for _, val := range data {
		if _, ok := tempSet[val]; ok || val == "" {
			continue
		}
		tempSet[val] = val
	}
	for one := range tempSet {
		result = append(result, one)
	}
	return
}
