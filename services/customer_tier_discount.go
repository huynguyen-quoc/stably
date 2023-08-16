package services

type discountFunc func(network string) string

var tierDiscountFiat = map[string]discountFunc{
	"1": func(network string) string {
		return "0"
	},
	"2": func(network string) string {
		return "25"
	},
	"3": func(network string) string {
		if network == "Wire" {
			return "25"
		}
		return "100"
	},
}

var tierDiscountCrypto = map[string]discountFunc{
	"1": func(network string) string {
		return "0"
	},
	"2": func(network string) string {
		return "25"
	},
	"3": func(network string) string {
		if network == "ethereum" {
			return "25"
		}
		return "50"
	},
}
