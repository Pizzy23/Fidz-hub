package lInterface

type Contract struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type WalletResponse struct {
	ID        string `json:"id"`
	Address   string `json:"address"`
	ProjectID string `json:"projectId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type TokenTypeResponse struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	MaxSupply   int                    `json:"maxSupply"`
	Traits      map[string]interface{} `json:"traits"`
	ImageUrl    string                 `json:"imageUrl"`
}

type TransferDetails struct {
	ContractId string `json:"contractId"`
	From       string `json:"from"`
	To         string `json:"to"`
	TokenId    string `json:"tokenId"`
}

type TransferResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type GetContract struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	Address        string `json:"address"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Description    string `json:"description"`
	BlockchainName string `json:"blockchainName"`
	BlockscanUrl   string `json:"blockscanUrl"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	DeployedAt     string `json:"deployedAt"`
}

type GetContractsResponse struct {
	Contracts []GetContract `json:"contracts"`
}

type TokenType struct {
	ID            string            `json:"id"`
	UriNumber     int               `json:"uriNumber"`
	ContractID    string            `json:"contractId"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	MaxSupply     int               `json:"maxSupply"`
	CurrentSupply int               `json:"currentSupply"`
	Traits        map[string]string `json:"traits"`
	ImageUrl      string            `json:"imageUrl"`
	MetadataUrl   string            `json:"metadataUrl"`
	CreatedAt     string            `json:"createdAt"`
	UpdatedAt     string            `json:"updatedAt"`
}
