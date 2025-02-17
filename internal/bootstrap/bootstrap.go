package bootstrap

import (
	"my-echo-app/internal/adapter/db"
	repoAdapter "my-echo-app/internal/adapter/repository" // Alias untuk mencegah konflik
	"my-echo-app/internal/app"
	"my-echo-app/internal/config"
	repoInterface "my-echo-app/internal/interfaces/repository" // Alias agar lebih jelas
	"my-echo-app/internal/interfaces/usecase"
	"my-echo-app/internal/transport/http/handlers"
)

// AppDependencies menyimpan semua instance yang dibutuhkan oleh aplikasi
type AppDependencies struct {
	OrderHandler *handlers.OrderHandler
}

// InitDependencies menginisialisasi semua dependency dalam aplikasi
func InitDependencies() *AppDependencies {
	cfg := config.LoadConfig()
	database := db.ConnectDB(cfg)

	//Order Instance
	var orderRepo repoInterface.OrderRepository = repoAdapter.NewOrderRepository(database)
	var orderUseCase usecase.OrderUseCase = app.NewOrderService(orderRepo)
	var orderHandler = handlers.NewOrderHandler(orderUseCase)

	return &AppDependencies{
		OrderHandler: orderHandler,
	}
}
