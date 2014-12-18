package validator

// NotBlank validator
type NotBlank struct{}

// Check a param should not be blank
func (e NotBlank) Check(param string) (string, bool) {
	return param, !isBlank(param)
}

// Validate for Validator interface
func (e NotBlank) Validate(param string) bool {
	_, ok := e.Check(param)
	return ok
}
