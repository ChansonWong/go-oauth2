package utils

import "time"

func CurrentTime() int64 {
	return time.Now().Unix()
}

func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

func CurrentTimeNano() int64{
	return time.Now().UnixNano()
}
