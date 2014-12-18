package validator

import "regexp"

const (
	//PatternURL (forward match)
	PatternURL string = "^(http|https|ftp):\\/\\/([a-zA-Z0-9][a-zA-Z0-9_-]*(?:\\.[a-zA-Z0-9][a-zA-Z0-9_-]*)+):?(\\d+)?\\/?"
	//PatternDomain Domain
	PatternDomain string = "^([a-zA-Z0-9][a-zA-Z0-9_-]*(?:\\.[a-zA-Z0-9][a-zA-Z0-9_-]*)+)$"
	//PatternEmail E-mail (Allow continuous dot)
	PatternEmail string = "^[^0-9][.a-zA-Z0-9_]+[@][a-zA-Z0-9_]+([.][a-zA-Z0-9_]+)*[.][a-zA-Z]{2,4}$"
	//PatternIpv4 IP Address (v4)
	PatternIpv4 string = "^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
)

// RegExp Check regular expression
type RegExp struct {
	Pattern string
}

// URL Check a URL
type URL struct{}

// Domain ... check a Domain
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

// Validate for Validator interface
func (d URL) Validate(param string) bool {

	re := RegExp{PatternURL}
	_, ok := re.Check(param)
	return ok
}

// Validate for Validator interface
func (d Domain) Validate(param string) bool {
	re := RegExp{PatternDomain}
	_, ok := re.Check(param)
	return ok
}

// Validate for Validator interface
func (d Email) Validate(param string) bool {

	re := RegExp{PatternEmail}
	_, ok := re.Check(param)
	return ok
}

// Validate for Validator interface
func (d Ipv4) Validate(param string) bool {

	re := RegExp{PatternIpv4}
	_, ok := re.Check(param)
	return ok
}
