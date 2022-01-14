package identity

type AccessResponse struct {
	DocType           string   `json:"docType"`
	ID                string   `json:"id"`
	Description       string   `json:"description,omitempty" metadata:",optional"` // access description
	ContractFunctions []string `json:"contractFunctions,omitempty"`                // contract functions
}

// AccessCreateRequest
type AccessCreateRequest struct {
	ContractName      string   `json:"contractName"`
	ContractFunctions []string `json:"contractFunctions,omitempty"` // contract functions
}

// AccessUpdateRequest
type AccessUpdateRequest struct {
	ID                string   `json:"id"`
	ContractFunctions []string `json:"contractFunctions"` // contract functions name
}