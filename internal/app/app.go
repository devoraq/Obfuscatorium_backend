package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" // Важно: v2
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcapi "github.com/devoraq/Obfuscatorium_backend/internal/api/grpc"
	"github.com/devoraq/Obfuscatorium_backend/internal/config"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/usecases"
	"github.com/devoraq/Obfuscatorium_backend/internal/infra/postgres"
	"github.com/devoraq/Obfuscatorium_backend/internal/infra/postgres/storages"
	userpb "github.com/devoraq/Obfuscatorium_backend/pkg/gen/go/user/v1"
)

type App struct {
	grpcServer *grpc.Server
	httpServer *http.Server
}

func New(cfg *config.Config) *App {
	// 1. Инфраструктура и логика
	db, err := postgres.New(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	userStorage := storages.NewUserStorage(db.GetConn())
	userUseCase := usecases.NewUserUseCase(userStorage)

	// 2. Настройка gRPC сервера с регистрацией сервисов
	gServer := grpcapi.NewServer(grpcapi.Dependencies{
		UserUseCase: userUseCase,
	})

	// 3. Настройка gRPC-Gateway (вместо старых хттп хендлеров)
	mux := runtime.NewServeMux()

	// Опции для подключения Gateway к gRPC (внутри одной сети — insecure)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Регистрируем Gateway. Он будет стучаться на порт gRPC (:50051)
	err = userpb.RegisterUserServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:50051",
		opts,
	)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	return &App{
		grpcServer: gServer,
		httpServer: &http.Server{
			Addr:    ":8080", // Порт для REST/JSON
			Handler: mux,
		},
	}
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Запуск gRPC (:50051)
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("grpc listen error: %v", err)
		}
		if err := a.grpcServer.Serve(lis); err != nil {
			log.Printf("grpc server stopped: %v", err)
		}
	}()

	// Запуск HTTP Gateway (:8080)
	go func() {
		log.Printf("Gateway (REST) started on :8080")
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("gateway error: %v", err)
		}
	}()

	<-ctx.Done()

	// Завершаем работу
	a.grpcServer.GracefulStop()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return a.httpServer.Shutdown(shutdownCtx)
}
