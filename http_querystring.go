package tripismapiutilities

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// HTTPGetQueryStringParameterCase is a parameter to convert parameter values to UPPER, lower or None
type HTTPGetQueryStringParameterCase string

// Determine if the parameter value should be converted to upper, lower or left as input
const (
	ParameterCaseUpper HTTPGetQueryStringParameterCase = "upper"
	ParameterCaseLower HTTPGetQueryStringParameterCase = "lower"
	ParameterCaseNone  HTTPGetQueryStringParameterCase = "none"
)

// HTTPGetQueryStringParameter validates and returns a query string parameter
func HTTPGetQueryStringParameter(r *http.Request, key string, required bool, requiredCase HTTPGetQueryStringParameterCase) (string, error) {
	if r == nil {
		return "", errors.New("missing r parameter")
	}
	if key == "" {
		return "", errors.New("missing key parameter")
	}

	q := r.URL.Query()

	qkey := strings.TrimSpace(q.Get(key))

	query, err := url.QueryUnescape(qkey)
	if err != nil {
		return "", errors.New("failed to unescape query parameter " + key + ": " + err.Error())
	}

	if len(query) == 0 && required {
		return "", errors.New("missing required parameter " + key)
	}

	switch requiredCase {
	case ParameterCaseUpper:
		return strings.ToUpper(query), nil
	case ParameterCaseLower:
		return strings.ToLower(query), nil
	}

	return query, nil
}
