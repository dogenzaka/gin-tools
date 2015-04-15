package validator

import "strconv"

// Range validator
type Range struct {
	Min int
	Max int
}

// Check a param within a range
func (r Range) Check(param string) (int, bool) {

	if isBlank(param) {
		return 0, true
	}

	v, err := strconv.Atoi(param)
	return v, err == nil && !(v < r.Min) && !(v > r.Max)
}

// Validate for Validator interface
func (r Range) Validate(param string) bool {
	_, ok := r.Check(param)
	return ok
}
