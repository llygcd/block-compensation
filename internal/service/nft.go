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

type NftMint struct {
	repo repository.INftRepo
}

func (n *NftMint) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var nftMint = msg.Msg.(*nft.DocMsgNFTMint)
	nftResp, err := opb_client.QueryNFT(nftMint.Denom, nftMint.Id)
	if err != nil {
		return err
	}

	denom, err := opb_client.QueryDenom(nftMint.Denom)
	if err != nil {
		return err
	}

	one, err := n.repo.FindOne(nftMint.Denom, nftMint.Id)
	if err != nil {
		if err == qmgo.ErrNoSuchDocuments {
			return n.repo.Save(&dto.Nft{
				DenomID:         denom.Denom.Id,
				NftID:           nftResp.NFT.Id,
				CreateTime:      int(time.Now().Unix()),
				Data:            nftResp.NFT.Data,
				DenomName:       denom.Denom.Name,
				LastBlockHeight: int(tx.Height),
				LastBlockTime:   int(tx.Time),
				NftName:         nftResp.NFT.Name,
				Owner:           nftResp.NFT.Owner,
				UpdateTime:      int(time.Now().Unix()),
				URI:             nftResp.NFT.URI,
			})
		}
		return err
	}

	one.DenomID = denom.Denom.Id
	one.NftID = nftResp.NFT.Id
	one.Data = nftResp.NFT.Data
	one.DenomName = denom.Denom.Name
	one.LastBlockHeight = int(tx.Height)
	one.LastBlockTime = int(tx.Time)
	one.NftName = nftResp.NFT.Name
	one.Owner = nftResp.NFT.Owner
	one.UpdateTime = int(time.Now().Unix())
	one.URI = nftResp.NFT.URI
	return n.repo.Update(one)
}

type NftEdit struct {
	repo repository.INftRepo
}

func (n *NftEdit) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var nftEdit = msg.Msg.(*nft.DocMsgNFTEdit)
	nftResp, err := opb_client.QueryNFT(nftEdit.Denom, nftEdit.Id)
	if err != nil {
		return err
	}

	denom, err := opb_client.QueryDenom(nftEdit.Denom)
	if err != nil {
		return err
	}

	one, err := n.repo.FindOne(nftEdit.Denom, nftEdit.Id)
	if err != nil {
		if err == qmgo.ErrNoSuchDocuments {
			return n.repo.Save(&dto.Nft{
				DenomID:         denom.Denom.Id,
				NftID:           nftResp.NFT.Id,
				CreateTime:      int(time.Now().Unix()),
				Data:            nftResp.NFT.Data,
				DenomName:       denom.Denom.Name,
				LastBlockHeight: int(tx.Height),
				LastBlockTime:   int(tx.Time),
				NftName:         nftResp.NFT.Name,
				Owner:           nftResp.NFT.Owner,
				UpdateTime:      int(time.Now().Unix()),
				URI:             nftResp.NFT.URI,
			})
		}
		return err
	}

	one.DenomID = denom.Denom.Id
	one.NftID = nftResp.NFT.Id
	one.Data = nftResp.NFT.Data
	one.DenomName = denom.Denom.Name
	one.LastBlockHeight = int(tx.Height)
	one.LastBlockTime = int(tx.Time)
	one.NftName = nftResp.NFT.Name
	one.Owner = nftResp.NFT.Owner
	one.UpdateTime = int(time.Now().Unix())
	one.URI = nftResp.NFT.URI
	return n.repo.Update(one)
}

type NftTransfer struct {
	repo repository.INftRepo
}

func (n *NftTransfer) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var nftTransfer = msg.Msg.(*nft.DocMsgNFTTransfer)
	nftResp, err := opb_client.QueryNFT(nftTransfer.Denom, nftTransfer.Id)
	if err != nil {
		return err
	}

	denom, err := opb_client.QueryDenom(nftTransfer.Denom)
	if err != nil {
		return err
	}

	one, err := n.repo.FindOne(nftTransfer.Denom, nftTransfer.Id)
	if err != nil {
		return n.repo.Save(&dto.Nft{
			DenomID:         denom.Denom.Id,
			NftID:           nftResp.NFT.Id,
			CreateTime:      int(time.Now().Unix()),
			Data:            nftResp.NFT.Data,
			DenomName:       denom.Denom.Name,
			LastBlockHeight: int(tx.Height),
			LastBlockTime:   int(tx.Time),
			NftName:         nftResp.NFT.Name,
			Owner:           nftResp.NFT.Owner,
			UpdateTime:      int(time.Now().Unix()),
			URI:             nftResp.NFT.URI,
		})
	}

	one.DenomID = denom.Denom.Id
	one.NftID = nftResp.NFT.Id
	one.Data = nftResp.NFT.Data
	one.DenomName = denom.Denom.Name
	one.LastBlockHeight = int(tx.Height)
	one.LastBlockTime = int(tx.Time)
	one.NftName = nftResp.NFT.Name
	one.Owner = nftResp.NFT.Owner
	one.UpdateTime = int(time.Now().Unix())
	one.URI = nftResp.NFT.URI
	return n.repo.Update(one)
}

type NftBurn struct {
	repo repository.INftRepo
}

func (n *NftBurn) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var nftBurn = msg.Msg.(*nft.DocMsgNFTBurn)

	one, err := n.repo.FindOne(nftBurn.Denom, nftBurn.Id)
	if err != nil {
		return err
	}

	return n.repo.Delete(one)
}
