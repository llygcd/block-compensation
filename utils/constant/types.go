package constant

const (
	TxStatusFail = iota
	TxStatusSuccess
	EnvNameConfigFilePath = "CONFIG_FILE_PATH"

	NoSupportMsgTypeTag = "no support msg parse"
	//unable to resolve type URL /cosmos.bank.v1beta1.MsgSend
	ErrNoSupportTxPrefix = "unable to resolve type URL"

	//cannot find transaction 601bf70ccdee4dde1c8be0d2_f018677a in queue for document {sync_task ObjectIdHex(\"601bdb0ccdee4dd7c214d167\")}
	ErrDbNotFindTransaction                = "cannot find transaction"
	IbcRecvPacketEventTypeWriteAcknowledge = "write_acknowledgement"
	IbcRecvPacketEventAttriKeyPacketAck    = "packet_ack"

	NFT_INFO_DO_NOT_MODIFY = "[do-not-modify]"
	ISSUE_DENOM            = "issue_denom"
	TRANSFER_DENOM         = "transfer_denom"
	BURN_NFT               = "burn_nft"
	TRANSFER_NFT           = "transfer_nft"
	EDIT_NFT               = "edit_nft"
	MINT_NFT               = "mint_nft"
)
