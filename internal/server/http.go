// Package server sets up the Gin engine, attaches global middleware
// and registers all route groups exported by the router package.
package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/AurChatOrg/aurchat-server/internal/router"
)

func NewHTTPServer(cfg *config.Config, log *zap.Logger) *http.Server {
	if cfg.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()

	// Global middleware
	engine.Use(ginzap.Ginzap(log, time.RFC3339, false))
	engine.Use(ginzap.RecoveryWithZap(log, true))

	// Register Route
	router.RegisterAPI(engine, cfg) // /api/*
	//router.RegisterWS(engine, cfg)  // /ws
	//router.RegisterRTC(engine, cfg) // /rtc/*

	return &http.Server{
		Addr:         cfg.HTTP.Listen,
		Handler:      engine,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}
