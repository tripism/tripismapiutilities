package tripismapiutilities

import (
	"errors"
	"net/http"
	"strings"
)

type HTTPGetQueryStringParameterCase string

// Determine if the parameter value should be converted to upper, lower or left as input
const (
	ParameterCaseUpper HTTPGetQueryStringParameterCase = "upper"
	ParameterCaseLower HTTPGetQueryStringParameterCase = "lower"
	ParameterCaseNone  HTTPGetQueryStringParameterCase = "none"
)

// HTTPGetQueryStringParameter validates and returns a query string parameter
func HTTPGetQueryStringParameter(r *http.Request, p string, required bool, requiredCase HTTPGetQueryStringParameterCase) (string, error) {
	if r == nil {
		return "", errors.New("missing r parameter")
	}
	if p == "" {
		return "", errors.New("missing p parameter")
	}

	q := r.URL.Query()

	qp := strings.TrimSpace(q.Get(p))

	if len(qp) == 0 && required {
		return "", errors.New("missing required parameter " + p)
	}

	if requiredCase == ParameterCaseUpper {
		return strings.ToUpper(qp), nil
	} else if requiredCase == ParameterCaseLower {
		return strings.ToLower(qp), nil
	}

	return qp, nil
}
