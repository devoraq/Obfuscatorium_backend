package exceptions

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidEmail      = errors.New("invalid email format")
	ErrUserNotFound      = errors.New("user not found")
	
	// Storage errors
	ErrNoFieldsToUpdate  = errors.New("no fields to update")
	ErrDatabaseError      = errors.New("database error")
	ErrQueryBuildFailed   = errors.New("failed to build query")
	ErrQueryExecutionFailed = errors.New("db execution failed")
	
	// UseCase errors
	ErrPasswordHashFailed = errors.New("failed to hash password")
	ErrUpdateMaskRequired  = errors.New("update_mask is required")
	
	// Validation errors
	ErrUsernameRequired = errors.New("username is required")
	ErrUsernameTooShort = errors.New("username must be at least 3 characters")
	ErrUsernameTooLong = errors.New("username must be at most 50 characters")
	ErrUsernameInvalidFormat = errors.New("username can only contain letters, numbers and underscores")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordTooShort = errors.New("password must be at least 6 characters")
	ErrPasswordTooLong = errors.New("password must be at most 100 characters")
	ErrEmailTooLong = errors.New("email is too long (max 255 characters)")
	ErrBioTooLong = errors.New("bio must be at most 500 characters")
)
