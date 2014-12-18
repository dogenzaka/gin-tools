package validator

// FixedLength validator is check a param equals fixed length
type FixedLength struct {
	Fixed int
}

// Validate for Validator interface
func (f FixedLength) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	return f.Fixed == len(param)
}
