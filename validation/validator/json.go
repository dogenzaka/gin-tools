package validator

import "encoding/json"

// JSON validator
type JSON struct{}

// Check a param is json variable
func (j JSON) Check(param string) (m map[string]interface{}, ok bool) {

	if isBlank(param) {
		return
	}

	ok = json.Unmarshal([]byte(param), &m) == nil

	return
}

// Validate for Validator interface
func (j JSON) Validate(param string) bool {
	_, ok := j.Check(param)
	return ok
}
