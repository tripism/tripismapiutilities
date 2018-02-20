package APIUtils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
)

// encodeString encodes a string into Base64
func encodeString(i map[string]interface{}) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(dst, []byte(b))
	return string(dst), nil
}

// decodeString decodes a Base64 encoded string
func decodeString(s string) (map[string]interface{}, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
	n, err := base64.StdEncoding.Decode(dst, []byte(s))
	if err != nil {
		return nil, err
	}
	var v map[string]interface{}
	if err := unmarshal(bytes.NewReader(dst[:n]), &v); err != nil {
		return nil, err
	}
	return v, nil
}

// unmarshal unmarshalles a Reader Object and returns as a mapped string interface the JSON document
// See https://ahmetalpbalkan.com/blog/golang-json-decoder-pitfalls/ for why to change from decode.
func unmarshal(r io.Reader, obj interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}
