package validator

// Enum validator is check a param is included in Enums
type Enum struct {
	Enums []string
}

// Validate for Validator interface
func (e Enum) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	for _, enum := range e.Enums {
		if param == enum {
			return true
		}
	}

	return false
}
