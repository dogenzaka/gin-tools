package validation

// Validator interface provides a "Validate" method
type Validator interface {
	Validate(param string) bool
}
