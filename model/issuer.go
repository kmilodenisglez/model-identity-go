package identity

type IssuerCreateRequest struct {
	Name      string `json:"name,omitempty" metadata:",optional"` // can be obtained directly from the certificate "CommonName"
	CertPem   string `json:"certPem"`                             // issuer PEM certificate in base64
	PublicKey string `json:"publicKey,omitempty" metadata:",optional"`
	ByDefault bool   `json:"byDefault,omitempty" metadata:",optional"`
}

type IssuerQueryResponse struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	PublicKey   string            `json:"publicKey"`
	Attrs       Attrs             `json:"attr,omitempty" metadata:",optional"`        // subject cert
	AttrsExtras map[string]string `json:"attrsExtras,omitempty" metadata:",optional"` // non-standard X.509 certificate extension asn1 1.2.3.4.5.6.7.8.1, etc.
	IssuedTime  string            `json:"issuedTime"`
	ExpiresTime string            `json:"expiresTime"`
	Active      bool              `json:"active"`
	ByDefault   bool              `json:"byDefault"`
}

type IssuerHistoryQueryResponse struct {
	Record   *IssuerQueryResponse `json:"record"`
	TxID     string               `json:"txID"`
	Time     string               `json:"time"`
	IsDelete bool                 `json:"isDelete"`
}

type IssuerRichQuerySelector struct {
	Selector struct {
		DocType  string   `json:"docType"`
		ID       string   `json:"id,omitempty" metadata:",optional"` // participant id: used in the composite key to store the participant in the ledger
		Name     string   `json:"name,omitempty" metadata:",optional"`
		Active   bool     `json:"active,omitempty" metadata:",optional"`
		UseIndex []string `json:"use_index,omitempty" metadata:",optional"`
	} `json:"selector"`
	PageSize int    `json:"pageSize,omitempty" metadata:",optional"`
	Bookmark string `json:"bookmark,omitempty" metadata:",optional"`
}
