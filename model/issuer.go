package identity

// IssuerCreateRequest
type IssuerCreateRequest struct {
	Name      string `json:"name,omitempty" metadata:",optional"`    // can be obtained directly from the certificate "CommonName"
	CertPem   string `json:"certPem,omitempty" metadata:",optional"` // issuer PEM certificate in base64
	PublicKey string `json:"publicKey"`
	ByDefault bool   `json:"byDefault,omitempty" metadata:",optional"`
}

// IssuerQueryResponse
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