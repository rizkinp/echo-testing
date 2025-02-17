package db

import (
	"database/sql"
	"fmt"
	"log"
	"my-echo-app/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database tidak dapat diakses: %v", err)
	}

	log.Println("Sukses koneksi ke database")
	return db
}
