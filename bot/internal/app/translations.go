package app

type lang struct {
	ru string
	en string
}

type i18n struct {
	translations map[string]lang
}

func newI18n() *i18n {
	return &i18n{
		translations: map[string]lang{
			"mainMessage": {
				ru: "<b>Конвертер валют</b>\n\nНажмите кнопку ниже, чтобы открыть мини-приложение.",
				en: "<b>Currency converter</b>\n\nTap the button below to open the MiniApp.",
			},
			"openWebApp": {
				ru: "Открыть",
				en: "Open",
			},
		},
	}
}

func (i *i18n) get(key string, lang string) string {
	if _, ok := i.translations[key]; !ok {
		return ""
	}

	switch lang {
	case "ru":
		return i.translations[key].ru
	case "en":
		return i.translations[key].en
	}

	return ""
}
