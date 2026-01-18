package validator

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,50}$`)
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

// ValidateUsername проверяет валидность username
func ValidateUsername(username string) error {
	if username == "" {
		return errors.New("username is required")
	}
	if len(username) < 4 {
		return errors.New("username must be at least 3 characters")
	}
	if len(username) > 50 {
		return errors.New("username must be at most 50 characters")
	}
	if !usernameRegex.MatchString(username) {
		return errors.New("username can only contain letters, numbers and underscores")
	}
	return nil
}

// ValidateEmail проверяет валидность email
func ValidateEmail(email string) error {
	if email == "" {
		return nil // Email опциональный
	}
	if len(email) > 255 {
		return errors.New("email is too long (max 255 characters)")
	}
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword проверяет валидность пароля
func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password is required")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	if len(password) > 100 {
		return errors.New("password must be at most 100 characters")
	}
	return nil
}

// ValidateBio проверяет валидность bio
func ValidateBio(bio string) error {
	if bio == "" {
		return nil // Bio опциональный
	}
	if len(bio) > 500 {
		return errors.New("bio must be at most 500 characters")
	}
	return nil
}

// ValidateUser проверяет все поля пользователя
func ValidateUser(username, password string, email *string) error {
	// Валидация username
	if err := ValidateUsername(username); err != nil {
		return fmt.Errorf("username validation failed: %w", err)
	}

	// Валидация password
	if err := ValidatePassword(password); err != nil {
		return fmt.Errorf("password validation failed: %w", err)
	}

	// Валидация email (если передан)
	if email != nil && *email != "" {
		if err := ValidateEmail(*email); err != nil {
			return fmt.Errorf("email validation failed: %w", err)
		}
	}

	return nil
}
