package user

import (
	"context"

	"github.com/devoraq/Obfuscatorium_backend/internal/api/grpc/errors"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/models"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/usecases"
	userpb "github.com/devoraq/Obfuscatorium_backend/pkg/gen/go/user/v1"
	"github.com/devoraq/Obfuscatorium_backend/pkg/token"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	userpb.UnimplementedUserServiceServer
	uc *usecases.UserUseCase
}

func NewUserService(uc *usecases.UserUseCase) userpb.UserServiceServer {
	return &userService{uc: uc}
}

func (s *userService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u := &models.User{Username: req.Username, Password: req.Password, Email: req.Email}

	createdUser, err := s.uc.CreateUser(ctx, u)
	if err != nil {
		return nil, errors.MapError(err)
	}

	return &userpb.CreateUserResponse{User: ToProtoUser(createdUser)}, nil
}

func (s *userService) LoginUser(ctx context.Context, req *userpb.LoginUserRequest) (*userpb.LoginUserResponse, error) {
	u, err := s.uc.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, errors.MapError(err)
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.GetPassword())); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid username or password")
	}

	accessToken, err := token.GenerateToken(u.ID, u.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate access token")
	}

	const expiresIn = 24 * 60 * 60 // 24 часа в секундах

	return &userpb.LoginUserResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
		User:        ToProtoUser(u),
	}, nil
}

func (s *userService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid uuid format: %v", err)
	}

	u, err := s.uc.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.MapError(err)
	}

	return &userpb.GetUserResponse{User: ToProtoUser(u)}, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid uuid format: %v", err)
	}

	u, err := s.uc.UpdateUser(ctx, id, req)
	if err != nil {
		return nil, errors.MapError(err)
	}

	return &userpb.UpdateUserResponse{User: ToProtoUser(u)}, nil
}

func (s *userService) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid uuid format: %v", err)
	}

	if err := s.uc.DeleteUser(ctx, id); err != nil {
		return nil, errors.MapError(err)
	}

	return &userpb.DeleteUserResponse{}, nil
}
