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

type QueryPaginator struct {
	PageSize int    `json:"pageSize"`
	Bookmark string `json:"bookmark,omitempty" metadata:",optional"`
}

type RichQuerySelector struct {
	QueryString map[string]interface{} `json:"queryString"`
	PageSize    int                    `json:"pageSize,omitempty" metadata:",optional"`
	Bookmark    string                 `json:"bookmark,omitempty" metadata:",optional"`
}

// PaginatedQueryResponse structure used for returning paginated query results and metadata
type PaginatedQueryResponse struct {
	Records             []interface{} `json:"records"` // fabric-contract-api-go not support return []*interfaces{} type
	FetchedRecordsCount int32         `json:"fetchedRecordsCount"`
	Bookmark            string        `json:"bookmark"`
}