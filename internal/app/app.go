package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

// Run запускает приложение
// Здесь происходит инициализация всех компонентов:
// - Config (internal/infra/config)
// - Logger (internal/infra/logger)
// - Database (internal/infra/postgres)
// - Repositories (internal/infra/postgres/storages)
// - UseCases (internal/domain/usecases)
// - Handlers (internal/api/http/handlers)
// - Router (internal/api/http)
// - Services (internal/services)
func Run() error {
	// TODO: Инициализация компонентов согласно новой архитектуре
	// 1. Загрузка конфигурации
	// 2. Инициализация логгера
	// 3. Подключение к БД
	// 4. Создание репозиториев
	// 5. Создание use cases
	// 6. Создание handlers
	// 7. Настройка роутера
	// 8. Запуск сервисов (workers)
	// 9. Запуск HTTP сервера

	app := fiber.New()

	// Graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := app.Listen(":8080"); err != nil {
			// log error
		}
	}()

	<-ctx.Done()
	return app.Shutdown()
}

