package utils

import "strconv"

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func IntToString(num int) string {
	return strconv.Itoa(num)
}
