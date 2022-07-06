package repository

import (
	"context"
	"github.com/llygcd/block-compensation/internal/model/dto"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type INftRepo interface {
	Save(r *dto.Nft) error
	SaveAll(r []dto.Nft) error
	Update(one dto.Nft) error
	FindOne(denomId, nftId string) (dto.Nft, error)
	Delete(d dto.Nft) error
}

type NftRepo struct {
	coll *qmgo.Collection
}

func NewNftRepo(cli *qmgo.Client, database string) *NftRepo {
	return &NftRepo{coll: cli.Database(database).Collection(dto.Nft{}.CollectionName())}
}

func (repo *NftRepo) Save(r *dto.Nft) error {
	_, err := repo.coll.InsertOne(context.Background(), r)
	return err
}
func (repo *NftRepo) SaveAll(r []dto.Nft) error {
	_, err := repo.coll.InsertMany(context.Background(), r)
	return err
}

func (repo *NftRepo) Update(one dto.Nft) error {
	return repo.coll.UpdateOne(context.Background(), bson.M{"denom_id": one.DenomID, "nft_id": one.NftID},
		bson.M{"$set": bson.M{
			"data":              one.Data,
			"denom_name":        one.DenomName,
			"last_block_height": one.LastBlockHeight,
			"last_block_time":   one.LastBlockTime,
			"nft_name":          one.NftName,
			"owner":             one.Owner,
			"UpdateTime":        one.UpdateTime,
			"uri":               one.URI,
		}})
}

func (repo *NftRepo) FindOne(denomId, nftId string) (dto.Nft, error) {
	one := dto.Nft{}
	err := repo.coll.Find(context.Background(), bson.M{"denom_id": denomId, "nft_id": nftId}).One(&one)
	return one, err
}

func (repo *NftRepo) Delete(d dto.Nft) error {
	return repo.coll.Remove(context.Background(), bson.M{"denom_id": d.DenomID, "nft_id": d.NftID})
}
