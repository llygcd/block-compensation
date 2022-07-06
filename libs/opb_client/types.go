package opb

func (o Owner) Convert() interface{} {
	var idcs []IDC
	for _, idc := range o.IDCollections {
		idcs = append(idcs, IDC{
			Denom:    idc.DenomId,
			TokenIDs: idc.TokenIds,
		})
	}
	return QueryOwnerResp{
		Address: o.Address,
		IDCs:    idcs,
	}
}

func (this BaseNFT) Convert() interface{} {
	return QueryNFTResp{
		ID:      this.Id,
		Name:    this.Name,
		URI:     this.URI,
		Data:    this.Data,
		Creator: this.Owner,
	}
}

type NFTs []BaseNFT

func (this Denom) Convert() interface{} {
	return QueryDenomResp{
		ID:      this.Id,
		Name:    this.Name,
		Schema:  this.Schema,
		Creator: this.Creator,
	}
}

type denoms []Denom

func (this denoms) Convert() interface{} {
	var denoms []QueryDenomResp
	for _, denom := range this {
		denoms = append(denoms, denom.Convert().(QueryDenomResp))
	}
	return denoms
}
