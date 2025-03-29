package auth

import (
	"regexp"
)

type UserCreds struct {
	EmailID         string
	Password        string
	ConfirmPassword string
}

func (u *UserCreds) Validate() (map[string][]string, bool) {
	errors := make(map[string][]string)

	// Validate email format with a basic regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.EmailID) {
		errors["email"] = append(errors["email"], "Invalid email format")
	}

	// Check password length (example: at least 8 characters)
	if len(u.Password) < 8 {
		errors["password"] = append(errors["password"], "Password must be at least 8 characters long")
	}

	// Check password strength: require at least one letter and one number
	hasLetter, hasNumber := false, false
	for _, ch := range u.Password {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			hasLetter = true
		}
		if ch >= '0' && ch <= '9' {
			hasNumber = true
		}
	}
	if !hasLetter || !hasNumber {
		errors["password"] = append(errors["password"], "Password must contain both letters and numbers")
	}

	// If provided, check if confirm password matches
	if u.ConfirmPassword != "" && u.Password != u.ConfirmPassword {
		errors["confirmPassword"] = append(errors["confirmPassword"], "Passwords do not match")
	}

	if len(errors) > 0 {
		return errors, false
	}

	return nil, true
}
