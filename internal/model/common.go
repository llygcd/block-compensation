package model

import (
	"github.com/llygcd/block-compensation/internal/api/rest"
	"github.com/llygcd/block-compensation/internal/repository"
	"github.com/llygcd/block-compensation/internal/service"
	"github.com/qiniu/qmgo/options"
)

type Docs interface {
	// CollectionName collection name
	CollectionName() string
	// Indexes ensure indexes
	Indexes() []options.IndexModel
	// PkKvPair primary key pair(used to find a unique record)
	PkKvPair() map[string]interface{}
}

type Repositories struct {
	DenomRepo *repository.DenomRepo
	NftRepo   *repository.NftRepo
	TxRepo    *repository.TxRepo
	BolckRepo *repository.BlockRepo
}

type Services struct {
	CompensationService *service.CompensationService
}

type Controllers struct {
	CompensationController *rest.CompensationController
}
