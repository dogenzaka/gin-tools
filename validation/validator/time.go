package validator

import "time"

// Time validator is check a param should be obey by format
type Time struct {
	Format string
}

// Validate for Validator interface
func (t Time) Check(param string) (time.Time, bool) {

	if isBlank(param) {
		return time.Time{}, true
	}

	v, err := time.Parse(t.Format, param)

	return v, err == nil
}

// Validate for Validator interface
func (t Time) Validate(param string) bool {
	_, ok := t.Check(param)
	return ok
}
