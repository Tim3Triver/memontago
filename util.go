package memontago

import (
	"fmt"
	"time"
)

// datetime2Time 传入一个	接口类型的时间，返回time.Time类型
// 只接受 string 标准格式、 int time.Time  参数
// 如果不符合格式返回nil
func datetime2Time(datetime interface{}, format string) (time.Time, error) {
	// 返回一个time.Time
	var input time.Time

	//	转换
	switch date := datetime.(type) {
	case int:
		input = time.Unix(int64(date), 0)
	case string:
		isNilInput, err := stringTimeToTime(date, format)
		if err != nil {
			return time.Time{}, err
		}
		input = *isNilInput
	case int64:
		input = time.Unix(date, 0)
	default:
		input = date.(time.Time)
	}
	return input, nil
}

// stringTimeToTime 字符串格式转time.Time
func stringTimeToTime(date, format string) (*time.Time, error) {
	var input time.Time
	var err error
	if config.Location == "" {
		if format == "" {
			input, err = time.Parse("2006-01-02 15:04:05", date)
		} else {
			input, err = time.Parse(format, date)
		}
		if err != nil {
			return nil, err
		}
	}

	var location *time.Location
	location, err = time.LoadLocation(config.Location)
	if err != nil {
		return nil, fmt.Errorf("时区错误")
	}
	if format == "" {
		input, err = time.ParseInLocation("2006-01-02 15:04:05", date, location)
	} else {
		input, err = time.ParseInLocation(format, date, location)
	}
	if err != nil {
		return nil, err
	}
	return &input, nil
}
