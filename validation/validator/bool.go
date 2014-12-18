package validator

import "strconv"

// Bool validator
type Bool struct{}

// Check a param is bool variable
func (b Bool) Check(param string) (bool, bool) {

	if isBlank(param) {
		return false, true
	}

	v, err := strconv.ParseBool(param)

	return v, err == nil
}

// Validate for Validator interface
func (b Bool) Validate(param string) bool {
	_, ok := b.Check(param)
	return ok
}
