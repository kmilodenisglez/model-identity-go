package tools

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"hash"
	"regexp"
	"strings"
	"time"
)

func CheckDid(did string) error {
	// check the did format
	if matched, _ := regexp.MatchString(`^did:`, did); !matched {
		return fmt.Errorf("the did format not match")
	}
	return nil
}

func CreateDid(hash string) (string, error) {
	did := fmt.Sprintf("did:%s", hash)
	if err := CheckDid(did); err != nil {
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
