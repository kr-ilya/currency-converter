package app

import (
	"context"
	"currency-telegram-webapp-backend/internal/config"
	"currency-telegram-webapp-backend/internal/provider"
	"currency-telegram-webapp-backend/internal/store"
	resp "currency-telegram-webapp-backend/pkg/api_response"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func notFoundHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(resp.New(false, utils.StatusMessage(fiber.StatusNotFound)))
}

type Translations map[provider.Currency]provider.Tr

type SymbolInfo struct {
	Rate   float64 `json:"rate"`
	HasImg bool    `json:"hasImg"`
}

type Symbols map[provider.Currency]SymbolInfo

type GetRateResult struct {
	Timestamp    int64        `json:"timestamp"`
	Symbols      Symbols      `json:"symbols"`
	Translations Translations `json:"translations,omitempty"`
}

// get currencies rates
// ex: /api/rates?tr=1
func (a *App) getRate(ctx *fiber.Ctx) error {

	withTranslations := false
	if ctx.Query("tr") == "1" {
		withTranslations = true
	}

	// get update time
	updateTime, err := a.store.GetUpdateTime()
	if err != nil {
		if !errors.Is(err, store.ErrNotFound) {
			a.log.Errorw("failed get updateTime", "error", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(resp.New(false, utils.StatusMessage(fiber.StatusInternalServerError)))
		}

		updateTime = 0
	}

	// fetch rates from provider
	// if the data was updated > updatePeriod ago or update time was not found
	if updateTime == 0 || time.Now().Unix()-updateTime > config.UpdatePeriod {
		data, err := a.p.Latest(context.Background())
		if err != nil {
			a.log.Errorw("failed receive rates from provider", "error", err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(resp.New(false, utils.StatusMessage(fiber.StatusInternalServerError)))
		}

		if err := a.store.SetRates(data.Base, data.Timestamp, &data.Rates); err != nil {
			a.log.Errorw("failed set rates to store", "error", err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(resp.New(false, utils.StatusMessage(fiber.StatusInternalServerError)))
		}

		updateTime = data.Timestamp
	}

	rates := &GetRateResult{
		Symbols:      make(Symbols),
		Translations: make(Translations),
	}

	rates.Timestamp = updateTime

	// get rates
	result, err := a.store.GetRates()
	if err != nil {
		a.log.Errorw("failed get rates from store", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp.New(false, utils.StatusMessage(fiber.StatusInternalServerError)))
	}

	for s, r := range *result {
		if provider.SymbolsList[provider.Currency(s)].Enabled {
			rates.Symbols[s] = SymbolInfo{
				Rate:   r,
				HasImg: provider.SymbolsList[provider.Currency(s)].HasImg,
			}

			if withTranslations {
				rates.Translations[provider.Currency(s)] = provider.SymbolsList[provider.Currency(s)].Translation
			}
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(resp.New(true, rates))
}
