package godap_user

import (
	b64 "encoding/base64"
)

// standard b64 encoding to not confuse encodings
func Encode(b []byte) string {
	return b64.StdEncoding.EncodeToString(b)
}

// standard b64 decoding to not confuse encodings
func Decode(s string) []byte {
	ret, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return ret
}
