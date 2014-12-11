package validator

import "strconv"

// MaxNumber ... check max number
type MaxNumber struct {
	Max int
}

// MaxNumberIfNotEmpty ... check a number if not empty
type MaxNumberIfNotEmpty struct {
	Max int
}

// Validate ... validate param of max number
func (m MaxNumber) Validate(param string) bool {

	if param == "" {
		return false
	}

	v, err := strconv.Atoi(param)
	return err == nil && !(v > m.Max)
}

// Validate ... validate param of max number if not empty
func (m MaxNumberIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	v, err := strconv.Atoi(param)
	return err == nil && !(v > m.Max)
}
