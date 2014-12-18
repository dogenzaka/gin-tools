package validator

import "regexp"

// RegExp ... check regular expression
type RegExp struct {
	Pattern string
}

// Validate ... validate param of regular expression
func (re RegExp) Validate(param string) bool {

	b := []byte(param)

	v, err := regexp.Match(re.Pattern, b)
	return err == nil && v
}
