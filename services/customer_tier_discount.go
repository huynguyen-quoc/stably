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

/**
  @param network name (eg: ethereum)
  @return the percentage of network discount (eg: 25% -> "25")
*/
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
	"4": func(network string) string {
		return "100"
	},
}
