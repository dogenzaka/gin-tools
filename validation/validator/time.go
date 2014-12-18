package validator

import (
	"strconv"
	"time"
)

// Time validator is check a param should be obey by format
type Time struct {
	Format string
}

// UnixTime validator is check a param
type UnixTime struct{}

// Check Validate for Validator interface
func (t Time) Check(param string) (time.Time, bool) {

	if isBlank(param) {
		return time.Time{}, true
	}

	v, err := time.Parse(t.Format, param)

	return v, err == nil
}

// CheckUnixTime Validate for Validator interface
func (ut UnixTime) CheckUnixTime(param string) (int, bool) {

	if isBlank(param) {
		return 0, true
	}

	v, err := strconv.Atoi(param)

	//Allow since '1970-01-01T00:00:00Z'
	return v, err == nil && v >= 0
}

// Validate for Validator interface
func (t Time) Validate(param string) bool {
	_, ok := t.Check(param)
	return ok
}

// Validate for Validator interface
func (ut UnixTime) Validate(param string) bool {
	_, ok := ut.CheckUnixTime(param)
	return ok
}
