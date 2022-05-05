package utils

import (
	"memontago"
	"time"
)

// Datetime2Time 传入一个	接口类型的时间，返回time.Time类型
// 只接受 string 标准格式、 int time.Time  参数
// 如果不符合格式返回nil
func Datetime2Time(datetime interface{}) time.Time {
	// 返回一个time.Time
	var input time.Time

	//	转换
	switch date := datetime.(type) {
	case int:
		input = time.Unix(int64(date), 0)
		return input
	case string:
		input = stringtime2time(date)
	default:
		return datetime.(time.Time)
	}
	return input
}

// stringtime2time 字符串格式转time.Time
func stringtime2time(date string) time.Time {
	var input time.Time
	if memontago.Configx.Location == "" {
		parseTime, err := time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			input = parseTime
		}
	} else {
		location, err := time.LoadLocation(memontago.Configx.Location)
		if err != nil {
			panic("时区加载失败")
		}
		parseTime, err := time.ParseInLocation("2006-01-02 15:04:05", date, location)
		if err != nil {
			input = parseTime
		}
	}
	return input
}
