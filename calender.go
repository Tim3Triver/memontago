package memontago

import (
	"fmt"
	"time"
)

// Calender 给定一个日期，返回给定的日历日期
func Calender(datetime interface{}) string {
	//datetime 转 time.Time
	inputTime := Datetime2Time(datetime)
	// 当天00:00:00的time
	now, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	//	second差
	diff := int(inputTime.Sub(now).Seconds())
	if diff < 0 {
		diff = diff / 86400
		diff--
	} else {
		diff = diff / 86400
	}
	r := ""
	switch {
	case diff == 0:
		r = "今天"
	case diff == 1:
		r = "明天"
	case diff == -1:
		r = "昨天"
	}
	if r != "" {
		return r + inputTime.Format("15:04")
	}
	weeks := weekday[fmt.Sprint(now.Weekday())]
	switch {
	case diff+weeks < 21 && diff+weeks > 13:
		r = "下"
	case diff+weeks < 14 && diff+weeks > 6:
		r = "这"
	case diff+weeks < 7 && diff+weeks >= 0:
		r = "上"
	default:
		return inputTime.Format("2006年01月02日 15:04")
	}
	return r + ChTrans[fmt.Sprint(inputTime.Weekday())] + inputTime.Format("15:04")
}
