package validator

// Enum validator
type Enum struct {
	Enums []string
}

// Check a param is included in Enums
func (e Enum) Check(param string) (string, bool) {

	if isBlank(param) {
		return param, true
	}

	for _, enum := range e.Enums {
		if param == enum {
			return param, true
		}
	}

	return param, false
}

// Validate for Validator interface
func (e Enum) Validate(param string) bool {
	_, ok := e.Check(param)
	return ok
}
