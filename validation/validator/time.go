package validator

import "time"

// Time validator is check a param should be obey by format
type Time struct {
	Format string
}

// Validate for Validator interface
func (t Time) Validate(param string) bool {

	if isBlank(param) {
		return true
	}

	_, err := time.Parse(t.Format, param)

	return err == nil
}
