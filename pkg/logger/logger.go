package logger

import (
	"log"
	"os"

	"github.com/labstack/echo/v4/middleware"
)

// InitLogger menginisialisasi logger global
func InitLogger() middleware.LoggerConfig {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return middleware.LoggerConfig{
		Format: `method=${method}, uri=${uri}, status=${status}, error=${error}, latency=${latency_human}` + "\n",
		Output: os.Stdout,
	}
}
