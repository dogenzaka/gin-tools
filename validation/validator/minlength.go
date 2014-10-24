package validator

// MinLength ... check min length
type MinLength struct {
	Min int
}

// Validate ... validate param of min length
func (m MinLength) Validate(param string) bool {
	if len(param) < m.Min {
		return false
	}
	return true
}
