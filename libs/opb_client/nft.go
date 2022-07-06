package opb

import (
	"context"
)

func QueryOwner(creator, denom string, pagination *PageRequest) (QueryOwnerResponse, error) {
	if len(denom) == 0 {
		return QueryOwnerResp{}, sdk.Wrapf("denom is required")
	}

	conn, err := nc.GenConn()

	if err != nil {
		return QueryOwnerResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Owner(
		context.Background(),
		&QueryOwnerRequest{
			Owner:      creator,
			DenomId:    denom,
			Pagination: tidyPagination(pagination),
		},
	)
	if err != nil {
		return QueryOwnerResp{}, sdk.Wrap(err)
	}

	return res.Owner.Convert().(QueryOwnerResp), nil
}

/*func QueryDenom(denom string) (QueryDenomResp, sdk.Error) {
	conn, err := nc.GenConn()

	if err != nil {
		return QueryDenomResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Denom(
		context.Background(),
		&QueryDenomRequest{DenomId: denom},
	)
	if err != nil {
		return QueryDenomResp{}, sdk.Wrap(err)
	}

	return res.Denom.Convert().(QueryDenomResp), nil
}

func QueryNFT(denom, tokenID string) (QueryNFTResp, sdk.Error) {
	if len(denom) == 0 {
		return QueryNFTResp{}, sdk.Wrapf("denom is required")
	}

	if len(tokenID) == 0 {
		return QueryNFTResp{}, sdk.Wrapf("tokenID is required")
	}

	conn, err := nc.GenConn()

	if err != nil {
		return QueryNFTResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).NFT(
		context.Background(),
		&QueryNFTRequest{
			DenomId: denom,
			TokenId: tokenID,
		},
	)
	if err != nil {
		return QueryNFTResp{}, sdk.Wrap(err)
	}

	return res.NFT.Convert().(QueryNFTResp), nil
}*/

func tidyPagination(pagination *PageRequest) *PageRequest {
	pagination.CountTotal = false
	if pagination.Limit == 0 || pagination.Limit > 100 {
		pagination.Limit = 100
	}
	return pagination
}
