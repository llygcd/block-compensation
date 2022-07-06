package opb_client

import (
	"context"
	"errors"
	"github.com/llygcd/block-compensation/config"
	"github.com/llygcd/block-compensation/pkg/opb_client/nft"
	"google.golang.org/grpc"
)

var (
	Conn *grpc.ClientConn
)

func Init(conf *config.Config) *grpc.ClientConn {
	conn, err := grpc.Dial(conf.Server.GrpcUrls, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}

func QueryOwner(creator, denom string) (*nft.QueryOwnerResponse, error) {
	if len(denom) == 0 {
		return nil, errors.New("denom is required")
	}

	res, err := nft.NewQueryClient(Conn).Owner(
		context.Background(),
		&nft.QueryOwnerRequest{
			Owner:   creator,
			DenomId: denom,
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func QueryDenom(denom string) (*nft.QueryDenomResponse, error) {

	res, err := nft.NewQueryClient(Conn).Denom(
		context.Background(),
		&nft.QueryDenomRequest{DenomId: denom},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func QueryNFT(denom, tokenID string) (*nft.QueryNFTResponse, error) {
	if len(denom) == 0 {
		return nil, errors.New("denom is required")
	}

	if len(tokenID) == 0 {
		return nil, errors.New("tokenID is required")
	}

	res, err := nft.NewQueryClient(Conn).NFT(
		context.Background(),
		&nft.QueryNFTRequest{
			DenomId: denom,
			TokenId: tokenID,
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
