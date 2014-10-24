package validator

// MaxLength ... check max length
type MaxLength struct {
	Max int
}

// Validate ... max length validate
func (m MaxLength) Validate(param string) bool {
	if len(param) > m.Max {
		return false
	}
	return true
}
