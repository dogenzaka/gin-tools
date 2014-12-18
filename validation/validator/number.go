package validator

import "strconv"

// UInt validator
type UInt struct{}

// Check a param is unsigned integer
func (u UInt) Check(param string) (int, bool) {

	if isBlank(param) {
		return 0, true
	}

	v, err := strconv.Atoi(param)

	return v, err == nil && v > 0
}

// Validate for Validator interface
func (u UInt) Validate(param string) bool {
	_, ok := u.Check(param)
	return ok
}
