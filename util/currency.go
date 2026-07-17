package util

const (
	USD = "USD"
	EUR = "EUR"
	GHS = "GHS"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GHS:
		return true
	}
	return false
}
