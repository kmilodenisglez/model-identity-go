package identity

// GetRequest Get request
type GetRequest struct {
	ID string `json:"id"`
}

// Transaction typical transaction structure
type Transaction struct {
	ID        string `json:"id"` // user id or did
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}

// TransactionD transaction definition with document field. Useful for Tx that involve
// document as part of the business process
type TransactionD struct {
	ID        string `json:"id"` // user id or did
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
	Document  string `json:"document"`
}

