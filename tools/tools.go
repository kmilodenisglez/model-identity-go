package identity

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	identity "github.com/ic-matcom/model-identity-go/model"
	"hash"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func MatchDidFormat(did string) error {
	// check the did format
	if matched, _ := regexp.MatchString(`^did:`, did); !matched {
		return fmt.Errorf("the did format not match")
	}
	return nil
}

func CreateDid(publicKey string) (string, error) {
	checksum, _ := Checksum(HashSha256, []byte(publicKey))
	did := fmt.Sprintf("did:%s", checksum)
	if err := MatchDidFormat(did); err != nil {
		return "", fmt.Errorf("error generating the DID")
	}
	return did, nil
}

func GetTimestampRFC3339(timestamp *timestamp.Timestamp) string {
	tm := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
	return tm.Format(time.RFC3339)
}

func ParseRFC3339toTime(tm string) (time.Time, error) {
	t1, err := time.Parse(time.RFC3339, tm)
	if err != nil {
		return time.Time{}, err
	}
	return t1, nil
}

const (
	HashSha256 = "SHA256"
)

// Checksum returns the checksum of some data, using a specified algorithm.
// It only returns an error when an invalid algorithm is used. The valid ones
// are SHA256
func Checksum(algorithm string, data []byte) (checksum string, err error) {
	var _hash hash.Hash
	switch strings.ToUpper(algorithm) {
	case HashSha256:
		_hash = sha256.New()
	default:
		msg := "invalid algorithm parameter passed go Checksum: %s"
		return checksum, fmt.Errorf(msg, algorithm)
	}
	_hash.Write(data)
	str := hex.EncodeToString(_hash.Sum(nil))
	return str, nil
}

var (
	oidUserDNI = []int{0, 9, 2342, 19200300, 100, 1, 1}
)

// GetAttrsCert populates Attrs asset with multi-value RDNs:
//
//  - Name
//  - DNI
//  - Country
//  - Company
//  - OrganizationalUnit
//  - Locality
//  - Province
//
func GetAttrsCert(cert *x509.Certificate) identity.Attrs {
	attrs := identity.Attrs{
		Name:               cert.Subject.CommonName,
		Country:            GetFirstElem(cert.Subject.Country),
		Company:            GetFirstElem(cert.Subject.Organization),
		OrganizationalUnit: GetFirstElem(cert.Subject.OrganizationalUnit),
		Locality:           GetFirstElem(cert.Subject.Locality),
		Province:           GetFirstElem(cert.Subject.Province),
	}

	FillFromParsedCert(cert, &attrs)

	return attrs
}

// FillFromParsedCert to extract non-standard or others attributes
func FillFromParsedCert(cert *x509.Certificate, attrs *identity.Attrs) {
	// reading the OID from list of unparsed objects from Subject
	for _, n := range cert.Subject.Names {
		t := n.Type
		if len(t) == 4 && t[0] == 2 && t[1] == 5 && t[2] == 4 {
			switch t[3] {
			case 12:
				attrs.Position = n.Value.(string)
			}
		} else if n.Type.Equal(oidUserDNI) {
			if dni, ok := n.Value.(string); ok {
				attrs.DNI = dni
			}
		}
	}
}

func GetFirstElem(arr []string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

// GetTransactions returns callables functions of Contract
// used to populate ContractFunctions field in identity:Access
//
// Arguments:
//		0: contract ContractInterface
// Returns:
//		0: []string
func GetTransactions(contract interface{}) []string {
	avoidMap := map[string]string{"": "", "GetTransactionContextHandler": "", "InitLedger": "", "GetBeforeTransaction": "", "GetAfterTransaction": "", "GetIgnoredTransaction": "", "GetUnknownTransaction": ""}
	fooType := reflect.TypeOf(contract)
	var methods = make([]string, 0)
	for i := 0; i < fooType.NumMethod(); i++ {
		methodName := fooType.Method(i).Name
		_, exists := avoidMap[methodName]
		if !exists && !strings.Contains(methodName, "OnlyDev") {
			methods = append(methods, methodName)
		}
	}

	return UniqueStrings(methods)
}

// UniqueStrings returns a copy if the passed slice with only unique string results.
func UniqueStrings(input []string) []string {
	unique := make([]string, 0, len(input))
	m := make(map[string]bool, len(unique))

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			unique = append(unique, val)
		}
	}

	return unique
}