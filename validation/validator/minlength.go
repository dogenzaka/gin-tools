package validator

// MinLength validator
type MinLength struct {
	Min int
}

// Check a param greater than min length
func (m MinLength) Check(param string) (string, bool) {

	if isBlank(param) {
		return param, true
	}

	return param, !(len(param) < m.Min)
}

// Validate for validator interface
func (m MinLength) Validate(param string) bool {
	_, ok := m.Check(param)
	return ok
}
