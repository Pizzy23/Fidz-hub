package inter

type TokenCreateInput struct {
	Traits      *map[string]string `json:"traits,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	MaxSupply   int                `json:"maxSupply"`
	ImageURL    string             `json:"imageUrl"`
}

type TokenCreateOutput struct {
	ID            string            `json:"id"`
	URINumber     int               `json:"uriNumber"`
	ContractID    string            `json:"contractId"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	MaxSupply     int               `json:"maxSupply"`
	CurrentSupply int               `json:"currentSupply"`
	Traits        map[string]string `json:"traits,omitempty"`
	ImageURL      string            `json:"imageUrl"`
	MetadataURL   string            `json:"metadataUrl"`
	CreatedAt     string            `json:"createdAt"`
	UpdatedAt     string            `json:"updatedAt"`
}
