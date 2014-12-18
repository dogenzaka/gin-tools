package validator

// FixedLength validator
type FixedLength struct {
	Fixed int
}

// Check a param equals fixed length
func (f FixedLength) Check(param string) (string, bool) {

	if isBlank(param) {
		return param, true
	}

	return param, f.Fixed == len(param)
}

// Validate for Validator interface
func (f FixedLength) Validate(param string) bool {
	_, ok := f.Check(param)
	return ok
}
