package memontago

import (
	"strconv"
	"time"
)

var gobalOptions = []string{}

// Parse 解析时间，并输出与当前时间的关系
// 如：1 second ago
func Parse(datetime interface{}, options ...string) string {
	//转化
	input := datetime2Time(datetime)
	//	添加options
	gobalOptions = append(gobalOptions, options...)

	//计算秒数
	second := time.Now().Sub(input).Seconds()
	//将来的时间
	if second < 0 {
		// 将来的时间 打一个标识符 upcoming
		gobalOptions = append(gobalOptions, "upcoming")
		second = -second
	}
	//	统计结果 kind number
	kindtime, number := calculateTheResult(int(second))
	//	result
	return getWords(kindtime, number)
}

func getWords(kindtime string, number int) string {
	switch config.Language {
	case "ch":
		if optionIsEnabled("upcoming") { // 未来的时间
			return strconv.Itoa(number) + ChTrans[kindtime] + ChTrans["later"]
		}
		// 过去的时间
		if kindtime == "seconds" && number <= 30 && optionIsEnabled("online") {
			return ChTrans["online"]

		} else if kindtime == "seconds" && number > 30 && number <= 60 && optionIsEnabled("justNow") {
			return ChTrans["justNow"]
		}
		return strconv.Itoa(number) + ChTrans[kindtime] + ChTrans["ago"]

	case "en":
		// 数量为单数
		if number == 1 {
			kindtime = kindtime[:len(kindtime)-1]
		}
		if optionIsEnabled("upcoming") { // 未来的时间
			return strconv.Itoa(number) + " " + EnTrans[kindtime] + " " + EnTrans["later"]
		}
		// 过去的时间
		if kindtime == "second" || kindtime == "seconds" && number <= 30 && optionIsEnabled("online") {
			return EnTrans["online"]
		} else if kindtime == "seconds" && number > 30 && number <= 60 && optionIsEnabled("justNow") {
			return EnTrans["justNow"]
		}
		return strconv.Itoa(number) + " " + EnTrans[kindtime] + " " + EnTrans["ago"]
	}
	return ""
}
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

func optionIsEnabled(option string) bool {
	for _, options := range gobalOptions {
		if options == option {
			return true
		}
	}
	return false
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
