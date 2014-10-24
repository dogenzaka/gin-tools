package validator

// FixedLength ... check fixed length
type FixedLength struct {
	Fixed int
}

// FixedLengthIfNotEmpty ... check fixed length if not empty
type FixedLengthIfNotEmpty struct {
	Fixed int
}

// Validate ... validate param of fixed length
func (f FixedLength) Validate(param string) bool {
	if f.Fixed == len(param) {
		return true
	}
	return false
}

// Validate ... validate param of fixed length if not empty
func (f FixedLengthIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	if f.Fixed == len(param) {
		return true
	}
	return false
}
