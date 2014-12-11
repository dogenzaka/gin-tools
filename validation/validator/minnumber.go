package validator

import "strconv"

// MinNumber ... check min number
type MinNumber struct {
	Min int
}

// MinNumberIfNotEmpty ... check a number if not empty
type MinNumberIfNotEmpty struct {
	Min int
}

// Validate ... validate param of min number
func (m MinNumber) Validate(param string) bool {

	if param == "" {
		return false
	}

	v, err := strconv.Atoi(param)
	return err == nil && !(v < m.Min)
}

// Validate ... validate param of min number if not empty
func (m MinNumberIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	v, err := strconv.Atoi(param)
	return err == nil && !(v < m.Min)
}
