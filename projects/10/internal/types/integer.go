package types

import "strconv"

func IsIntegerConstant(s string) bool {
	parsedNum, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return false
	}
	// valid range 0 < parsedNum < 32767
	if parsedNum < 0 || parsedNum > 32767 {
		return false
	}
	return true
}

func GetIntegerConstant(token string) int32 {
	// TODO: error handling
	parsedNum, _ := strconv.ParseInt(token, 10, 32)
	return int32(parsedNum)
}
