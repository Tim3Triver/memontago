package memontago

import (
	"log"
	"strings"
)

var formatMap = map[string]string{
	"MMM":  "01",
	"Do":   "02",
	"YY":   "06",
	"MMMM": "01",
	"YYYY": "2006",
	"h":    "15",
	"mm":   "04",
	"ss":   "05",
	"a":    "下午",
}

// Format 格式化时间 format:"MMM Do YY"
func Format(datetime interface{}, format string) string {
	inputTime := datetime2Time(datetime)
	//切割字符串
	patterns := strings.Split(format, " ")
	temp := ""
	//判断map中是否存在
	for i := 0; i < len(patterns); i++ {
		pattern := patterns[i]
		if _, ok := formatMap[pattern]; !ok {
			log.Fatal("输入格式不对")
		}
		fmat := formatMap[pattern]
		switch fmat {
		case "01":
			temp += inputTime.Format(monthMap[fmat]) + " "
		case "02":
			temp += inputTime.Format(dayMap[fmat]) + " "
		default:
			temp += inputTime.Format(formatMap[fmat])
		}
	}
	return temp
}
