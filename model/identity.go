package identity

// ***********  SCHEMES  ***********

// ParticipantGetRequest
type ParticipantGetRequest struct {
	Did string `json:"did"`
}

// ParticipantCreateRequest
type ParticipantCreateRequest struct {
	Did        string   `json:"did"`
	PublicKey  string   `json:"publicKey"`
	IssuerID   string   `json:"issuerID,omitempty" metadata:",optional"`
	CreatorDid string   `json:"creatorDid,omitempty" metadata:",optional"` // creator did, "empty if is SelfParticipant"
	CertPem    string   `json:"certPem,omitempty" metadata:",optional"`    // PEM certificate in base64
	Roles      []string `json:"roles,omitempty" metadata:",optional"`      // role id list
}

// ParticipantDeleteRequest
type ParticipantDeleteRequest struct {
	UserDid   string `json:"userDid"`
	CallerDid string `json:"callerDid"` // caller DID
}

// ParticipantResponse
type ParticipantResponse struct {
	Did     string                    `json:"did"`
	Roles   []string                  `json:"roles,omitempty" metadata:",optional"`   // role id
	Creator *ParticipantCreateRequest `json:"creator,omitempty" metadata:",optional"` // creator, "empty if is SelfParticipant"
}

// ParticipantDeletedPayload
type ParticipantDeletedPayload struct {
	MspID    string `json:"mspID"`    // mspID of the invoking client
	Time     string `json:"time"`     // exact time when the operation is confirmed
	CallerID string `json:"callerID"` // participant ID that invokes the tx
}

// ParticipantQueryResponse
type ParticipantQueryResponse struct {
	ParticipantID string   `json:"participantID"`
	Did           string   `json:"did"`
	PublicKey     string   `json:"publicKey"`
	Roles         []string `json:"roles,omitempty"` // role id list
}

// ParticipantHistoryQueryResponse structure used for returning result of history query
type ParticipantHistoryQueryResponse struct {
	TxID     string                    `json:"txID"`
	Record   *ParticipantQueryResponse `json:"record"`
	Time     string                    `json:"time"`
	IsDelete bool                      `json:"isDelete"`
}

// PaginatedQueryResponse structure used for returning paginated query results and metadata
type PaginatedQueryResponse struct {
	Records             []*ParticipantResponse `json:"records"`
	FetchedRecordsCount int32                  `json:"fetchedRecordsCount"`
	Bookmark            string                 `json:"bookmark"`
}
