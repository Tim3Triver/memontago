package memontago

import "time"

func Add(dateTime interface{}, number int, DurUnit string) time.Time {
	input := Datetime2Time(dateTime)
	//如果number是0 直接返回转化后的时间
	if number == 0 {
		return input
	}
	durNumber := time.Duration(number)
	switch DurUnit {
	case "second":
		return input.Add(durNumber * time.Second)
	case "minute":
		return input.Add(durNumber * time.Minute)
	case "hour":
		return input.Add(durNumber * time.Hour)
	case "week":
		return input.Add(durNumber * time.Hour * 24 * 7)
	case "day":
		return input.AddDate(0, 0, number)
	case "month": // 一个月 默认按照 30 天计算
		return input.AddDate(0, number, 0)
	case "year":
		return input.AddDate(number, 0, 0)
	default:
		return input
	}
}
