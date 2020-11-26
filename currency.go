package tripismapiutilities

import (
	"strings"
)

// CurrencySymbolFromCurrencyISO returns a currency symbol for a given currency ISO3 value
func CurrencySymbolFromCurrencyISO(currencyISO string) string {
	if len(currencyISO) == 0 {
		return ""
	}
	switch strings.ToUpper(currencyISO) {
	case "AED":
		return "د.إ"
	case "AFN":
		return "؋"
	case "ALL":
		return "Lek"
	case "AMD":
		return "֏"
	case "ANG":
		return "ƒ"
	case "AOA":
		return "Kz"
	case "ARS":
		return "$"
	case "AUD":
		return "$"
	case "AWG":
		return "Afl."
	case "AZN":
		return "₼"
	case "BAM":
		return "KM"
	case "BBD":
		return "$"
	case "BDT":
		return "৳"
	case "BGN":
		return "лв."
	case "BHD":
		return ".د.ب"
	case "BIF":
		return "FBu"
	case "BMD":
		return "$"
	case "BND":
		return "$"
	case "BOB":
		return "Bs"
	case "BRL":
		return "R$"
	case "BSD":
		return "$"
	case "BTN":
		return "Nu"
	case "BWP":
		return "P"
	case "BYN":
		return "Br"
	case "BZD":
		return "BZ$"
	case "CAD":
		return "$"
	case "CDF":
		return "FC"
	case "CHF":
		return "CHF"
	case "CLP":
		return "$"
	case "CNY":
		return "¥"
	case "COP":
		return "$"
	case "CRC":
		return "₡"
	case "CUC":
		return "CUC$"
	case "CUP":
		return "₱"
	case "CVE":
		return "$"
	case "CZK":
		return "Kč"
	case "DJF":
		return "Fdj"
	case "DKK":
		return "kr."
	case "DOP":
		return "RD$"
	case "DZD":
		return "DA"
	case "EGP":
		return "£"
	case "ERN":
		return "Nkf"
	case "ETB":
		return "Br"
	case "EUR":
		return "€"
	case "FJD":
		return "$"
	case "FKP":
		return "£"
	case "GBP":
		return "£"
	case "GEL":
		return "₾"
	case "GHS":
		return "GH₵"
	case "GIP":
		return "£"
	case "GMD":
		return "D"
	case "GNF":
		return "FG"
	case "GTQ":
		return "Q"
	case "GYD":
		return "$"
	case "HKD":
		return "$"
	case "HNL":
		return "L"
	case "HRK":
		return "kn"
	case "HTG":
		return "G"
	case "HUF":
		return "Ft"
	case "IDR":
		return "Rp"
	case "ILS":
		return "₪"
	case "INR":
		return "₹"
	case "IQD":
		return "د.ع"
	case "IRR":
		return "﷼"
	case "ISK":
		return "kr"
	case "JMD":
		return "J$"
	case "JOD":
		return "د.أ"
	case "JPY":
		return "¥"
	case "KES":
		return "KSh"
	case "KGS":
		return "лв"
	case "KHR":
		return "៛"
	case "KMF":
		return "CF"
	case "KPW":
		return "₩"
	case "KRW":
		return "₩"
	case "KWD":
		return "ك"
	case "KYD":
		return "$"
	case "KZT":
		return "₸"
	case "LAK":
		return "₭"
	case "LBP":
		return "ل.ل"
	case "LKR":
		return "₨"
	case "LRD":
		return "$"
	case "LSL":
		return "L"
	case "LYD":
		return "LD"
	case "MAD":
		return "DH"
	case "MDL":
		return "L"
	case "MGA":
		return "Ar"
	case "MKD":
		return "ден"
	case "MMK":
		return "K"
	case "MNT":
		return "₮"
	case "MOP":
		return "MOP$"
	case "MRO":
		return "UM"
	case "MUR":
		return "₨"
	case "MVR":
		return "Rf"
	case "MWK":
		return "MK"
	case "MXN":
		return "$"
	case "MYR":
		return "RM"
	case "MZN":
		return "MT"
	case "NAD":
		return "$"
	case "NGN":
		return "₦"
	case "NIO":
		return "C$"
	case "NOK":
		return "kr"
	case "NPR":
		return "₨"
	case "NZD":
		return "$"
	case "OMR":
		return "﷼"
	case "PAB":
		return "B/."
	case "PEN":
		return "S/."
	case "PGK":
		return "K"
	case "PHP":
		return "₱"
	case "PKR":
		return "₨"
	case "PLN":
		return "zł"
	case "PYG":
		return "Gs"
	case "QAR":
		return "﷼"
	case "RON":
		return "lei"
	case "RSD":
		return "РСД"
	case "RUB":
		return "₽"
	case "RWF":
		return "FRw"
	case "SAR":
		return "﷼"
	case "SBD":
		return "$"
	case "SCR":
		return "₨"
	case "SDG":
		return "£SD"
	case "SEK":
		return "kr"
	case "SGD":
		return "$"
	case "SHP":
		return "£"
	case "SLL":
		return "Le"
	case "SOS":
		return "S"
	case "SRD":
		return "$"
	case "SSP":
		return "SS£"
	case "STN":
		return "Db"
	case "SVC":
		return "$"
	case "SYP":
		return "£"
	case "SZL":
		return "L"
	case "THB":
		return "฿"
	case "TJS":
		return "cомонӣ"
	case "TMT":
		return "T"
	case "TND":
		return "DT"
	case "TOP":
		return "T$"
	case "TRY":
		return "₺"
	case "TTD":
		return "TT$"
	case "TWD":
		return "NT$"
	case "TZS":
		return "TSh"
	case "UAH":
		return "₴"
	case "UGX":
		return "USh"
	case "USD":
		return "$"
	case "UYU":
		return "$U"
	case "UZS":
		return "лв"
	case "VES":
		return "Bs. S"
	case "VND":
		return "₫"
	case "VUV":
		return "VT"
	case "WST":
		return "$"
	case "XAF":
		return "FCFA"
	case "XCD":
		return "$"
	case "XOF":
		return "CFA"
	case "XPF":
		return "F"
	case "YER":
		return "﷼"
	case "ZAR":
		return "R"
	case "ZMW":
		return "ZK"
	case "ZWL":
		return "$"
	}

	return strings.ToUpper(currencyISO)
}

