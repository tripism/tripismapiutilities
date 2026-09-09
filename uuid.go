package tripismapiutilities

import (
	uuid "github.com/google/uuid"
)

// GetUUID generates a UUID/GUID (version 1 format) using the google uuid package
func GetUUID() string {
	// Version 1, based on timestamp and MAC address (RFC 4122)
	u1, err := uuid.NewUUID()
	if err != nil {
		return uuid.New().String()
	}
	return u1.String()
}

// GetUUIDv4 generates a UUID/GUID (version 4 format) using the google uuid package
func GetUUIDv4() string {
	// Version 4, based on random numbers (RFC 4122)
	u4, err := uuid.NewRandom()
	if err != nil {
		return uuid.New().String()
	}
	return u4.String()
}

// GetUUIDv6 generates a UUID/GUID (version 6 format) using the google uuid package
func GetUUIDv6() string {
	// Version 6, based on timestamp and random numbers (RFC 4122)
	u6, err := uuid.NewV6()
	if err != nil {
		return uuid.New().String()
	}
	return u6.String()
}

// GetUUIDv7 generates a UUID/GUID (version 7 format) using the google uuid package
func GetUUIDv7() string {
	// Version 7, based on Unix timestamp and random numbers (RFC 4122)
	u7, err := uuid.NewV7()
	if err != nil {
		return uuid.New().String()
	}
	return u7.String()
}
