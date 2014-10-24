package validator

import "encoding/hex"

type ObjectId struct {
}

type ObjectIdIfNotEmpty struct {
}

func (o ObjectId) Validate(param string) bool {

	if len(param) != 24 {
		return false
	}
	_, err := hex.DecodeString(param)
	return err == nil
}

func (o ObjectIdIfNotEmpty) Validate(param string) bool {

	if param == "" {
		return true
	}

	if len(param) != 24 {
		return false
	}
	_, err := hex.DecodeString(param)
	return err == nil
}
