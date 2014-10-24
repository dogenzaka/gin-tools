package validator

import "encoding/hex"

// ObjectID ... check objectID
type ObjectID struct {
}

// ObjectIDIfNotEmpty ... check objectID if not empty
type ObjectIDIfNotEmpty struct {
}

// Validate ... validate param of objectID
func (o ObjectID) Validate(param string) bool {

	if len(param) != 24 {
		return false
	}
	_, err := hex.DecodeString(param)
	return err == nil
}

// Validate ... validate param of objectID if not empty
func (o ObjectIDIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	if len(param) != 24 {
		return false
	}
	_, err := hex.DecodeString(param)
	return err == nil
}
