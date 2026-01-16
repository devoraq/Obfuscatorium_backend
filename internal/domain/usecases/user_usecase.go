package usecases

import (
	"context"
	"fmt"

	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/devoraq/Obfuscatorium_backend/pkg/validator"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo UserRepositoryI
}

func NewUserUseCase(userRepo UserRepositoryI) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc UserUseCase) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := validator.ValidateUser(user.Username, user.Password, &user.Email); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := models.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Email:    user.Email,
	}

	return uc.userRepo.Create(ctx, &newUser)
}

func (uc UserUseCase) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}
