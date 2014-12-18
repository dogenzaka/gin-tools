package validator

import "strconv"

// MaxNumber validator is check a param less than max number
type MaxNumber struct {
	Max int
}

// Validate for Validator interface
func (m MaxNumber) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	v, err := strconv.Atoi(param)
	return err == nil && !(v > m.Max)
}
