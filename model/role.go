package identity

// RoleCreateRequest
type RoleCreateRequest struct {
	Name              string   `json:"name"`
	ContractFunctions []string `json:"contractFunctions,omitempty"` // contract functions
}

// RoleUpdateRequest
type RoleUpdateRequest struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	ContractFunctions []string `json:"contractFunctions"` // contract functions of Access
}

type RoleResponse struct {
	DocType           string   `json:"docType"`
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description,omitempty" metadata:",optional"` // description
	ContractFunctions []string `json:"contractFunctions,omitempty"`                // contract functions
}
