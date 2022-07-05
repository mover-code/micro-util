package tools

import "time"

// 日期时间字符串转时间戳（秒）
func DateTime2Timestamp(datetime string) int64 {
	local, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, local)
	return tmp.Unix() //转化为时间戳 类型是int64

}

// 纯日期字符串转时间戳（秒）
func Date2Timestamp(datetime string) int64 {
	local, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02", datetime, local)
	return tmp.Unix() //转化为时间戳 类型是int64
}

// 时间戳(秒)转时间字符串
func Timestamp2Date(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
func Timestamp2min(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
func TimestampToDataWithDay(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02")
}
func TimesToNum(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("20060102")
}

// 秒转换为时分秒
func FormatSecond(seconds int64) (day, hour, minute, second int64) {
	day = seconds / (24 * 3600)
	hour = (seconds - day*3600*24) / 3600
	minute = (seconds - day*24*3600 - hour*3600) / 60
	second = seconds - day*24*3600 - hour*3600 - minute*60
	return day, hour, minute, second
}

//通过时间戳获取00:00:00和23.59.59时间戳
// 例data  : 2020-11-11
func GetDateTime(date string) (int64, int64) {
	//获取当前时区
	loc, _ := time.LoadLocation("Local")

	//日期当天0点时间戳(拼接字符串)
	startDate := date + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)

	//日期当天23时59分时间戳
	endDate := date + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	//返回当天0点和23点59分的时间戳
	return startTime.Unix(), end.Unix()
}

//通过时间戳获取00:00:00和23.59.59时间戳

func GetTodayDateTime() (int64, int64) {
	date := time.Now().Format("2006-01-02")
	//获取当前时区
	loc, _ := time.LoadLocation("Local")

	//日期当天0点时间戳(拼接字符串)
	startDate := date + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)

	//日期当天23时59分时间戳
	endDate := date + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	//返回当天0点和23点59分的时间戳
	return startTime.Unix(), end.Unix()
}

//获取昨天日期
//return：2021-01-01
func GetYesterdayDate() string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	logDay := yesTime.Format("2006-01-02")
	return logDay
}

//获取当天日期
//return：2021-01-01
func GetTodayDate() string {
	return time.Now().Format("2006-01-02")
}
