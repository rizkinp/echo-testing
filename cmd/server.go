package cmd

import (
	"my-echo-app/internal/bootstrap"
	"my-echo-app/internal/transport/http/routes"
	"my-echo-app/pkg/logger"
	"my-echo-app/pkg/tracing"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// StartServer menginisialisasi server Echo dengan Jaeger & Middleware
func StartServer() {
	e := echo.New()

	// ✅ Inisialisasi Jaeger
	_, closer := tracing.InitTracer("echo-server")
	defer closer()
	// ✅ Inisialisasi Logger
	logger.InitLogger()
	// ✅ Setup Jaeger Tracing Middleware
	c := jaegertracing.New(e, nil)
	defer c.Close()
	// ✅ Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	// ✅ Middleware Logging & Recovery
	e.Use(middleware.LoggerWithConfig(logger.InitLogger()))
	// ✅ Middleware Recovery
	e.Use(middleware.Recover())
	// ✅ Tambahkan Middleware Prometheus
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e) // 🔥 Aktifkan endpoint `/metrics` otomatis!
	// ✅ Inisialisasi Dependencies
	deps := bootstrap.InitDependencies()
	// ✅ Setup Routes
	cfg := &routes.RoutesConfig{
		OrderHandler: deps.OrderHandler,
	}
	routes.SetupRoutes(e, cfg)
	// ✅ Jalankan Server
	e.Logger.Fatal(e.Start(":5173"))
}
