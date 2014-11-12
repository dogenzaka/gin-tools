package validator

type Enum struct {
	Enums []string
}

type EnumIfNotEmpty struct {
	Enums []string
}

func (e Enum) Validate(param string) bool {

	for _, enum := range e.Enums {
		if param == enum {
			return true
		}
	}

	return false
}

func (e EnumIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	for _, enum := range e.Enums {
		if param == enum {
			return true
		}
	}

	return false
}
