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
		return "فلس"
	case "AFN":
		return "؋"
	case "ALL":
		return "Lek"
	case "AMD":
		return "dram"
	case "ANG":
		return "ƒ"
	case "AOA":
		return "Kz"
	case "ARS":
		return "$"
	case "AUD":
		return "$"
	case "AWG":
		return "ƒ"
	case "AZN":
		return "ман"
	case "BAM":
		return "KM"
	case "BBD":
		return "$"
	case "BDT":
		return "৳"
	case "BGN":
		return "лв"
	case "BHD":
		return ".د.ب"
	case "BIF":
		return "FBu"
	case "BMD":
		return "$"
	case "BND":
		return "$"
	case "BOB":
		return "$b"
	case "BRL":
		return "R$"
	case "BSD":
		return "$"
	case "BTN":
		return "Nu"
	case "BWP":
		return "P"
	case "BYN":
		return "p."
	case "BYR":
		return "p."
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
		return "₱"
	case "CUP":
		return "₱"
	case "CVE":
		return "$"
	case "CZK":
		return "Kč"
	case "DJF":
		return "Fdj"
	case "DKK":
		return "kr"
	case "DOP":
		return "RD$"
	case "DZD":
		return "دج"
	case "EGP":
		return "£"
	case "ERN":
		return "Nfk"
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
		return "ع.د"
	case "IRR":
		return "IRR"
	case "ISK":
		return "kr"
	case "JMD":
		return "J$"
	case "JOD":
		return "ينار"
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
		return "K.D."
	case "KYD":
		return "$"
	case "KZT":
		return "лв"
	case "LAK":
		return "₭"
	case "LBP":
		return "£"
	case "LKR":
		return "Rs"
	case "LRD":
		return "$"
	case "LSL":
		return "L"
	case "LYD":
		return "LD"
	case "MAD":
		return ".م."
	case "MDL":
		return "MDL"
	case "MGA":
		return "Ar"
	case "MKD":
		return "ден"
	case "MMK":
		return "L"
	case "MNT":
		return "₮"
	case "MOP":
		return "MOP$"
	case "MRO":
		return "UM"
	case "MUR":
		return "Rs"
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
		return "N$"
	case "NGN":
		return "XXX"
	case "NIO":
		return "₦"
	case "NOK":
		return "kr"
	case "NPR":
		return "Rs"
	case "NZD":
		return "$"
	case "OMR":
		return "OMR"
	case "PAB":
		return "B/."
	case "PEN":
		return "S/."
	case "PGK":
		return "K"
	case "PHP":
		return "Php"
	case "PKR":
		return "Rs"
	case "PLN":
		return "zł"
	case "PYG":
		return "Gs"
	case "QAR":
		return "QAR"
	case "RON":
		return "lei"
	case "RSD":
		return "Дин."
	case "RUB":
		return "руб"
	case "RWF":
		return "FRw"
	case "SAR":
		return "SAR"
	case "SBD":
		return "$"
	case "SCR":
		return "Rs"
	case "SDG":
		return "£"
	case "SEK":
		return "kr"
	case "SGD":
		return "S$"
	case "SHP":
		return "£"
	case "SLL":
		return "Le"
	case "SOS":
		return "S"
	case "SRD":
		return "$"
	case "SSP":
		return "£"
	case "STD":
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
		return "TJS"
	case "TMT":
		return "m"
	case "TND":
		return "DT"
	case "TOP":
		return "T$"
	case "TRY":
		return "TL"
	case "TTD":
		return "TT$"
	case "TWD":
		return "NT$"
	case "TZS":
		return "shilingi"
	case "UAH":
		return "₴"
	case "UGX":
		return "Ush"
	case "USD":
		return "$"
	case "UYU":
		return "$U"
	case "UZS":
		return "лв"
	case "VEF":
		return "Bs"
	case "VND":
		return "₫"
	case "VUV":
		return "VT"
	case "WST":
		return "WS$"
	case "XAF":
		return "FCFA"
	case "XCD":
		return "$"
	case "XOF":
		return "CFA"
	case "XPF":
		return "F"
	case "YER":
		return "YER"
	case "ZAR":
		return "R"
	case "ZMW":
		return "ZMW"
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
		return "Armenian Dram"
	case "ANG":
		return "Netherlands Antillean Guilder"
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
		return "Bulgarian Lev"
	case "BHD":
		return "Bahraini Dinar"
	case "BIF":
		return "Burundi Franc"
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
		return "Ngultrum"
	case "BWP":
		return "Pula"
	case "BYN":
		return "Belarusian Ruble"
	case "BYR":
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
		return "Yuan Renminbi"
	case "COP":
		return "Colombian Peso"
	case "CRC":
		return "Costa Rican Colon"
	case "CUC":
		return "Peso Convertible"
	case "CUP":
		return "Cuban Peso"
	case "CVE":
		return "Cabo Verde Escudo"
	case "CZK":
		return "Czech Koruna"
	case "DJF":
		return "Djibouti Franc"
	case "DKK":
		return "Danish Krone"
	case "DOP":
		return "Dominican Peso"
	case "DZD":
		return "Algerian Dinar"
	case "EGP":
		return "Egyptian Pound"
	case "ERN":
		return "Nakfa"
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
		return "Lari"
	case "GHS":
		return "Ghana Cedi"
	case "GIP":
		return "Gibraltar Pound"
	case "GMD":
		return "Dalasi"
	case "GNF":
		return "Guinea Franc"
	case "GTQ":
		return "Quetzal"
	case "GYD":
		return "Guyana Dollar"
	case "HKD":
		return "Hong Kong Dollar"
	case "HNL":
		return "Lempira"
	case "HRK":
		return "Kuna"
	case "HTG":
		return "Gourde"
	case "HUF":
		return "Forint"
	case "IDR":
		return "Rupiah"
	case "ILS":
		return "New Israeli Sheqel"
	case "INR":
		return "Indian Rupee"
	case "IQD":
		return "Iraqi Dinar"
	case "IRR":
		return "Iranian Rial"
	case "ISK":
		return "Iceland Krona"
	case "JMD":
		return "Jamaican Dollar"
	case "JOD":
		return "Jordanian Dinar"
	case "JPY":
		return "Yen"
	case "KES":
		return "Kenyan Shilling"
	case "KGS":
		return "Som"
	case "KHR":
		return "Riel"
	case "KMF":
		return "Comoro Franc"
	case "KPW":
		return "North Korean Won"
	case "KRW":
		return "Won"
	case "KWD":
		return "Kuwaiti Dinar"
	case "KYD":
		return "Cayman Islands Dollar"
	case "KZT":
		return "Tenge"
	case "LAK":
		return "Kip"
	case "LBP":
		return "Lebanese Pound"
	case "LKR":
		return "Sri Lanka Rupee"
	case "LRD":
		return "Liberian Dollar"
	case "LSL":
		return "Loti"
	case "LYD":
		return "Libyan Dinar"
	case "MAD":
		return "Moroccan Dirham"
	case "MDL":
		return "Moldovan Leu"
	case "MGA":
		return "Malagasy Ariary"
	case "MKD":
		return "Denar"
	case "MMK":
		return "Kyat"
	case "MNT":
		return "Tugrik"
	case "MOP":
		return "Pataca"
	case "MRO":
		return "Ouguiya"
	case "MUR":
		return "Mauritius Rupee"
	case "MVR":
		return "Rufiyaa"
	case "MWK":
		return "Malawi Kwacha"
	case "MXN":
		return "Mexican Peso"
	case "MYR":
		return "Malaysian Ringgit"
	case "MZN":
		return "Mozambique Metical"
	case "NAD":
		return "Namibia Dollar"
	case "NGN":
		return "Naira"
	case "NIO":
		return "Cordoba Oro"
	case "NOK":
		return "Norwegian Krone"
	case "NPR":
		return "Nepalese Rupee"
	case "NZD":
		return "New Zealand Dollar"
	case "OMR":
		return "Rial Omani"
	case "PAB":
		return "Balboa"
	case "PEN":
		return "Sol"
	case "PGK":
		return "Kina"
	case "PHP":
		return "Philippine Peso"
	case "PKR":
		return "Pakistan Rupee"
	case "PLN":
		return "Zloty"
	case "PYG":
		return "Guarani"
	case "QAR":
		return "Qatari Rial"
	case "RON":
		return "Romanian Leu"
	case "RSD":
		return "Serbian Dinar"
	case "RUB":
		return "Russian Ruble"
	case "RWF":
		return "Rwanda Franc"
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
		return "Leone"
	case "SOS":
		return "Somali Shilling"
	case "SRD":
		return "Surinam Dollar"
	case "SSP":
		return "South Sudanese Pound"
	case "STD":
		return "Dobra"
	case "SVC":
		return "El Salvador Colon"
	case "SYP":
		return "Syrian Pound"
	case "SZL":
		return "Lilangeni"
	case "THB":
		return "Baht"
	case "TJS":
		return "Somoni"
	case "TMT":
		return "Turkmenistan New Manat"
	case "TND":
		return "Tunisian Dinar"
	case "TOP":
		return "Pa'anga"
	case "TRY":
		return "Turkish Lira"
	case "TTD":
		return "Trinidad and Tobago Dollar"
	case "TWD":
		return "New Taiwan Dollar"
	case "TZS":
		return "Tanzanian Shilling"
	case "UAH":
		return "Hryvnia"
	case "UGX":
		return "Uganda Shilling"
	case "USD":
		return "US Dollar"
	case "UYU":
		return "Peso Uruguayo"
	case "UZS":
		return "Uzbekistan Sum"
	case "VEF":
		return "Bolívar"
	case "VND":
		return "Dong"
	case "VUV":
		return "Vatu"
	case "WST":
		return "Tala"
	case "XAF":
		return "CFA Franc BEAC"
	case "XCD":
		return "East Caribbean Dollar"
	case "XOF":
		return "CFA Franc BCEAO"
	case "XPF":
		return "CFP Franc"
	case "YER":
		return "Yemeni Rial"
	case "ZAR":
		return "Rand"
	case "ZMW":
		return "Zambian Kwacha"
	case "ZWL":
		return "Zimbabwe Dollar"
	}

	return strings.ToUpper(currencyISO)
}
