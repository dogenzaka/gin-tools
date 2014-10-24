package validator

import "strconv"

// PNumber ... check a positive number
type PNumber struct {
}

// PNumberIfNotEmpty ... check a positive number if not empty
type PNumberIfNotEmpty struct {
}

// Validate ... validate param for a positive number
func (p PNumber) Validate(param string) bool {
	v, err := strconv.Atoi(param)
	if err == nil && v > 0 {
		return true
	}
	return false
}

// Validate ... validate param for a positive number if not empty
func (p PNumberIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	v, err := strconv.Atoi(param)
	if err == nil && v > 0 {
		return true
	}
	return false
}
