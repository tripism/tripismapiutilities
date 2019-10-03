package tripismapiutilities

import (
	uuid "github.com/satori/go.uuid"
)

// GetUUID generates a UUID/GUID (version 1 format) using the satori/go.uuid package
func GetUUID() string {
	// Version 1, based on timestamp and MAC address (RFC 4122)
	u1 := uuid.NewV1()
	return u1.String()
}

// GetUUIDv4 generates a UUID/GUID (version 4 format) using the satori/go.uuid package
func GetUUIDv4() string {
	// Version 4, based on random numbers (RFC 4122)
	u4 := uuid.NewV4()
	return u4.String()
}
