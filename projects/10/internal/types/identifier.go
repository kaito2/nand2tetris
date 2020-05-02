package types

// Check s is valid identifier token
func IsIdentifier(s string) bool {
	if isDigit(int32(s[0])) {
		return false
	}
	for _, c := range s {
		if !isValidCharacter(c) {
			return false
		}
	}
	return true
}

func isValidCharacter(c int32) bool {
	if isAlphabet(c) {
		return true
	}
	if isDigit(c) {
		return true
	}
	if c == '_' {
		return true
	}
	return false
}

func isAlphabet(c int32) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c int32) bool {
	return c >= '0' && c <= '9'
}
