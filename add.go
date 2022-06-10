package memontago

import "time"

// Add 支持在给定时间上增加 number * DurUnit
func Add(dateTime interface{}, number int, DurUnit string) (time.Time, error) {
	input, err := datetime2Time(dateTime)
	if err != nil {
		return time.Time{}, err
	}
	//如果number是0 直接返回转化后的时间
	if number == 0 {
		return input, nil
	}
	durNumber := time.Duration(number)
	switch DurUnit {
	case "second":
		return input.Add(durNumber * time.Second), nil
	case "minute":
		return input.Add(durNumber * time.Minute), nil
	case "hour":
		return input.Add(durNumber * time.Hour), nil
	case "week":
		return input.Add(durNumber * time.Hour * 24 * 7), nil
	case "day":
		return input.AddDate(0, 0, number), nil
	case "month": // 一个月 默认按照 31 天计算
		return input.AddDate(0, number, 0), nil
	case "year":
		return input.AddDate(number, 0, 0), nil
	default:
		return input, nil
	}
}
