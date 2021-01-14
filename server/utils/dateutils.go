package utils

import "time"

func CurrentTime() int64 {
	return time.Now().Unix()
}

func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

func CurrentTimeNano() int64 {
	return time.Now().UnixNano()
}

func DayToSecond(days time.Duration) int {
	seconds := days.Seconds()
	return int(seconds)
}

func DayDuration(days time.Duration) time.Duration {
	return days * 24 * time.Hour
}

func HourDuration(hours time.Duration) time.Duration {
	return hours * time.Hour
}

func MinuteDuration(minutes time.Duration) time.Duration {
	return minutes * time.Minute
}

func SecondDuration(seconds time.Duration) time.Duration {
	return seconds * time.Second
}

func NextDaysSecond(days time.Duration) int {
	now := time.Now()
	dayDuration, _ := time.ParseDuration(days.String() + "d")
	time := now.Add(dayDuration)
	return time.Second()
}
