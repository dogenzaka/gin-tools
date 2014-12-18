package validator

import "strconv"

// PNumber validator is check a param is positive number
type PNumber struct {
}

// Validate for Validator interface
func (p PNumber) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	v, err := strconv.Atoi(param)

	return err == nil && v > 0
}
