package repository

import (
	"context"
	"github.com/llygcd/block-compensation/internal/model/dto"
	"github.com/qiniu/qmgo"
)

type ITxRepo interface {
	Save(tx []*dto.Tx) error
}

func NewTxRepo(cli *qmgo.Client, database string) *TxRepo {
	return &TxRepo{coll: cli.Database(database).Collection(dto.Tx{}.CollectionName())}
}

type TxRepo struct {
	coll *qmgo.Collection
}

func (repo *TxRepo) Save(tx []*dto.Tx) error {
	_, err := repo.coll.InsertMany(context.Background(), tx)
	return err
}
