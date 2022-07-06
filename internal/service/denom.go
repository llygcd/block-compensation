package service

import (
	"github.com/kaifei-bianjie/msg-parser/modules/nft"
	"github.com/kaifei-bianjie/msg-parser/types"
	"github.com/llygcd/block-compensation/internal/model/dto"
	"github.com/llygcd/block-compensation/internal/repository"
	"github.com/llygcd/block-compensation/pkg/opb_client"
	"github.com/qiniu/qmgo"
	"time"
)

type DenomIssue struct {
	repo repository.IDenomRepo
}

func (n *DenomIssue) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var issueDenom = msg.Msg.(*nft.DocMsgIssueDenom)
	denom, err := opb_client.QueryDenom(issueDenom.Id)
	if err != nil {
		return err
	}

	one, err := n.repo.FindOne(denom.Denom.Id)
	if err != nil {
		if err == qmgo.ErrNoSuchDocuments {
			return n.repo.Save(&dto.Denom{
				Name:            denom.Denom.Name,
				DenomID:         denom.Denom.Id,
				JsonSchema:      denom.Denom.Schema,
				Creator:         issueDenom.Sender,
				Owner:           denom.Denom.Owner,
				Txhash:          tx.TxHash,
				Height:          int(tx.Height),
				Time:            int(tx.Time),
				CreateTime:      int(time.Now().Unix()),
				LastBlockHeight: int(tx.Height),
				LastBlockTime:   int(tx.Time),
			})
		}
		return err
	}
	one.Name = denom.Denom.Name
	one.DenomID = denom.Denom.Id
	one.JsonSchema = denom.Denom.Schema
	one.Creator = issueDenom.Sender
	one.Owner = denom.Denom.Owner
	one.Txhash = tx.TxHash
	one.Height = int(tx.Height)
	one.Time = int(tx.Time)
	one.LastBlockHeight = int(tx.Height)
	one.LastBlockTime = int(tx.Time)
	return n.repo.Update(one)
}

type DenomTransfer struct {
	repo repository.IDenomRepo
}

func (n *DenomTransfer) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var transferDenom = msg.Msg.(*nft.DocMsgTransferDenom)
	denom, err := opb_client.QueryDenom(transferDenom.Id)
	if err != nil {
		return err
	}

	one, err := n.repo.FindOne(denom.Denom.Id)
	if err != nil {
		if err == qmgo.ErrNoSuchDocuments {
			return n.repo.Save(&dto.Denom{
				Name:            denom.Denom.Name,
				DenomID:         denom.Denom.Id,
				JsonSchema:      denom.Denom.Schema,
				Creator:         transferDenom.Sender,
				Owner:           denom.Denom.Owner,
				Txhash:          tx.TxHash,
				Height:          int(tx.Height),
				Time:            int(tx.Time),
				CreateTime:      int(time.Now().Unix()),
				LastBlockHeight: int(tx.Height),
				LastBlockTime:   int(tx.Time),
			})
		}
		return err
	}

	one.Name = denom.Denom.Name
	one.DenomID = denom.Denom.Id
	one.JsonSchema = denom.Denom.Schema
	one.Owner = denom.Denom.Owner
	one.Txhash = tx.TxHash
	one.Height = int(tx.Height)
	one.Time = int(tx.Time)
	one.LastBlockHeight = int(tx.Height)
	one.LastBlockTime = int(tx.Time)
	return n.repo.Update(one)
}
