package main

import (
	"fmt"
	"strings"

	"code.google.com/p/go-uuid/uuid"
)

// ValidateUri Will attempt to validate URI. Will return error if URI is not valid.
func ValidateUri(uri string) error {

	if len(uri) < 3 || len(uri) > 20 {
		return fmt.Errorf("Could not validate (uri: %s) as length is not matched. URI must be between 3 and 20 characters long.", uri)
	}

	return nil
}

// GenerateUri Will return UUID without dashes as that's pretty much unique as it can go
func GenerateUri() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}
