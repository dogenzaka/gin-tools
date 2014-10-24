package validator

type FixedLength struct {
	Fixed int
}

type FixedLengthIfNotEmpty struct {
	Fixed int
}

func (f FixedLength) Validate(param string) bool {
	if f.Fixed == len(param) {
		return true
	}
	return false
}

func (f FixedLengthIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	if f.Fixed == len(param) {
		return true
	}
	return false
}
