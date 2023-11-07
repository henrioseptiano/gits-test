package utils

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
	if !emailRegex.MatchString(email) {
		return errors.New("Invalid email format")
	}
	return nil
}

func IsValidPhoneNumber(phoneNumber string) bool {
	regex := regexp.MustCompile(`^08[0-9]{8,10}$`)
	return regex.MatchString(phoneNumber)
}
