package validator

import "time"

// Time ... check time
type Time struct {
	Format string
}

// TimeIfNotEmpty ... check time if not empty
type TimeIfNotEmpty struct {
	Format string
}

// Validate ... validate param of objectID
func (t Time) Validate(param string) bool {
	_, err := time.Parse(t.Format, param)
	return err == nil
}

// Validate ... validate param of objectID if not empty
func (t TimeIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	_, err := time.Parse(t.Format, param)
	return err == nil
}
