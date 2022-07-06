package opb

// IDC defines a set of nft ids that belong to a specific
type IDC struct {
	Denom    string   `json:"denom" yaml:"denom"`
	TokenIDs []string `json:"token_ids" yaml:"token_ids"`
}

type QueryOwnerResp struct {
	Address string `json:"address" yaml:"address"`
	IDCs    []IDC  `json:"idcs" yaml:"idcs"`
}

// BaseNFT non fungible token definition
type QueryNFTResp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
	Creator string `json:"creator"`
	URIHash string `json:"uri_hash"`
}

type QueryDenomResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Schema      string `json:"schema"`
	Creator     string `json:"creator"`
	Description string `json:"description"`
	Uri         string `json:"uri"`
	UriHash     string `json:"uri_hash"`
	Data        string `json:"data"`
}

type QueryCollectionResp struct {
	Denom QueryDenomResp `json:"denom" yaml:"denom"`
	NFTs  []QueryNFTResp `json:"nfts" yaml:"nfts"`
}
