package app

import (
	"github.com/fasthttp/router"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"currency-telegram-bot/internal/config"
)

type App struct {
	conf *config.ServiceConfig
	log  *zap.SugaredLogger
	bot  *telego.Bot
	bh   *th.BotHandler
	i18n *i18n
}

func CreateApp(c *config.ServiceConfig, logger *zap.SugaredLogger) *App {
	app := &App{
		conf: c,
		log:  logger,
		i18n: newI18n(),
	}

	return app
}

func (app *App) Run() {
	var err error
	app.bot, err = telego.NewBot(app.conf.BotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		app.log.Fatalw("Create bot", "error", err)
	}

	srv := telego.FastHTTPWebhookServer{
		Server: &fasthttp.Server{},
		Router: router.New(),
	}

	updates, err := app.bot.UpdatesViaWebhook("/bot/webhook",
		telego.WithWebhookServer(srv),

		telego.WithWebhookSet(&telego.SetWebhookParams{
			URL: app.conf.WebhookBase + "/bot/webhook",
		}),
	)

	if err != nil {
		app.log.Fatalw("Updates via webhook", "error", err)
	}

	app.bh, err = th.NewBotHandler(app.bot, updates)
	if err != nil {
		app.log.Fatalw("Bot handler", "error", err)
	}

	// register bot handlers
	app.registerHandlers()

	// start bot handler
	go app.bh.Start()
	app.log.Info("Handling updates...")

	// start bot webhook server
	go func() {
		err = app.bot.StartWebhook(app.conf.ListenAddress)
		if err != nil {
			app.log.Fatalw("Failed to start webhook", "error", err)
		}
	}()
}

func (app *App) Shutdown() error {

	err := app.bot.StopWebhook()
	if err != nil {
		return err
	}

	// Stop handling updates
	app.bh.Stop()

	return app.bot.DeleteWebhook(nil)
}
