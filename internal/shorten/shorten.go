package shorten

import (
	"crypto/sha256"
	"encoding/hex"
)

func Shorten(original string) string {
	checksum := sha256.Sum256([]byte(original))
	code := hex.EncodeToString(checksum[:])
	return code[:8]
}