// CurrencyDescriptionFromCurrencyISO returns a currency description for a given currency ISO3 value
func CurrencyDescriptionFromCurrencyISO(currencyISO string) string {
	if len(currencyISO) == 0 {
		return ""
	}
	switch strings.ToUpper(currencyISO) {
	case "AED":
		return "UAE Dirham"
	case "AFN":
		return "Afghani"
	case "ALL":
		return "Lek"
	case "AMD":
		return "Dram"
	case "ANG":
		return "Guilder"
	case "AOA":
		return "Kwanza"
	case "ARS":
		return "Argentine Peso"
	case "AUD":
		return "Australian Dollar"
	case "AWG":
		return "Aruban Florin"
	case "AZN":
		return "Azerbaijanian Manat"
	case "BAM":
		return "Convertible Mark"
	case "BBD":
		return "Barbados Dollar"
	case "BDT":
		return "Taka"
	case "BGN":
		return "Lev"
	case "BHD":
		return "Bahraini Dinar"
	case "BIF":
		return "Burundian Franc"
	case "BMD":
		return "Bermudian Dollar"
	case "BND":
		return "Brunei Dollar"
	case "BOB":
		return "Boliviano"
	case "BRL":
		return "Brazilian Real"
	case "BSD":
		return "Bahamian Dollar"
	case "BTN":
		return "Bhutanese Ngultrum"
	case "BWP":
		return "Botswana Pula"
	case "BYN":
		return "Belarusian Ruble"
	case "BZD":
		return "Belize Dollar"
	case "CAD":
		return "Canadian Dollar"
	case "CDF":
		return "Congolese Franc"
	case "CHF":
		return "Swiss Franc"
	case "CLP":
		return "Chilean Peso"
	case "CNY":
		return "Chinese Yuan"
	case "COP":
		return "Colombian Peso"
	case "CRC":
		return "Costa Rican Colon"
	case "CUC":
		return "Cuban Convertible Peso"
	case "CUP":
		return "Cuban Peso"
	case "CVE":
		return "Cape Verde Escudo"
	case "CZK":
		return "Czech Koruna"
	case "DJF":
		return "Djiboutian Franc"
	case "DKK":
		return "Danish Krone"
	case "DOP":
		return "Dominican Peso"
	case "DZD":
		return "Algerian Dinar"
	case "EGP":
		return "Egyptian Pound"
	case "ERN":
		return "Eritrean Nakfa"
	case "ETB":
		return "Ethiopian Birr"
	case "EUR":
		return "Euro"
	case "FJD":
		return "Fiji Dollar"
	case "FKP":
		return "Falkland Islands Pound"
	case "GBP":
		return "Pound Sterling"
	case "GEL":
		return "Georgian Lari"
	case "GHS":
		return "Ghanaian Cedi"
	case "GIP":
		return "Gibraltar Pound"
	case "GMD":
		return "Gambian Dalasi"
	case "GNF":
		return "Guinean Franc"
	case "GTQ":
		return "Guatemalan Quetzal"
	case "GYD":
		return "Guyanese Dollar"
	case "HKD":
		return "Hong Kong Dollar"
	case "HNL":
		return "Honduran Lempira"
	case "HRK":
		return "Croatian Kuna"
	case "HTG":
		return "Haitian Gourde"
	case "HUF":
		return "Hungarian Forint"
	case "IDR":
		return "Indonesian Rupiah"
	case "ILS":
		return "Israeli Shekel"
	case "INR":
		return "Indian Rupee"
	case "IQD":
		return "Iraqi Dinar"
	case "IRR":
		return "Iranian Rial"
	case "ISK":
		return "Icelandic Króna"
	case "JMD":
		return "Jamaican Dollar"
	case "JOD":
		return "Jordanian Dinar"
	case "JPY":
		return "Japanese Yen"
	case "KES":
		return "Kenyan Shilling"
	case "KGS":
		return "Kyrgyzstani Som"
	case "KHR":
		return "Cambodian Riel"
	case "KMF":
		return "Comorian Franc"
	case "KPW":
		return "North Korean Won"
	case "KRW":
		return "South Korean Won"
	case "KWD":
		return "Kuwaiti Dinar"
	case "KYD":
		return "Cayman Islands Dollar"
	case "KZT":
		return "Kazakhstani Tenge"
	case "LAK":
		return "Lao Kip"
	case "LBP":
		return "Lebanese Pound"
	case "LKR":
		return "Sri Lankan Rupee"
	case "LRD":
		return "Liberian Dollar"
	case "LSL":
		return "Lesotho Loti"
	case "LYD":
		return "Libyan Dinar"
	case "MAD":
		return "Moroccan Dirham"
	case "MDL":
		return "Moldovan Leu"
	case "MGA":
		return "Malagasy Ariary"
	case "MKD":
		return "Macedonian Denar"
	case "MMK":
		return "Myanmar Kyat"
	case "MNT":
		return "Mongolian Tughrik"
	case "MOP":
		return "Macanese Pataca"
	case "MRO":
		return "Mauritanian Ouguiya"
	case "MUR":
		return "Mauritian Rupee"
	case "MVR":
		return "Maldivian Rufiyaa"
	case "MWK":
		return "Malawian Kwacha"
	case "MXN":
		return "Mexican Peso"
	case "MYR":
		return "Malaysian Ringgit"
	case "MZN":
		return "Mozambican Metical"
	case "NAD":
		return "Namibian Dollar"
	case "NGN":
		return "Nigerian Naira"
	case "NIO":
		return "Nicaraguan Córdoba"
	case "NOK":
		return "Norwegian Krone"
	case "NPR":
		return "Nepalese Rupee"
	case "NZD":
		return "New Zealand Dollar"
	case "OMR":
		return "Omani Rial"
	case "PAB":
		return "Panamanian Balboa"
	case "PEN":
		return "Peruvian Sol"
	case "PGK":
		return "Papua New Guinean Kina"
	case "PHP":
		return "Philippine Peso"
	case "PKR":
		return "Pakistani Rupee"
	case "PLN":
		return "Polish Złoty"
	case "PYG":
		return "Paraguayan Guaraní"
	case "QAR":
		return "Qatari Rial"
	case "RON":
		return "Romanian Leu"
	case "RSD":
		return "Serbian Dinar"
	case "RUB":
		return "Russian Ruble"
	case "RWF":
		return "Rwandan Franc"
	case "SAR":
		return "Saudi Riyal"
	case "SBD":
		return "Solomon Islands Dollar"
	case "SCR":
		return "Seychelles Rupee"
	case "SDG":
		return "Sudanese Pound"
	case "SEK":
		return "Swedish Krona"
	case "SGD":
		return "Singapore Dollar"
	case "SHP":
		return "Saint Helena Pound"
	case "SLL":
		return "Sierra Leonean Leone"
	case "SOS":
		return "Somali Shilling"
	case "SRD":
		return "Surinamese Dollar"
	case "SSP":
		return "South Sudanese Pound"
	case "STN":
		return "São Toméan Dobra"
	case "SVC":
		return "Salvadoran Colón"
	case "SYP":
		return "Syrian Pound"
	case "SZL":
		return "Swazi Lilangeni"
	case "THB":
		return "Thai Baht"
	case "TJS":
		return "Tajikistani Somoni"
	case "TMT":
		return "Turkmenistan Manat"
	case "TND":
		return "Tunisian Dinar"
	case "TOP":
		return "Tongan Pa'anga"
	case "TRY":
		return "Turkish Lira"
	case "TTD":
		return "Trinidad and Tobago Dollar"
	case "TWD":
		return "New Taiwan Dollar"
	case "TZS":
		return "Tanzanian Shilling"
	case "UAH":
		return "Ukrainian Hryvnia"
	case "UGX":
		return "Ugandan Shilling"
	case "USD":
		return "United States Dollar"
	case "UYU":
		return "Uruguayan Peso"
	case "UZS":
		return "Uzbekistan Som"
	case "VEF":
		return "Venezuelan Bolívar"
	case "VES":
		return "Venezuelan Bolívar Soberano"
	case "VND":
		return "Vietnamese Đồng"
	case "VUV":
		return "Vanuatu Vatu"
	case "WST":
		return "Samoan Tala"
	case "XAF":
		return "CFA Franc BEAC"
	case "XCD":
		return "East Caribbean Dollar"
	case "XOF":
		return "CFA Franc BCEAO"
	case "XPF":
		return "CFP Franc (Franc Pacifique)"
	case "YER":
		return "Yemeni Rial"
	case "ZAR":
		return "South African Rand"
	case "ZMW":
		return "Zambian Kwacha"
	case "ZWL":
		return "Zimbabwean Dollar"
	}

	return strings.ToUpper(currencyISO)
}
