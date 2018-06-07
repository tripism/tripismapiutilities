package tripismapiutilities

import (
	"net/http"
	"time"
)

// HTTPClient is a client with timeout set (default one doesn't have it and may hang forever)
var HTTPClient = &http.Client{Timeout: time.Duration(60) * time.Second}
