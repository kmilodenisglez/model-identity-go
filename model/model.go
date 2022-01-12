package identity

import "fmt"

// Issuer are the companies issuer certs and attributes in the network
type Issuer struct {
	DocType     string            `json:"docType"` // docType is used to distinguish the various types of objects in state database
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	CertPem     string            `json:"certPem"`     // issuer PEM certificate in base64
	Attrs       Attrs             `json:"attrs"`       // subject cert
	AttrsExtras map[string]string `json:"attrsExtras"` // non-standard X.509 certificate extension asn1 1.2.3.4.5.6.7.8.1, etc.
	IssuedTime  string            `json:"issuedTime"`
	ExpiresTime string            `json:"expiresTime"`
	Active      bool              `json:"active"`
	ByDefault   bool              `json:"byDefault"`
}

// Access
type Access struct {
	DocType           string            `json:"docType"`
	ID                string            `json:"id"`                                         // contract name
	Description       string            `json:"description,omitempty" metadata:",optional"` // access description
	ContractFunctions map[string]string `json:"contractFunctions"`                          // contract functions name
}

// Role data
type Role struct {
	DocType           string            `json:"docType"`
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Description       string            `json:"description,omitempty" metadata:",optional"` // description
	ContractFunctions map[string]string `json:"contractFunctions,omitempty"`                // Access contract functions
}

// The Participant Identity asset
type Participant struct {
	DocType     string            `json:"docType"` // docType is used to distinguish the various types of objects in state database
	ID          string            `json:"id"`      // participant id: used in the composite key to store the participant in the ledger
	Did         string            `json:"did"`     // "did:hash_public_key"
	PublicKey   string            `json:"publicKey"`
	IssuerID    string            `json:"issuerId"`                               // issuer id
	Creator     string            `json:"creator,omitempty" metadata:",optional"` // Participant ID that created the identity, "empty if is SelfParticipant"
	Roles       []string          `json:"roles"`                                  // role id
	Attrs       Attrs             `json:"attrs"`                                  // Subject certificate
	AttrsExtras map[string]string `json:"attrsExtras"`                            // non-standard X.509 certificate extension asn1 1.2.3.4.5.6.7.8.1, etc.
	Time        string            `json:"time"`
	IssuedTime  string            `json:"issuedTime"`
	ExpiresTime string            `json:"expiresTime"`
	Active      bool              `json:"active"`
	MspID       string            `json:"mspID"` // MSP ID of the client
}

// attrContains returns true if the named attribute is found
func (p *Participant) attrContains(name string) bool {
	_, ok := p.AttrsExtras[name]
	return ok
}

// attrValue returns an attribute's value
func (p *Participant) attrValue(name string) (string, bool, error) {
	attr, ok := p.AttrsExtras[name]
	return attr, ok, nil
}

// attrValueTrue returns nil if the value of attribute 'name' is true;
// otherwise, an appropriate error is returned.
func (p *Participant) attrValueTrue(name string) error {
	val, ok, err := p.attrValue(name)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("attribute '%s' was not found", name)
	}
	if val != "true" {
		return fmt.Errorf("attribute '%s' is not true", name)
	}
	return nil
}

type Attrs struct {
	Name               string `json:"name"`
	DNI                string `json:"dni"`
	Company            string `json:"company"`
	Position           string `json:"position"`
	Country            string `json:"country"`
	Province           string `json:"province"`
	Locality           string `json:"locality"`
	OrganizationalUnit string `json:"organizationalUnit"`
}