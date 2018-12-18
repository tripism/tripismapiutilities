package tripismapiutilities

import (
	"net/http"

	"github.com/baloo32/gomobiledetect"
)

// IsMobileDevice returns true if HTTP request is made by mobile device
func IsMobileDevice(r *http.Request) bool {
	device := mobiledetect.NewMobileDetect(r, nil)

	if device.IsMobile() && !device.IsTablet() {
		return true
	}
	return false
}
