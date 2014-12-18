package validator

import "encoding/hex"

// ObjectID validator
type ObjectID struct{}

// Check a param should be 24-bits HEX
func (o ObjectID) Check(param string) (string, bool) {

	if isBlank(param) {
		return param, true
	}

	if len(param) != 24 {
		return param, false
	}

	_, err := hex.DecodeString(param)

	return param, err == nil
}

// Validate for Validator interface
func (o ObjectID) Validate(param string) bool {
	_, ok := o.Check(param)
	return ok
}
