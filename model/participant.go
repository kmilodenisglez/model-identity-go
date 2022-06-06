package identity

// ***********  SCHEMES  ***********

type ParticipantGetRequest struct {
	Did string `json:"did"`
}

type ParticipantCreateRequest struct {
	DID        string   `json:"did,omitempty" metadata:",optional"`
	PublicKey  string   `json:"publicKey"` // required
	IssuerID   string   `json:"issuerID,omitempty" metadata:",optional"`
	CreatorDid string   `json:"creatorDid,omitempty" metadata:",optional"` // creator did, "empty if is SelfParticipant"
	CertPem    string   `json:"certPem,omitempty" metadata:",optional"`    // PEM certificate in base64, used to valid publicKey and to store participant attributes
	Roles      []string `json:"roles,omitempty" metadata:",optional"`      // role id list
}

type ParticipantUpdateRequest struct {
	DID         string                 `json:"did,omitempty" metadata:",optional"`
	PublicKey   string                 `json:"publicKey,omitempty" metadata:",optional"`
	IssuerID    string                 `json:"issuerID,omitempty" metadata:",optional"`
	CertPem     string                 `json:"certPem,omitempty" metadata:",optional"` // PEM certificate in base64, used only to store participant attributes
	Roles       []string               `json:"roles,omitempty" metadata:",optional"`   // role id list
	Active      bool                   `json:"active"`
	AttrsExtras map[string]interface{} `json:"attrsExtras,omitempty" metadata:",optional"`
}

type ParticipantDeleteRequest struct {
	UserDid   string `json:"userDid"`
	CallerDid string `json:"callerDid"` // caller DID
}

type ParticipantResponse struct {
	DID     string                    `json:"did"`
	Roles   []string                  `json:"roles,omitempty" metadata:",optional"`   // role id
	Creator *ParticipantCreateRequest `json:"creator,omitempty" metadata:",optional"` // creator, "empty if is SelfParticipant"
}

type ParticipantDeletedPayload struct {
	MspID    string `json:"mspID"`    // mspID of the invoking client
	Time     string `json:"time"`     // exact time when the operation is confirmed
	CallerID string `json:"callerID"` // participant ID that invokes the tx
}

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

type ParticipantRichQuerySelector struct {
	Selector struct {
		DocType   string   `json:"docType"`
		DID       string   `json:"did,omitempty" metadata:",optional"` // participant id
		PublicKey string   `json:"publicKey,omitempty" metadata:",optional"`
		Active    bool     `json:"active" metadata:",optional"`
		UseIndex  []string `json:"use_index,omitempty" metadata:",optional"`
	} `json:"selector"`
	PageSize int    `json:"pageSize,omitempty" metadata:",optional"`
	Bookmark string `json:"bookmark,omitempty" metadata:",optional"`
}