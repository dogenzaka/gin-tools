package validation

// Validator .. validator interface
type Validator interface {
	Validate(param string) bool
}
