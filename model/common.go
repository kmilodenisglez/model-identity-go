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

type ParticipantRichQuerySelector struct {
	Selector struct {
		DocType   string `json:"docType"`
		ID        string `json:"id,omitempty" metadata:",optional"`  // participant id: used in the composite key to store the participant in the ledger
		Did       string `json:"did,omitempty" metadata:",optional"` // "did:hash_public_key"
		PublicKey string `json:"publicKey,omitempty" metadata:",optional"`
		Active    bool   `json:"active,omitempty" metadata:",optional"`
	} `json:"selector"`
}

type IssuerRichQuerySelector struct {
	Selector struct {
		DocType string `json:"docType"`
		ID      string `json:"id,omitempty" metadata:",optional"` // participant id: used in the composite key to store the participant in the ledger
		Name    string `json:"name,omitempty" metadata:",optional"`
		Active  bool   `json:"active,omitempty" metadata:",optional"`
	} `json:"selector"`
}

// PaginatedQueryResponse structure used for returning paginated query results and metadata
type PaginatedQueryResponse struct {
	Records             []*interface{} `json:"records"`
	FetchedRecordsCount int32          `json:"fetchedRecordsCount"`
	Bookmark            string         `json:"bookmark"`
}