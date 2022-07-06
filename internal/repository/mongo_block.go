package repository

import (
	"context"
	"github.com/llygcd/block-compensation/internal/model/dto"
	"github.com/qiniu/qmgo"
)

type IBlockRepo interface {
	Save(tx *dto.Block) error
}

func NewBlockRepo(cli *qmgo.Client, database string) *BlockRepo {
	return &BlockRepo{coll: cli.Database(database).Collection(dto.Block{}.CollectionName())}
}

type BlockRepo struct {
	coll *qmgo.Collection
}

func (repo *BlockRepo) Save(tx *dto.Block) error {
	_, err := repo.coll.InsertOne(context.Background(), tx)
	return err
}
