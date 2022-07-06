package service

import (
	"github.com/llygcd/block-compensation/handlers"
	"github.com/llygcd/block-compensation/internal/repository"
	"github.com/llygcd/block-compensation/pkg/pool"
	"github.com/sirupsen/logrus"
)

type CompensationService struct {
	denomRepo  repository.IDenomRepo
	nftRepo    repository.INftRepo
	blockRepo  repository.IBlockRepo
	txRepo     repository.ITxRepo
	poolClient *pool.Client
}

func NewCompensationService(denomRepo repository.IDenomRepo, nftRepo repository.INftRepo, blockRepo repository.IBlockRepo, txRepo repository.ITxRepo, poolClient *pool.Client) *CompensationService {
	return &CompensationService{
		denomRepo:  denomRepo,
		nftRepo:    nftRepo,
		blockRepo:  blockRepo,
		txRepo:     txRepo,
		poolClient: poolClient,
	}
}

func (s *CompensationService) Compensation(height int64) {
	txs, txes, err := handlers.ParseBlockAndTxs(height, s.poolClient)
	if err != nil {
		logrus.Error(txs, txes)
	}

	if err = s.blockRepo.Save(txs); err != nil {
		logrus.Error(err)
	}

	if err = s.txRepo.Save(txes); err != nil {
		logrus.Error(err)
	}

	if len(txes) > 0 {
		for _, tx := range txes {
			for _, msg := range tx.DocTxMsgs {
				txTypeStrategy := GetTxTypeStrategy(msg.Type, s)
				txTypeStrategy.Compensation(msg, tx)
			}
		}
	}

}
