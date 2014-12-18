package validator

import "encoding/hex"

// ObjectID validator is check a param should be 24-bits HEX
type ObjectID struct{}

// Validate for Validator interface
func (o ObjectID) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	if len(param) != 24 {
		return false
	}

	_, err := hex.DecodeString(param)

	return err == nil
}
