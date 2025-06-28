package digiflazz

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// VerifyWebhookSignature memverifikasi signature dari request webhook yang masuk.
// Gunakan fungsi ini di dalam HTTP handler Anda untuk memastikan request berasal dari Digiflazz.
//
// secret: Secret key yang Anda atur di dashboard Digiflazz.
// body: Raw body dari request HTTP POST.
// signatureHeader: Nilai dari header "X-Hub-Signature".
func VerifyWebhookSignature(secret string, body []byte, signatureHeader string) error {
	if signatureHeader == "" {
		return errors.New("missing X-Hub-Signature header")
	}

	parts := strings.SplitN(signatureHeader, "=", 2)
	if len(parts) != 2 || parts[0] != "sha1" {
		return fmt.Errorf("invalid signature format: %s", signatureHeader)
	}

	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write(body)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(parts[1]), []byte(expectedMAC)) {
		return errors.New("signature mismatch")
	}

	return nil
}