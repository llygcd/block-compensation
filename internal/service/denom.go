package service

/*type DenomIssue struct {
	repo repository.IDenomRepo
}

func (n *DenomIssue) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var issueDenom = msg.Msg.(*nft.DocMsgIssueDenom)
	denom, err := opb.QueryDenom(issueDenom.Id)
	if err != nil {
		return errors.New(err.Error())
	}
	owner, err := opb.QueryOwner(denom.Creator, denom.Id)
	if err != nil {
		return errors.New(err.Error())
	}
	return n.repo.Save(&dto.Denom{
		Name:            denom.Name,
		DenomID:         denom.Id,
		JsonSchema:      denom.Schema,
		Creator:         denom.Creator,
		Owner:           owner.Address,
		Txhash:          tx.TxHash,
		Height:          tx.Height,
		Time:            tx.Time,
		CreateTime:      time.Now().Unix(),
		LastBlockHeight: tx.Height,
		LastBlockTime:   tx.Time,
	})
}

type DenomTransfer struct {
	repo repository.IDenomRepo
}

func (n *DenomTransfer) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var transferDenom = msg.Msg.(*nft.DocMsgTransferDenom)
	denom, err := opb.QueryDenom(transferDenom.Id)
	if err != nil {
		return errors.New(err.Error())
	}
	owner, err := opb.QueryOwner(denom.Creator, denom.Id)
	if err != nil {
		return errors.New(err.Error())
	}
	one, err1 := n.repo.FindOne(denom.Id)
	if err1 != nil {
		return errors.New(err1.Error())
	}

	one.Creator = denom.Creator
	one.Owner = owner.Address
	one.Txhash = tx.TxHash
	one.Height = tx.Height
	one.Time = tx.Time
	one.LastBlockHeight = tx.Height
	one.LastBlockTime = tx.Time
	return n.repo.Update(one)
}*/
