package app

import (
	"context"
	"errors"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"fmt"

	resp "currency-telegram-webapp-backend/pkg/api_response"

	"currency-telegram-webapp-backend/internal/client"
	"currency-telegram-webapp-backend/internal/config"
	"currency-telegram-webapp-backend/internal/provider"
	"currency-telegram-webapp-backend/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/utils"
)

type App struct {
	server       *fiber.App
	conf         *config.ServiceConfig
	log          *zap.SugaredLogger
	store        store.Store
	baseCurrency provider.Currency
	p            provider.Provider
	api          *client.Caller
}

func CreateApp(c *config.ServiceConfig, logger *zap.SugaredLogger) *App {
	clientApi := &client.Caller{
		Client: &fasthttp.Client{},
	}

	app := &App{
		conf: c,
		log:  logger,
		server: fiber.New(fiber.Config{
			AppName:                  "backend currency telegram webapp",
			DisableStartupMessage:    false,
			JSONEncoder:              json.Marshal,
			JSONDecoder:              json.Unmarshal,
			ErrorHandler:             errorHandler(logger),
			EnableSplittingOnParsers: true,
		}),
		store:        store.NewRedisStore(&c.Redis),
		baseCurrency: "",
		p:            provider.NewFixedProvider(clientApi, c.FixerAccessToken),
	}

	// auth middleware
	app.server.Use(app.AuthMw)

	app.server.Get("/api/rates", app.getRate)

	// panic recover
	app.server.Use(recover.New())

	// custom 404 handler
	app.server.Use(notFoundHandler)

	if err := app.initBaseCurrency(); err != nil {
		app.log.Fatalw("Init base currency", "error", err)
	}

	return app
}

func (app *App) Run() {
	go func() {
		if err := app.server.Listen(app.conf.ListenAddress); err != nil {
			app.log.Fatalw("Start server", "error", err)
		}
	}()
}

func (app *App) Shutdown() error {
	return app.server.ShutdownWithTimeout(config.ShutdownTimeout)
}

func errorHandler(log *zap.SugaredLogger) func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError
		errData := utils.StatusMessage(code)

		// Retrieve the custom status code if it's a *fiber.Error
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
			errData = fiberErr.Message
		}

		log.Errorf("API panic recovered: %v", err)

		return ctx.Status(code).JSON(resp.New(false, errData))
	}
}

func (app *App) initBaseCurrency() error {
	base, err := app.store.GetBase()
	if err != nil {
		if !errors.Is(err, store.ErrNotFound) {
			return fmt.Errorf("failed get base currency %w", err)
		}

		data, err := app.p.Latest(context.Background())
		if err != nil {
			return err
		}

		if err := app.store.SetRates(data.Base, data.Timestamp, &data.Rates); err != nil {
			return err
		}
	}

	app.baseCurrency = base
	return nil
}
