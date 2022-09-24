package memontago

import (
	"math"
	"time"
)

// Parse 解析时间，并输出与当前时间的关系
// 如：1 second ago
func Parse(datetime interface{}) (string, error) {
	// 标记
	status := "ago"

	// 转化
	input, err := datetime2Time(datetime)
	if err != nil {
		return "", err
	}

	// 计算秒数
	second := time.Now().Sub(input).Seconds()
	second = math.Round(second)
	// 将来的时间
	if second < 0 {
		status = "upComing" // 将来的时间 打一个标识符 upComing
		second = -second
	}
	//	统计结果 kind number
	kindTime, number := calculateTheResult(int(second))
	//	result
	return GetWord(kindTime, number, status), nil
}

// calculateTheResult 计算时间 以及 时间单位
func calculateTheResult(second int) (string, int) {
	minutes, hours, days, weeks, months, years := getTimeCalculations(second)
	switch {
	case second < 60:
		return "seconds", second
	case minutes < 60:
		return "minutes", minutes
	case hours < 24:
		return "hours", hours
	case days < 7:
		return "days", days
	case days <= 30: // 三十天以内
		return "weeks", weeks
	case months < 12:
		return "months", months
	default:
		return "years", years
	}
}

// getTimeCalculations 根据秒数转化时间
func getTimeCalculations(seconds int) (int, int, int, int, int, int) {
	minutes := seconds / 60
	hours := seconds / 3600
	days := seconds / 86400
	weeks := seconds / 604800
	months := seconds / 2629440
	years := seconds / 31553280

	return minutes, hours, days, weeks, months, years
}
