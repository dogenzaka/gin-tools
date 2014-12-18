package validator

import "strconv"

// MinNumber validator is check a param greater than min number
type MinNumber struct {
	Min int
}

// Validate for Validator interface
func (m MinNumber) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	v, err := strconv.Atoi(param)
	return err == nil && !(v < m.Min)
}
