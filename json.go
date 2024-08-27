package tripismapiutilities

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Unmarshal unmarshalles a Reader Object and returns as a mapped string interface the JSON document
// See https://ahmetalpbalkan.com/blog/golang-json-decoder-pitfalls/ for why to change from decode.
func Unmarshal(r io.Reader, obj interface{}) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}

// Marshal is a helper function to marshal an object to JSON byte array
func Marshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

// TestJSON POSTs a marshalled JSON byte array to httpbin.org/post and returns the parsed response
func TestJSON(buf []byte) (string, error) {
	testJSONUrl := "https://httpbin.org/post"

	res, err := http.Post(testJSONUrl, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// DecodeAnyJSON reads the body of an http.Response or http.Request Body, decodes any unknown JSON and returns as an interface
func DecodeAnyJSON(b io.ReadCloser) (interface{}, error) {
	var j interface{}
	err := json.NewDecoder(b).Decode(&j)
	if err != nil {
		return nil, err
	}
	return j, nil
}
