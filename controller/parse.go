package controller

import (
	"math"
	"memontago"
	"memontago/utils"
	"time"
)

var gobalOptions = []string{}

func ParseTime(datatime interface{}, options ...string) string {
	//转化
	input := utils.Datetime2Time(datatime)
	//	添加options
	gobalOptions = append(gobalOptions, options...)

	//计算秒数
	second := time.Now().Sub(input).Seconds()

	//	统计结果 kind number
	kindtime, number := calculateTheResult(int(second))
	//	result
	return getWords(kindtime, number)
}

func getWords(kindtime string, number int) string {
	//lang : kindtime : single or plural
	//	需要使用不同的语言  各种语言需要缓存下来,之后取的时候直接
	return ""
}
func calculateTheResult(second int) (string, int) {
	minutes, hours, days, weeks, months, years := getTimeCalculations(float64(second))
	switch {
	case second < 60:
		return "seconds", second
	case minutes < 60:
		return "minutes", minutes
	case hours < 24:
		return "hours", hours
	case days < 7:
		return "days", days
	case weeks < 5:
		return "weeks", weeks
	case months < 12:
		return "months", months
	default:
		return "years", years
	}
}

func optionIsEnabled(option string) bool {
	for _, options := range memontago.GlobalOptions {
		if options == option {
			return true
		}
	}
	return false
}

// getTimeCalculations 根据秒数转化时间
func getTimeCalculations(seconds float64) (int, int, int, int, int, int) {
	minutes := math.Round(seconds / 60)
	hours := math.Round(seconds / 3600)
	days := math.Round(seconds / 86400)
	weeks := math.Round(seconds / 604800)
	months := math.Round(seconds / 2629440)
	years := math.Round(seconds / 31553280)

	return int(minutes), int(hours), int(days), int(weeks), int(months), int(years)
}
