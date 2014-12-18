package validator

import "strconv"

// MinNumber validator
type MinNumber struct {
	Min int
}

// Check a param greater than min number
func (m MinNumber) Check(param string) (int, bool) {

	if isBlank(param) {
		return 0, true
	}

	v, err := strconv.Atoi(param)
	return v, err == nil && !(v < m.Min)
}

// Validate for Validator interface
func (m MinNumber) Validate(param string) bool {
	_, ok := m.Check(param)
	return ok
}
