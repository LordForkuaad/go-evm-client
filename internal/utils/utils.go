package utils

import (
	"fmt"
)

// Contains accepts a string and a slice of strings,
// it then verifies if the element exists in the slice.
func Contains(s *[]string, str string) bool {
	for _, v := range *s {
		if v == str {
			return true
		}
	}
	return false
}

// ValidateLength accepts a slice of strings and validates that the
// length of the slice matches the expected amount
func ValidateLength(s *[]string, expected int) error {
	listLen := len(*s)
	if listLen != expected {
		err := fmt.Errorf("error: %d arguments does not match required %d",
			listLen, expected)
		return err
	}
	return nil
}

// RequiredFlagVerification accepts a slice of strings and verifies that
// each string is not empty
func RequiredFlagVerification(flags *[]string) bool {
	for _, flag := range *flags {
		if len(flag) == 0 {
			return false
		}
	}
	return true
}