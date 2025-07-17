package main

import (
	"app/internal/application"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// Configure structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Change to slog.LevelDebug for more verbose logs
	}))
	slog.SetDefault(logger)

	slog.Info("Starting Fantasy Store API...")

	cfg := &application.ConfigApplicationDefault{
		Db: &mysql.Config{
			User:   os.Getenv("DB_USER"),
			Passwd: os.Getenv("DB_PASSWORD"),
			Net:    "tcp",
			Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
			DBName: os.Getenv("DB_NAME"),
		},
		Addr: os.Getenv("HOST") + ":" + os.Getenv("PORT"),
	}

	app := application.NewApplicationDefault(cfg)

	slog.Info("Application configured",
		"database", cfg.Db.Addr+"/"+cfg.Db.DBName,
		"server_address", cfg.Addr)

	// - set up
	slog.Info("Setting up application...")
	err := app.SetUp()
	if err != nil {
		slog.Error("Failed to set up application", "error", err)
		return
	}
	slog.Info("Application setup completed successfully")

	// - run
	slog.Info("Starting server...")
	err = app.Run()
	if err != nil {
		slog.Error("Failed to run application", "error", err)
		return
	}
}
