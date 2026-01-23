package usecases

import (
	"context"

	"github.com/devoraq/Obfuscatorium_backend/internal/domain/exceptions"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	userpb "github.com/devoraq/Obfuscatorium_backend/pkg/gen/go/user/v1"
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

func (uc *UserUseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return uc.userRepo.GetByUsername(ctx, username)
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := validator.ValidateUser(user.Username, user.Password, &user.Email); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, exceptions.ErrPasswordHashFailed
	}

	newUser := models.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Email:    user.Email,
	}

	return uc.userRepo.Create(ctx, &newUser)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, id uuid.UUID, req *userpb.UpdateUserRequest) (*models.User, error) {
	updates := make(map[string]any)

	if req.UpdateMask == nil || len(req.UpdateMask.GetPaths()) == 0 {
		return nil, exceptions.ErrUpdateMaskRequired
	}

	for _, path := range req.UpdateMask.GetPaths() {
		switch path {
		case "username":
			if req.Username != nil {
				updates["username"] = req.GetUsername()
			}
		case "email":
			if req.Email != nil {
				updates["email"] = req.GetEmail()
			}
		case "bio":
			if req.Bio != nil {
				updates["bio"] = req.GetBio()
			}
		case "avatar":
			if req.Avatar != nil {
				updates["avatar"] = req.GetAvatar()
			}
		case "password":
			if req.Password != nil && req.GetPassword() != "" {
				hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
				if err != nil {
					return nil, exceptions.ErrPasswordHashFailed
				}
				updates["password_hash"] = string(hash)
			}
		case "role":
			if req.Role != nil {
				updates["role"] = req.GetRole()
			}
		}
	}

	return uc.userRepo.Update(ctx, id, updates)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return uc.userRepo.Delete(ctx, id)
}
