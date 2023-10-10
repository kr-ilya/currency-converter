package app

import (
	resp "currency-telegram-webapp-backend/pkg/api_response"
	"encoding/hex"
	"io"
	"net/url"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"crypto/hmac"
	"crypto/sha256"
)

// Auth middleware
func (a *App) AuthMw(ctx *fiber.Ctx) error {
	q := ctx.Query("auth")
	if q == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(resp.New(false, utils.StatusMessage(fiber.StatusUnauthorized)))
	}

	params, err := url.ParseQuery(q)
	if err != nil {
		a.log.Errorw("failed ParseQuery auth", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp.New(false, utils.StatusMessage(fiber.StatusInternalServerError)))
	}

	data := make([]string, 0, len(params)-1)
	for k, v := range params {
		if k != "hash" {
			data = append(data, k+"="+v[0])
		}
	}

	sort.Strings(data)

	res := ""
	for _, v := range data {
		if res != "" {
			res += "\n"
		}

		res += v
	}

	secretKey := hmac.New(sha256.New, []byte("WebAppData"))
	if _, err = io.WriteString(secretKey, a.conf.BotToken); err != nil {
		a.log.Errorw("failed create hmac", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp.New(false, utils.StatusMessage(fiber.StatusInternalServerError)))
	}

	hash := hmac.New(sha256.New, secretKey.Sum(nil))
	if _, err = io.WriteString(hash, res); err != nil {
		a.log.Errorw("failed create hmac", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp.New(false, utils.StatusMessage(fiber.StatusInternalServerError)))
	}

	if hex.EncodeToString(hash.Sum(nil)) != params.Get("hash") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(resp.New(false, utils.StatusMessage(fiber.StatusUnauthorized)))
	}

	return ctx.Next()
}
