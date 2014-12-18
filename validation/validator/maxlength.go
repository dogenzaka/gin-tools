package validator

// MaxLength validator is check a param is less than max length
type MaxLength struct {
	Max int
}

// Validate for Validator interface
func (m MaxLength) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	return !(len(param) > m.Max)
}
