package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/AurChatOrg/aurchat-server/docs"
	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/logger"
	"github.com/AurChatOrg/aurchat-server/internal/repo"
	"github.com/AurChatOrg/aurchat-server/internal/server"
	"go.uber.org/zap"
)

// @title        AurChat Server
// @version      1.0
// @description  High-performance chat backend written in Go
// @license.name	MIT License
// @BasePath     /api/v1
func main() {
	cfg := config.Load()                  // Load Config
	log := logger.InitLogger(cfg.App.Env) // Init Logger
	defer func() { _ = log.Sync() }()

	if err := repo.InitPostgres(cfg); err != nil { // Init Postgres
		log.Error("Init postgres", zap.Error(err))
		os.Exit(1)
	}
	log.Info("Init postgres", zap.String("status", "success"))

	srv := server.NewHTTPServer(cfg, log)

	// Start http server
	go func() {
		log.Info("Starting gateway", zap.String("listen", cfg.HTTP.Listen))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("Server stopped", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", zap.Error(err))
	}
	log.Info("Server exited")
}
