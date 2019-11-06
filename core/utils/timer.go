package utils

import "time"

func NowMillisecond() int {
	return int(time.Now().UnixNano() / 1000000)
}
