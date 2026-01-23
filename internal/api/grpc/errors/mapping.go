package errors

import (
	"errors"

	"github.com/devoraq/Obfuscatorium_backend/internal/domain/exceptions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, exceptions.ErrUserAlreadyExists) {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	if errors.Is(err, exceptions.ErrInvalidEmail) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, exceptions.ErrUserNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, exceptions.ErrNoFieldsToUpdate) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, exceptions.ErrUpdateMaskRequired) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, exceptions.ErrPasswordHashFailed) {
		return status.Error(codes.Internal, err.Error())
	}
	if errors.Is(err, exceptions.ErrDatabaseError) {
		return status.Error(codes.Internal, err.Error())
	}
	if errors.Is(err, exceptions.ErrQueryBuildFailed) {
		return status.Error(codes.Internal, err.Error())
	}
	if errors.Is(err, exceptions.ErrQueryExecutionFailed) {
		return status.Error(codes.Internal, err.Error())
	}

	// Validation errors
	if errors.Is(err, exceptions.ErrUsernameRequired) ||
		errors.Is(err, exceptions.ErrUsernameTooShort) ||
		errors.Is(err, exceptions.ErrUsernameTooLong) ||
		errors.Is(err, exceptions.ErrUsernameInvalidFormat) ||
		errors.Is(err, exceptions.ErrPasswordRequired) ||
		errors.Is(err, exceptions.ErrPasswordTooShort) ||
		errors.Is(err, exceptions.ErrPasswordTooLong) ||
		errors.Is(err, exceptions.ErrEmailTooLong) ||
		errors.Is(err, exceptions.ErrBioTooLong) {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return status.Error(codes.Internal, "internal server error")
}
