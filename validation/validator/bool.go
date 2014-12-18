package validator

import "strconv"

// Bool validator is check a param is bool variable
type Bool struct{}

// Validate for Validator interface
func (b Bool) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	_, err := strconv.ParseBool(param)

	return err == nil
}
