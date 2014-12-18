package validator

// NotBlank validator is check a param should not be blank
type NotBlank struct{}

// Validate for Validator interface
func (e NotBlank) Validate(param string) bool {
	return !isBlank(param)
}
