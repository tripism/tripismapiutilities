package tripismapiutilities

import (
	"regexp"
)

// emailRegexp is the Regular Expression to check email format validity
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ValidateEmail performs a regular expression check against a given email address
func ValidateEmail(email string) bool {
	return emailRegexp.MatchString(email)
}
