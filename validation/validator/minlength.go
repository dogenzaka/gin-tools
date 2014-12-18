package validator

// MinLength validator is check a param greater than min length
type MinLength struct {
	Min int
}

// Validate : param of min length
func (m MinLength) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	return !(len(param) < m.Min)
}
