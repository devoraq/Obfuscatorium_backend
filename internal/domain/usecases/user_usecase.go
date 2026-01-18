package usecases

import (
	"context"
	"fmt"

	"github.com/devoraq/Obfuscatorium_backend/internal/api/http/dto"
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

func (uc *UserUseCase) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
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

func (uc *UserUseCase) UpdateUser(ctx context.Context, id uuid.UUID, req dto.UpdateUserRequest) (*models.User, error) {
	updates := make(map[string]any)

	if req.Username != nil {
		updates["username"] = *req.Username
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.Bio != nil {
		updates["bio"] = *req.Bio
	}
	if req.Avatar != nil {
		updates["avatar"] = *req.Avatar
	}

	if req.Password != nil && *req.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		updates["password_hash"] = string(hash)
	}

	return uc.userRepo.Update(ctx, id, updates)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return uc.userRepo.Delete(ctx, id)
}
