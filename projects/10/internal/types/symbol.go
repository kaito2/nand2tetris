package types

type Symbol string

var allSymbols = []Symbol{
	"{", "}", "(", ")", "[", "]", ".", ",", ";", "+", "-", "*", "/", "&", "|", "<", ">", "=", "~",
}

func IsSymbol(s string) bool {
	for _, symbol := range allSymbols {
		if Symbol(s) == symbol {
			return true
		}
	}
	return false
}
