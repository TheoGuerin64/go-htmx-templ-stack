package main

import (
	"log/slog"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/database"
	"project/internal/handlers"
	"project/internal/middleware"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

func handleKillSignal(f func()) {
	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)

	go f()

	<-killSignal
}

func setupLogger() {
	handler := slog.NewTextHandler(os.Stdout, &config.SlogHandlerOptions)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func setupSentry(settings *config.Settings) error {
	if err := sentry.Init(config.SentryClientOptions(settings.SentryDSN)); err != nil {
		return err
	}
	return nil
}

func setupMiddlewares(e *echo.Echo, settings *config.Settings, db *bun.DB) {
	e.Use(middleware.GetContextMiddleware(settings, db))
	e.Use(echomiddleware.RequestLoggerWithConfig(config.GetLoggerConfig(slog.Default())))
	e.Use(echomiddleware.RecoverWithConfig(config.RecoverConfig))
	if settings.Environnement.IsProd() {
		e.Use(sentryecho.New(config.SentryEchoOptions))
	}
}

func setupRoutes(e *echo.Echo) {
	e.Static("/assets", "assets")
	e.HTTPErrorHandler = handlers.NewErrorHandler().ServeHTTPError
	e.GET("/", handlers.NewHomeHandler().ServeHTTP)
}

func main() {
	setupLogger()
	settings, err := config.LoadSettings()
	if err != nil {
		slog.Error("failed to load settings", slog.Any("error", err))
		os.Exit(1)
	}
	db, err := database.Open(settings)
	if err != nil {
		slog.Error("failed to open database", slog.Any("error", err))
		os.Exit(1)
	}
	if settings.Environnement.IsProd() {
		if err := setupSentry(settings); err != nil {
			slog.Error("failed to setup sentry", slog.Any("error", err))
			os.Exit(1)
		}
	}

	e := echo.New()
	setupMiddlewares(e, settings, db)
	setupRoutes(e)

	handleKillSignal(func() {
		if err := e.Start(settings.Address); err != nil {
			slog.Error("failed to start server", slog.Any("error", err))
			os.Exit(1)
		}
	})

	slog.Info("shutting down server")
}
