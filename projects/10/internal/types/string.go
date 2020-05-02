package types

func isString(s string) bool {
	if int32(s[0]) != '"' || int32(s[len(s)-1]) != '"' {
		return false
	}
	return true
}

func GetString(token string) string {
	return token[1 : len(token)-1]
}
