package validator

// MaxLength validator
type MaxLength struct {
	Max int
}

// Check a param is less than max length
func (m MaxLength) Check(param string) (string, bool) {

	if isBlank(param) {
		return param, true
	}

	return param, !(len(param) > m.Max)
}

// Validate for Validator interface
func (m MaxLength) Validate(param string) bool {
	_, ok := m.Check(param)
	return ok
}
