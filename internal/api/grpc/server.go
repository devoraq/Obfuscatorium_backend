package grpc

import (
	"google.golang.org/grpc"

	"github.com/devoraq/Obfuscatorium_backend/internal/api/grpc/user"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/usecases"
	userpb "github.com/devoraq/Obfuscatorium_backend/pkg/gen/go/user/v1"
)

// Dependencies содержит все зависимости для регистрации gRPC сервисов
type Dependencies struct {
	UserUseCase *usecases.UserUseCase
}

// NewServer создаёт и настраивает gRPC сервер с зарегистрированными сервисами
func NewServer(deps Dependencies) *grpc.Server {
	server := grpc.NewServer()

	// Регистрация сервисов
	userService := user.NewUserService(deps.UserUseCase)
	userpb.RegisterUserServiceServer(server, userService)

	return server
}