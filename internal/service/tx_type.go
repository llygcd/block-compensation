package service

import (
	"github.com/llygcd/block-compensation/internal/model/dto"
	//"github.com/bianjieai/opb-sdk-go/pkg/app/sdk/client"
	"github.com/kaifei-bianjie/msg-parser/types"
)

type TxTypeStrategy interface {
	Compensation(msg types.TxMsg, tx *dto.Tx) error
}

func GetTxTypeStrategy(txType string, serv *CompensationService) TxTypeStrategy {
	var txTypeStrategy TxTypeStrategy

	/*	switch txType {
		case TypeIssueDenom:
			txTypeStrategy = &DenomIssue{serv.denomRepo}
		case TypeTransferDenom:
			txTypeStrategy = &DenomTransfer{serv.denomRepo}
		case TypeNFTMint:
			txTypeStrategy = &NftMint{serv.nftRepo}
		case TypeNFTEdit:
			txTypeStrategy = &NftEdit{serv.nftRepo}
		case TypeNFTTransfer:
			txTypeStrategy = &NftTransfer{serv.nftRepo}
		case TypeNFTBurn:
			txTypeStrategy = &NftBurn{serv.nftRepo}
		}*/
	return txTypeStrategy
}

const (
	TypeIssueDenom    = "issue_denom"
	TypeTransferDenom = "transfer_denom"
	TypeNFTMint       = "mint_nft"
	TypeNFTEdit       = "edit_nft"
	TypeNFTTransfer   = "transfer_nft"
	TypeNFTBurn       = "burn_nft"
)
