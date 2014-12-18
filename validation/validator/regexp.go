package validator

import "regexp"

const (
	// URL (forward match)
	C_PATTERN_URL string = "^(http|https|ftp):\\/\\/([a-zA-Z0-9][a-zA-Z0-9_-]*(?:\\.[a-zA-Z0-9][a-zA-Z0-9_-]*)+):?(\\d+)?\\/?"
	// Domain
	C_PATTERN_DOMAIN string = "^([a-zA-Z0-9][a-zA-Z0-9_-]*(?:\\.[a-zA-Z0-9][a-zA-Z0-9_-]*)+)$"
	// E-mail (Allow continuous dot)
	C_PATTERN_EMAIL string = "^[^0-9][.a-zA-Z0-9_]+[@][a-zA-Z0-9_]+([.][a-zA-Z0-9_]+)*[.][a-zA-Z]{2,4}$"
	// IP Address (v4)
	C_PATTERN_IPV4 string = "^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
)

// RegExp ... check regular expression
type RegExp struct {
	Pattern string
}

// Url ... check a positive number
type Url struct{}

// Domain ... check a positive number
type Domain struct{}

// Email ... check a email
type Email struct{}

// Ipv4 ... check a IP Address(v4)
type Ipv4 struct{}

// Check a param regular expression
func (re RegExp) Check(param string) (string, bool) {

	if param == "" {
		return param, true
	}
	b := []byte(param)

	v, err := regexp.Match(re.Pattern, b)
	return param, err == nil && v
}

// Validate for Validator interface
func (re RegExp) Validate(param string) bool {
	_, ok := re.Check(param)
	return ok
}

// Validate ... validate param of regular expression
func (d Url) Validate(param string) bool {

	re := RegExp{C_PATTERN_URL}
	_, ok := re.Check(param)
	return ok
}

// Validate ... validate param of regular expression
func (d Domain) Validate(param string) bool {
	re := RegExp{C_PATTERN_DOMAIN}
	_, ok := re.Check(param)
	return ok
}

// Validate ... validate param of regular expression
func (d Email) Validate(param string) bool {

	re := RegExp{C_PATTERN_EMAIL}
	_, ok := re.Check(param)
	return ok
}

// Validate ... validate param of regular expression
func (d Ipv4) Validate(param string) bool {

	re := RegExp{C_PATTERN_IPV4}
	_, ok := re.Check(param)
	return ok
}
