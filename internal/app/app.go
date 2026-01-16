package app

import (
	"log"
	"net/http"

	api "github.com/devoraq/Obfuscatorium_backend/internal/api/http"
	"github.com/devoraq/Obfuscatorium_backend/internal/api/http/handlers"
	"github.com/devoraq/Obfuscatorium_backend/internal/config"
	"github.com/devoraq/Obfuscatorium_backend/internal/domain/usecases"
	"github.com/devoraq/Obfuscatorium_backend/internal/infra/postgres"
	"github.com/devoraq/Obfuscatorium_backend/internal/infra/postgres/storages"
)

func Run() error {
	//Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключение к БД
	db, err := postgres.New(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userStorage := storages.NewUserStorage(db.GetConn())
	userUseCase := usecases.NewUserUseCase(userStorage)
	userHandler := handlers.NewUserHandler(userUseCase)

	router := api.NewRouter(api.Deps{
		UserHandler: userHandler,
	})

	return http.ListenAndServe(":8080", router)
}
