package service

/*type NftMint struct {
	repo repository.INftRepo
}

func (n *NftMint) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var nftMint = msg.Msg.(*nft.DocMsgNFTMint)
	nftResp, err := opb.QueryNFT(nftMint.Denom, nftMint.Id)
	if err != nil {
		return errors.New(err.Error())
	}

	denom, err := opb.QueryDenom(nftMint.Denom)
	if err != nil {
		return errors.New(err.Error())
	}

	_, err2 := n.repo.FindOne(nftMint.Denom, nftMint.Id)
	if err2 != nil {
		if err2 == qmgo.ErrNoSuchDocuments {
			return n.repo.Save(&dto.Nft{
				DenomID:         nftMint.Denom,
				NftID:           nftResp.Id,
				CreateTime:      time.Now().Unix(),
				Data:            nftResp.Data,
				DenomName:       denom.Name,
				LastBlockHeight: tx.Height,
				LastBlockTime:   tx.Time,
				NftName:         nftResp.Name,
				Owner:           nftMint.Recipient,
				UpdateTime:      time.Now().Unix(),
				URI:             nftResp.URI,
				IsDelete:        false,
			})
		}
		return errors.New(err.Error())
	}

	return n.repo.Save(&dto.Nft{
		DenomID:         nftMint.Denom,
		NftID:           nftResp.Id,
		CreateTime:      time.Now().Unix(),
		Data:            nftResp.Data,
		DenomName:       denom.Name,
		LastBlockHeight: tx.Height,
		LastBlockTime:   tx.Time,
		NftName:         nftResp.Name,
		Owner:           nftMint.Recipient,
		UpdateTime:      time.Now().Unix(),
		URI:             nftResp.URI,
		IsDelete:        false,
	})
}

type NftEdit struct {
	repo repository.INftRepo
}

func (n *NftEdit) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var nftEdit = msg.Msg.(*nft.DocMsgNFTEdit)
	nftResp, err := opb.QueryNFT(nftEdit.Denom, nftEdit.Id)
	if err != nil {
		return errors.New(err.Error())
	}

	one, err2 := n.repo.FindOne(nftEdit.Denom, nftEdit.Id)
	if err2 != nil {
		if err2 == qmgo.ErrNoSuchDocuments {
			denom, err := opb.QueryDenom(nftEdit.Denom)
			if err != nil {
				return errors.New(err.Error())
			}
			return n.repo.Save(&dto.Nft{
				DenomID:         nftEdit.Denom,
				NftID:           nftResp.Id,
				CreateTime:      time.Now().Unix(),
				Data:            nftResp.Data,
				DenomName:       denom.Name,
				LastBlockHeight: tx.Height,
				LastBlockTime:   tx.Time,
				NftName:         nftResp.Name,
				Owner:           nftEdit.Sender,
				UpdateTime:      time.Now().Unix(),
				URI:             nftResp.URI,
			})
		}
		return errors.New(err.Error())
	}

	one.Data = nftResp.Data
	one.LastBlockHeight = tx.Height
	one.LastBlockTime = tx.Time
	one.NftName = nftResp.Name
	one.URI = nftResp.URI
	one.UpdateTime = time.Now().Unix()
	return n.repo.Update(one)
}

type NftTransfer struct {
	repo repository.INftRepo
}

func (n *NftTransfer) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var nftTransfer = msg.Msg.(*nft.DocMsgNFTTransfer)
	nftResp, err := opb.QueryNFT(nftTransfer.Denom, nftTransfer.Id)
	if err != nil {
		return errors.New(err.Error())
	}

	one, err2 := n.repo.FindOne(nftTransfer.Denom, nftTransfer.Id)
	if err2 != nil {
		denom, err := opb.QueryDenom(nftTransfer.Denom)
		if err != nil {
			return errors.New(err.Error())
		}
		return n.repo.Save(&dto.Nft{
			DenomID:         nftTransfer.Denom,
			NftID:           nftResp.Id,
			CreateTime:      time.Now().Unix(),
			Data:            nftResp.Data,
			DenomName:       denom.Name,
			LastBlockHeight: tx.Height,
			LastBlockTime:   tx.Time,
			NftName:         nftResp.Name,
			Owner:           nftTransfer.Recipient,
			UpdateTime:      time.Now().Unix(),
			URI:             nftResp.URI,
		})
	}

	one.Data = nftResp.Data
	one.LastBlockHeight = tx.Height
	one.LastBlockTime = tx.Time
	one.NftName = nftResp.Name
	one.URI = nftResp.URI
	one.UpdateTime = time.Now().Unix()
	return n.repo.Update(one)
}

type NftBurn struct {
	repo repository.INftRepo
}

func (n *NftBurn) Compensation(msg types.TxMsg, tx *dto.Tx) error {
	var nftBurn = msg.Msg.(*nft.DocMsgNFTBurn)

	one, err := n.repo.FindOne(nftBurn.Denom, nftBurn.Id)
	if err != nil {
		return err
	}

	return n.repo.Delete(one)
}*/
