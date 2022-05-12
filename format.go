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
	inputTime := Datetime2Time(datetime)
	//切割字符串
	s := strings.Split(format, " ")
	temp := ""
	//判断map中是否存在
	for i := 0; i < len(s); i++ {
		if _, ok := formatMap[s[i]]; ok {
			if !ok {
				log.Fatal("输入格式不对")
				return ""
			}
		}
		if formatMap[s[i]] == "01" {
			temp += inputTime.Format(monthMap[formatMap[s[i]]]) + " "
		} else if formatMap[s[i]] == "02" {
			temp += inputTime.Format(dayMap[formatMap[s[i]]]) + " "
		} else {
			temp += inputTime.Format(formatMap[s[i]])
		}
	}
	return temp
}
