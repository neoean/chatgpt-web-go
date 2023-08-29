package types

import "strconv"

func InterfaceToInt(i interface{}) int {
	switch i.(type) {
	case int:
		return i.(int)
	case string:
		iStr := i.(string)
		iInt, _ := strconv.Atoi(iStr)
		return iInt
	case int64:
		return int(i.(int64))
	default:
		return 0
	}
}

func InterfaceToInt64(i interface{}) int64 {
	switch i.(type) {
	case int:
		return i.(int64)
	case string:
		iStr := i.(string)

		iInt, _ := strconv.ParseInt(iStr, 10, 64)
		return iInt
	case int64:
		return i.(int64)
	default:
		return 0
	}
}
