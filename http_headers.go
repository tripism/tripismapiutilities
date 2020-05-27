package tripismapiutilities

import (
	"errors"
	"net/http"
	"strings"
)

// Supported Custom HTTP Headers for the server
const (
	HTTPHeaderKey                = "X-Tripism-Key"
	HTTPHeaderAuth               = "X-Tripism-Auth"
	HTTPHeaderUserID             = "X-Tripism-UserID"
	HTTPHeaderDeviceChannel      = "X-Tripism-Channel"
	HTTPHeaderDeviceUserAgent    = "X-Tripism-User-Agent"
	HTTPHeaderSource             = "X-Tripism-Source"
	HTTPHeaderResultsCount       = "X-Tripism-Results-Count"
	HTTPHeaderBoardContentsCount = "X-Tripism-Board-Contents-Count"
)

// Supported Headers for Device Channel
const (
	HTTPHeaderDeviceChannelDesktop = "DESKTOP"
	HTTPHeaderDeviceChannelMobile  = "MOBILE"
)

// HTTPHeaderGetDeviceChannel validates and returns the value of the header X-Tripism-Channel from an HTTP Request
func HTTPHeaderGetDeviceChannel(r *http.Request) (string, error) {
	if r == nil {
		return "", errors.New("missing r parameter")
	}

	header := strings.ToUpper(strings.TrimSpace(r.Header.Get(HTTPHeaderDeviceChannel)))
	if header == "" {
		return HTTPHeaderDeviceChannelDesktop, nil
	}

	if !(header == HTTPHeaderDeviceChannelDesktop || header == HTTPHeaderDeviceChannelMobile) {
		return "", errors.New("invalid header parameter " + header)
	}

	return header, nil
}
