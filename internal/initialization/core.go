package initialization

import (
	"github.com/llygcd/block-compensation/internal/api/rest"
	"github.com/llygcd/block-compensation/internal/model"
	"github.com/llygcd/block-compensation/internal/repository"
	"github.com/llygcd/block-compensation/internal/service"
	"github.com/llygcd/block-compensation/pkg/pool"
	"github.com/qiniu/qmgo"
)

func NewRepositories(mongoCli *qmgo.Client, dbName string) *model.Repositories {
	return &model.Repositories{
		DenomRepo: repository.NewDenomRepo(mongoCli, dbName),
		NftRepo:   repository.NewNftRepo(mongoCli, dbName),
		TxRepo:    repository.NewTxRepo(mongoCli, dbName),
		BolckRepo: repository.NewBlockRepo(mongoCli, dbName),
	}
}

func NewServices(r *model.Repositories, client *pool.Client) *model.Services {
	return &model.Services{
		CompensationService: service.NewCompensationService(r.DenomRepo, r.NftRepo, r.BolckRepo, r.TxRepo, client),
	}
}

func NewControllers(s *model.Services) *model.Controllers {
	return &model.Controllers{
		CompensationController: rest.NewCompensationController(s.CompensationService),
	}
}
