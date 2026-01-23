package validator

import (
	"regexp"

	"github.com/devoraq/Obfuscatorium_backend/internal/domain/exceptions"
)

var (
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,50}$`)
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

// ValidateUsername проверяет валидность username
func ValidateUsername(username string) error {
	if username == "" {
		return exceptions.ErrUsernameRequired
	}
	if len(username) < 3 {
		return exceptions.ErrUsernameTooShort
	}
	if len(username) > 50 {
		return exceptions.ErrUsernameTooLong
	}
	if !usernameRegex.MatchString(username) {
		return exceptions.ErrUsernameInvalidFormat
	}
	return nil
}

// ValidateEmail проверяет валидность email
func ValidateEmail(email string) error {
	if email == "" {
		return nil // Email опциональный
	}
	if len(email) > 255 {
		return exceptions.ErrEmailTooLong
	}
	if !emailRegex.MatchString(email) {
		return exceptions.ErrInvalidEmail
	}
	return nil
}

// ValidatePassword проверяет валидность пароля
func ValidatePassword(password string) error {
	if password == "" {
		return exceptions.ErrPasswordRequired
	}
	if len(password) < 6 {
		return exceptions.ErrPasswordTooShort
	}
	if len(password) > 100 {
		return exceptions.ErrPasswordTooLong
	}
	return nil
}

// ValidateBio проверяет валидность bio
func ValidateBio(bio string) error {
	if bio == "" {
		return nil // Bio опциональный
	}
	if len(bio) > 500 {
		return exceptions.ErrBioTooLong
	}
	return nil
}

// ValidateUser проверяет все поля пользователя
func ValidateUser(username, password string, email *string) error {
	// Валидация username
	if err := ValidateUsername(username); err != nil {
		return err
	}

	// Валидация password
	if err := ValidatePassword(password); err != nil {
		return err
	}

	// Валидация email (если передан)
	if email != nil && *email != "" {
		if err := ValidateEmail(*email); err != nil {
			return err
		}
	}

	return nil
}
