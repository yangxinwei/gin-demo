package utils

import "time"

var (
	TimeLocation, _ = time.LoadLocation("Asia/Chongqing") //当地时间
	TimeLocationUS, _ = time.LoadLocation("US/Eastern") //美国东部时间
)

func GetUnixTimestamp() int64 {
	return time.Now().Unix()
}

func GetTime()  {
	
}