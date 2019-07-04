package encrypt

import(
	"crypto/hmac"
	"crypto/sha256"
)

var key = []byte("xONuoQUBjKwnxmTm6YXQDRGDsL2fR8mq3iXQlP02FxkdwIFw47V-mv2ZrHPIqcr7")
var mac = hmac.New(sha256.New, key)

func SHA256(source string) []byte{
	mac.Write([]byte(source))
	return mac.Sum(nil)
}
