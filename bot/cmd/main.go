package main

import (
	a "currency-telegram-bot/internal/app"
	"currency-telegram-bot/internal/config"
	"currency-telegram-bot/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Get config: %v\n", err)
		return
	}

	log, err := logger.NewLogger(logger.Config{
		Type: c.LoggerType,
	})
	if err != nil {
		fmt.Printf("Create logger: %v\n", err)
		return
	}

	app := a.CreateApp(c, log)

	done := make(chan struct{}, 1)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		err = app.Shutdown()
		if err != nil {
			log.Fatalw("Shutdown", "error", err)
		}

		done <- struct{}{}
	}()

	// run bot
	app.Run()

	fmt.Println("Started")

	<-done
}
