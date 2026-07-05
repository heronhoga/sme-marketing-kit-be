package utils

import (
	"errors"
	"regexp"
)



func CheckEmailPattern(input string) error {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(input) {
		return errors.New("Invalid Email Format")
	}

	return nil
}