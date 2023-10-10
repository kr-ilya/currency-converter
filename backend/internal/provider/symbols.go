package provider

type Symbol struct {
	Enabled     bool
	HasImg      bool
	Translation Tr
}

type Tr struct {
	Ru string `json:"ru"`
	En string `json:"en"`
}

var SymbolsList = map[Currency]Symbol{
	Currency("AED"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Дирхам (ОАЭ)",
			En: "United Arab Emirates Dirham",
		},
	},
	Currency("AFN"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Афгани",
			En: "Afghan Afghani",
		},
	},
	Currency("ALL"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Албанский лек",
			En: "Albanian Lek",
		},
	},
	Currency("AMD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Армянский драм",
			En: "Armenian Dram",
		},
	},
	Currency("ANG"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Нидерландский антильский гульден",
			En: "Netherlands Antillean Guilder",
		},
	},
	Currency("AOA"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Ангольская кванза",
			En: "Angolan Kwanza",
		},
	},
	Currency("ARS"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Аргентинское песо",
			En: "Argentine Peso",
		},
	},
	Currency("AUD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Австралийский доллар",
			En: "Australian Dollar",
		},
	},
	Currency("AWG"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Арубанский флорин",
			En: "Aruban Florin",
		},
	},
	Currency("AZN"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Азербайджанский манат",
			En: "Azerbaijani Manat",
		},
	},
	Currency("BAM"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Боснийская конвертируемая марка",
			En: "Bosnia-Herzegovina Convertible Mark",
		},
	},
	Currency("BBD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Барбадосский доллар",
			En: "Barbadian Dollar",
		},
	},
	Currency("BDT"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Бангладешкая така",
			En: "Bangladeshi Taka",
		},
	},
	Currency("BGN"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Болгарский лев",
			En: "Bulgarian Lev",
		},
	},
	Currency("BHD"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Бахрейнский динар",
			En: "Bahraini Dinar",
		},
	},
	Currency("BIF"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Бурундийский франк",
			En: "Burundian Franc",
		},
	},
	Currency("BMD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Бермудский доллар",
			En: "Bermudan Dollar",
		},
	},
	Currency("BND"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Брунейский доллар",
			En: "Brunei Dollar",
		},
	},
	Currency("BOB"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Боливийский боливиано",
			En: "Bolivian Boliviano",
		},
	},
	Currency("BRL"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Бразильский реал",
			En: "Brazilian Real",
		},
	},
	Currency("BSD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Багамский доллар",
			En: "Bahamian Dollar",
		},
	},
	Currency("BTC"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Биткоин",
			En: "Bitcoin",
		},
	},
	Currency("BTN"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Бутанский нгултрум",
			En: "Bhutanese Ngultrum",
		},
	},
	Currency("BWP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Ботсванская пула",
			En: "Botswanan Pula",
		},
	},
	Currency("BYN"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Белорусский рубль",
			En: "Belarusian Ruble",
		},
	},
	Currency("BYR"): {
		Enabled: false,
		HasImg:  true,
		Translation: Tr{
			Ru: "Белорусский рубль",
			En: "Belarusian Ruble",
		},
	},
	Currency("BZD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Белизский доллар",
			En: "Belize Dollar",
		},
	},
	Currency("CAD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Канадский доллар",
			En: "Canadian Dollar",
		},
	},
	Currency("CDF"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Конголезский франк",
			En: "Congolese Franc",
		},
	},
	Currency("CHF"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Швейцарский франк",
			En: "Swiss Franc",
		},
	},
	Currency("CLF"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Условная расчетная единица Чили",
			En: "Chilean Unit of Account (UF)",
		},
	},
	Currency("CLP"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Чилийское песо",
			En: "Chilean Peso",
		},
	},
	Currency("CNY"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Китайский юань",
			En: "Chinese Yuan",
		},
	},
	Currency("COP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Колумбийское песо",
			En: "Colombian Peso",
		},
	},
	Currency("CRC"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Коста-риканский колон",
			En: "Costa Rican Colon",
		},
	},
	Currency("CUC"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Кубинское конвертируемое песо",
			En: "Cuban Convertible Peso",
		},
	},
	Currency("CUP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Кубинское песо",
			En: "Cuban Peso",
		},
	},
	Currency("CVE"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Эскудо Кабо-Верде",
			En: "Cape Verdean Escudo",
		},
	},
	Currency("CZK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Чешская крона",
			En: "Czech Republic Koruna",
		},
	},
	Currency("DJF"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Франк Джибути",
			En: "Djiboutian Franc",
		},
	},
	Currency("DKK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Датская крона",
			En: "Danish Krone",
		},
	},
	Currency("DOP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Доминиканское песо",
			En: "Dominican Peso",
		},
	},
	Currency("DZD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Алжирский динар",
			En: "Algerian Dinar",
		},
	},
	Currency("EGP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Египетский фунт",
			En: "Egyptian Pound",
		},
	},
	Currency("ERN"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Эритрейская накфа",
			En: "Eritrean Nakfa",
		},
	},
	Currency("ETB"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Эфиопский быр",
			En: "Ethiopian Birr",
		},
	},
	Currency("EUR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Евро",
			En: "Euro",
		},
	},
	Currency("FJD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Доллар Фиджи",
			En: "Fijian Dollar",
		},
	},
	Currency("FKP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Фунт Фолклендских островов",
			En: "Falkland Islands Pound",
		},
	},
	Currency("GBP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Фунт стерлингов",
			En: "British Pound Sterling",
		},
	},
	Currency("GEL"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Грузинский Лари",
			En: "Georgian Lari",
		},
	},
	Currency("GGP"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Гернсийский фунт",
			En: "Guernsey Pound",
		},
	},
	Currency("GHS"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Ганский седи",
			En: "Ghanaian Cedi",
		},
	},
	Currency("GIP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Гибралтарский фунт",
			En: "Gibraltar Pound",
		},
	},
	Currency("GMD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Гамбийский даласи",
			En: "Gambian Dalasi",
		},
	},
	Currency("GNF"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Гвинейский франк",
			En: "Guinean Franc",
		},
	},
	Currency("GTQ"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Гватемальский кетсаль",
			En: "Guatemalan Quetzal",
		},
	},
	Currency("GYD"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Гайанский доллар",
			En: "Guyanaese Dollar",
		},
	},
	Currency("HKD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Гонконгский доллар",
			En: "Hong Kong Dollar",
		},
	},
	Currency("HNL"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Гондурасская лемпира",
			En: "Honduran Lempira",
		},
	},
	Currency("HRK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Хорватская куна",
			En: "Croatian Kuna",
		},
	},
	Currency("HTG"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Гаитянский гурд",
			En: "Haitian Gourde",
		},
	},
	Currency("HUF"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Венгерский форинт",
			En: "Hungarian Forint",
		},
	},
	Currency("IDR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Индонезийская рупия",
			En: "Indonesian Rupiah",
		},
	},
	Currency("ILS"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Новый израильский шекель",
			En: "Israeli New Sheqel",
		},
	},
	Currency("IMP"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Фунты Острова Мэн",
			En: "Manx pound",
		},
	},
	Currency("INR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Индийская рупия",
			En: "Indian Rupee",
		},
	},
	Currency("IQD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Иракский динар",
			En: "Iraqi Dinar",
		},
	},
	Currency("IRR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Иранский риал",
			En: "Iranian Rial",
		},
	},
	Currency("ISK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Исландская крона",
			En: "Iceland Krona",
		},
	},
	Currency("JEP"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Джерсийский фунт",
			En: "Jersey Pound",
		},
	},
	Currency("JMD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Ямайский доллар",
			En: "Jamaican Dollar",
		},
	},
	Currency("JOD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Иорданский динар",
			En: "Jordanian Dinar",
		},
	},
	Currency("JPY"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Японская иена",
			En: "Japanese Yen",
		},
	},
	Currency("KES"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Кенийский шиллинг",
			En: "Kenyan Shilling",
		},
	},
	Currency("KGS"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Киргизский сом",
			En: "Kyrgystani Som",
		},
	},
	Currency("KHR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Камбоджийский риель",
			En: "Cambodian Riel",
		},
	},
	Currency("KMF"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Франк Комор",
			En: "Comorian Franc",
		},
	},
	Currency("KPW"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Северокорейская вона",
			En: "North Korean Won",
		},
	},
	Currency("KRW"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Южнокорейская вона",
			En: "South Korean Won",
		},
	},
	Currency("KWD"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Кувейтский динар",
			En: "Kuwaiti Dinar",
		},
	},
	Currency("KYD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Доллар Островов Кайман",
			En: "Cayman Islands Dollar",
		},
	},
	Currency("KZT"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Казахстанский тенге",
			En: "Kazakhstani Tenge",
		},
	},
	Currency("LAK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Лаосский кип",
			En: "Laotian Kip",
		},
	},
	Currency("LBP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Ливанский фунт",
			En: "Lebanese Pound",
		},
	},
	Currency("LKR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Шри-ланкийская рупия",
			En: "Sri Lankan Rupee",
		},
	},
	Currency("LRD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Либерийский доллар",
			En: "Liberian Dollar",
		},
	},
	Currency("LSL"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Лоти Лесото",
			En: "Lesotho Loti",
		},
	},
	Currency("LTL"): {
		Enabled: false,
		HasImg:  true,
		Translation: Tr{
			Ru: "Литовский лит",
			En: "Lithuanian Litas",
		},
	},
	Currency("LVL"): {
		Enabled: false,
		HasImg:  true,
		Translation: Tr{
			Ru: "Латвийский лат",
			En: "Latvian Lats",
		},
	},
	Currency("LYD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Ливийский динар",
			En: "Libyan Dinar",
		},
	},
	Currency("MAD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Марокканский дирхам",
			En: "Moroccan Dirham",
		},
	},
	Currency("MDL"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Молдавский лей",
			En: "Moldovan Leu",
		},
	},
	Currency("MGA"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Малагасийский ариари",
			En: "Malagasy Ariary",
		},
	},
	Currency("MKD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Македонский денар",
			En: "Macedonian Denar",
		},
	},
	Currency("MMK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Мьянманский кьят",
			En: "Myanma Kyat",
		},
	},
	Currency("MNT"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Монгольский тугрик",
			En: "Mongolian Tugrik",
		},
	},
	Currency("MOP"): {
		Enabled: false,
		HasImg:  true,
		Translation: Tr{
			Ru: "Патака Макао",
			En: "Macanese Pataca",
		},
	},
	Currency("MRO"): {
		Enabled: false,
		HasImg:  true,
		Translation: Tr{
			Ru: "Мавританская угия",
			En: "Mauritanian Ouguiya",
		},
	},
	Currency("MUR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Маврикийская рупия",
			En: "Mauritian Rupee",
		},
	},
	Currency("MVR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Мальдивская руфия",
			En: "Maldivian Rufiyaa",
		},
	},
	Currency("MWK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Малавийская квача",
			En: "Malawian Kwacha",
		},
	},
	Currency("MXN"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Мексиканское песо",
			En: "Mexican Peso",
		},
	},
	Currency("MYR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Малайзийский ринггит",
			En: "Malaysian Ringgit",
		},
	},
	Currency("MZN"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Мозамбикский метикал",
			En: "Mozambican Metical",
		},
	},
	Currency("NAD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Доллар Намибии",
			En: "Namibian Dollar",
		},
	},
	Currency("NGN"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Нигерийская найра",
			En: "Nigerian Naira",
		},
	},
	Currency("NIO"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Никарагуанская кордоба",
			En: "Nicaraguan Cordoba",
		},
	},
	Currency("NOK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Норвежская крона",
			En: "Norwegian Krone",
		},
	},
	Currency("NPR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Непальская рупия",
			En: "Nepalese Rupee",
		},
	},
	Currency("NZD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Новозеландский доллар",
			En: "New Zealand Dollar",
		},
	},
	Currency("OMR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Оманский риал",
			En: "Omani Rial",
		},
	},
	Currency("PAB"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Панамский бальбоа",
			En: "Panamanian Balboa",
		},
	},
	Currency("PEN"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Перуанский новый соль",
			En: "Peruvian Nuevo Sol",
		},
	},
	Currency("PGK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Кина Папуа-Новой Гвинеи",
			En: "Papua New Guinean Kina",
		},
	},
	Currency("PHP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Филиппинское песо",
			En: "Philippine Peso",
		},
	},
	Currency("PKR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Пакистанская рупия",
			En: "Pakistani Rupee",
		},
	},
	Currency("PLN"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Польский злотый",
			En: "Polish Zloty",
		},
	},
	Currency("PYG"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Парагвайский гуарани",
			En: "Paraguayan Guarani",
		},
	},
	Currency("QAR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Катарский риал",
			En: "Qatari Rial",
		},
	},
	Currency("RON"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Румынский лей",
			En: "Romanian Leu",
		},
	},
	Currency("RSD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Сербский динар",
			En: "Serbian Dinar",
		},
	},
	Currency("RUB"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Российский рубль",
			En: "Russian Ruble",
		},
	},
	Currency("RWF"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Франк Руанды",
			En: "Rwandan Franc",
		},
	},
	Currency("SAR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Саудовский риял",
			En: "Saudi Riyal",
		},
	},
	Currency("SBD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Доллар Соломоновых Островов",
			En: "Solomon Islands Dollar",
		},
	},
	Currency("SCR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Сейшельская рупия",
			En: "Seychellois Rupee",
		},
	},
	Currency("SDG"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Суданский фунт",
			En: "Sudanese Pound",
		},
	},
	Currency("SEK"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Шведская крона",
			En: "Swedish Krona",
		},
	},
	Currency("SGD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Сингапурский доллар",
			En: "Singapore Dollar",
		},
	},
	Currency("SHP"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Фунт Святой Елены",
			En: "Saint Helena Pound",
		},
	},
	Currency("SLE"): {
		Enabled: false,
		HasImg:  false,
		Translation: Tr{
			Ru: "Сьерра-леонский леоне",
			En: "Sierra Leonean Leone",
		},
	},
	Currency("SLL"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Сьерра-леонский леоне",
			En: "Sierra Leonean Leone",
		},
	},
	Currency("SOS"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Сомалийский шиллинг",
			En: "Somali Shilling",
		},
	},
	Currency("SRD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Суринамский доллар",
			En: "Surinamese Dollar",
		},
	},
	Currency("SSP"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Южносуданский фунт",
			En: "South Sudanese Pound",
		},
	},
	Currency("STD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Добра Сан-томе и Принсипи (до 2018)",
			En: "Dobra",
		},
	},
	Currency("SVC"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Сальвадорский колон",
			En: "Salvadoran Colon",
		},
	},
	Currency("SYP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Сирийский фунт",
			En: "Syrian Pound",
		},
	},
	Currency("SZL"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Свазилендский лилангени",
			En: "Swazi Lilangeni",
		},
	},
	Currency("THB"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Тайский бат",
			En: "Thai Baht",
		},
	},
	Currency("TJS"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Таджикский сомони",
			En: "Tajikistani Somoni",
		},
	},
	Currency("TMT"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Туркменский манат",
			En: "Turkmenistani Manat",
		},
	},
	Currency("TND"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Тунисский динар",
			En: "Tunisian Dinar",
		},
	},
	Currency("TOP"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Тонганская Паанга",
			En: "Tongan Pa’anga",
		},
	},
	Currency("TRY"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Турецкая лира",
			En: "Turkish Lira",
		},
	},
	Currency("TTD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Доллар Тринидада и Тобаго",
			En: "Trinidad and Tobago Dollar",
		},
	},
	Currency("TWD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Новый тайваньский доллар",
			En: "New Taiwan Dolla",
		},
	},
	Currency("TZS"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Танзанийский шиллинг",
			En: "Tanzanian Shilling",
		},
	},
	Currency("UAH"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Украинская гривна",
			En: "Ukrainian Hryvnia",
		},
	},
	Currency("UGX"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Угандийский шиллинг",
			En: "Ugandan Shilling",
		},
	},
	Currency("USD"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Доллар США",
			En: "United States Dollar",
		},
	},
	Currency("UYU"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Уругвайское песо",
			En: "Uruguayan Peso",
		},
	},
	Currency("UZS"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Узбекский сум",
			En: "zbekistan Som",
		},
	},
	Currency("VEF"): {
		Enabled: false,
		HasImg:  false,
		Translation: Tr{
			Ru: "Боливар фуэрте",
			En: "Venezuelan Bolivar",
		},
	},
	Currency("VES"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Суверенный боливар",
			En: "Sovereign Bolivar",
		},
	},
	Currency("VND"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Вьетнамский донг",
			En: "Vietnamese Dong",
		},
	},
	Currency("VUV"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Вату Вануату",
			En: "Vanuatu Vatu",
		},
	},
	Currency("WST"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Самоанская тала",
			En: "Samoan Tala",
		},
	},
	Currency("XAF"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Франк КФА BEAC",
			En: "CFA Franc BEAC",
		},
	},
	Currency("XAG"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Унция серебра",
			En: "Silver (troy ounce)",
		},
	},
	Currency("XAU"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Унция золота",
			En: "Gold (troy ounce)",
		},
	},
	Currency("XCD"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Восточнокарибский доллар",
			En: "East Caribbean Dollar",
		},
	},
	Currency("XDR"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Специальные права заимствования",
			En: "Special Drawing Rights",
		},
	},
	Currency("XOF"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Франк КФА BCEAO",
			En: "CFA Franc BCEAO",
		},
	},
	Currency("XPF"): {
		Enabled: false,
		HasImg:  false,
		Translation: Tr{
			Ru: "Тихоокеанский франк",
			En: "CFP Franc",
		},
	},
	Currency("YER"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Йеменский риал",
			En: "Yemeni Rial",
		},
	},
	Currency("ZAR"): {
		Enabled: true,
		HasImg:  true,
		Translation: Tr{
			Ru: "Южноафриканский рэнд",
			En: "South African Rand",
		},
	},
	Currency("ZMK"): {
		Enabled: false,
		HasImg:  true,
		Translation: Tr{
			Ru: "Замбийская квача",
			En: "Zambian Kwacha (pre-2013)",
		},
	},
	Currency("ZMW"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Замбийская квача",
			En: "Zambian Kwacha",
		},
	},
	Currency("ZWL"): {
		Enabled: true,
		HasImg:  false,
		Translation: Tr{
			Ru: "Доллар Зимбабве",
			En: "Zimbabwean Dollar",
		},
	},
}
