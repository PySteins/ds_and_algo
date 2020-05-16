package utils

import "errors"

func ElementNotNilCheck(e interface{}) error {
	if e == nil {
		return errors.New("element must not nil")
	}
	return nil
}
