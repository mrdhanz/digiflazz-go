package digiflazz

import (
	"crypto/md5"
	"encoding/hex"
)

// generateSign membuat signature MD5 sesuai format Digiflazz.
func generateSign(username, apiKey, identifier string) string {
	data := username + apiKey + identifier
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}