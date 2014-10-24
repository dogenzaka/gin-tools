package validator

import "strconv"

// Bool ... check bool
type Bool struct {
}

// BoolIfNotEmpty ... check bool if not empty
type BoolIfNotEmpty struct {
}

// Validate ... validate param of bool
func (b Bool) Validate(param string) bool {
	_, err := strconv.ParseBool(param)
	if err == nil {
		return true
	}
	return false
}

// Validate ... validate param of bool if not empty
func (b BoolIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	_, err := strconv.ParseBool(param)
	if err == nil {
		return true
	}
	return false
}
