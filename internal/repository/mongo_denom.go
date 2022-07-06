package repository

import (
	"context"
	"github.com/llygcd/block-compensation/internal/model/dto"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type IDenomRepo interface {
	Save(r *dto.Denom) error
	SaveAll(r []dto.Denom) error
	FindOne(denomId string) (dto.Denom, error)
	Update(r dto.Denom) error
}
type DenomRepo struct {
	coll *qmgo.Collection
}

func NewDenomRepo(cli *qmgo.Client, database string) *DenomRepo {
	return &DenomRepo{coll: cli.Database(database).Collection(dto.Denom{}.CollectionName())}
}

func (repo *DenomRepo) Save(r *dto.Denom) error {
	_, err := repo.coll.InsertOne(context.Background(), r)
	return err
}

func (repo *DenomRepo) SaveAll(r []dto.Denom) error {
	_, err := repo.coll.InsertMany(context.Background(), r)
	return err
}

func (repo *DenomRepo) Update(r dto.Denom) error {
	return repo.coll.UpdateOne(context.Background(), bson.M{"denom_id": r.DenomID},
		bson.M{"$set": bson.M{
			"creator": r.Creator,
			"owner":   r.Owner,
		}})
}

func (repo *DenomRepo) FindOne(denomId string) (dto.Denom, error) {
	one := dto.Denom{}
	err := repo.coll.Find(context.Background(), bson.M{"denom_id": denomId}).One(&one)
	return one, err
}
