package tripismapiutilities

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
)

// EncodeString encodes a string into Base64
func EncodeString(i map[string]interface{}) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(dst, []byte(b))
	return string(dst), nil
}

// DecodeString decodes a Base64 encoded string
func DecodeString(s string) (map[string]interface{}, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
	n, err := base64.StdEncoding.Decode(dst, []byte(s))
	if err != nil {
		return nil, err
	}
	var v map[string]interface{}
	if err := Unmarshal(bytes.NewReader(dst[:n]), &v); err != nil {
		return nil, err
	}
	return v, nil
}
