package memontago

import (
	"log"
	"time"
)

// datetime2Time 传入一个	接口类型的时间，返回time.Time类型
// 只接受 string 标准格式、 int time.Time  参数
// 如果不符合格式返回nil
func datetime2Time(datetime interface{}) time.Time {
	// 返回一个time.Time
	var input time.Time
	//	转换
	switch date := datetime.(type) {
	case int:
		input = time.Unix(int64(date), 0)
	case string:
		input = stringtime2time(date)
	case int64:
		input = time.Unix(date, 0)
	default:
		input = date.(time.Time)
	}
	return input
}

// stringtime2time 字符串格式转time.Time
func stringtime2time(date string) time.Time {
	var input time.Time
	if config.Location == "" {
		parseTime, err := time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			log.Fatalln(err)
		}
		input = parseTime
	} else {
		location, err := time.LoadLocation(config.Location)
		if err != nil {
			panic("时区加载失败")
		}
		parseTime, err := time.ParseInLocation("2006-01-02 15:04:05", date, location)
		if err != nil {
			log.Fatalln(err)
		}
		input = parseTime
	}
	return input
}
