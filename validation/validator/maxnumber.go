package validator

import "strconv"

// MaxNumber validator
type MaxNumber struct {
	Max int
}

// Check a param less than max number
func (m MaxNumber) Check(param string) (int, bool) {

	if isBlank(param) {
		return 0, true
	}

	v, err := strconv.Atoi(param)
	return v, err == nil && !(v > m.Max)
}

// Validate for Validator interface
func (m MaxNumber) Validate(param string) bool {
	_, ok := m.Check(param)
	return ok
}
