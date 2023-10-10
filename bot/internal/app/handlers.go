package app

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (app *App) registerHandlers() {
	app.bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := update.Message.From.ID
		lang := update.Message.From.LanguageCode

		// Creating inline keyboard
		inlineKeyboard := tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton(app.i18n.get("openWebApp", lang)).WithWebApp(&telego.WebAppInfo{URL: app.conf.WebAppUrl}),
			))

		// Call method sendMessage.
		// Send a message to sender with the same text (echo bot).
		// (https://core.telegram.org/bots/api#sendmessage)
		bot.SendMessage(
			tu.Message(
				tu.ID(chatID),
				app.i18n.get("mainMessage", lang),
			).WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML),
		)
	}, th.AnyMessage())
}
