package utils

import (
	"unicode"
)

// IsStrongPassword checks if the given password meets the criteria for a strong password.
// A strong password must:
// - Be at least 8 characters long
// - Contain at least one lowercase letter
// - Contain at least one uppercase letter
// - Contain at least one digit
// - Contain at least one special character from the set @#$%^&*()_+!
//
// Parameters:
// - password: The password string to be validated.
//
// Returns:
// - bool: true if the password is strong, false otherwise.
func IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}
