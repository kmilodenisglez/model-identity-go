package identity

type AccessResponse struct {
	DocType           string   `json:"docType"`
	ID                string   `json:"id"`
	Description       string   `json:"description,omitempty" metadata:",optional"` // access description
	ContractFunctions []string `json:"contractFunctions,omitempty"`                // contract functions
}
