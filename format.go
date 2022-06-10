package memontago

import (
	"fmt"
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
func Format(datetime interface{}, format string) (string, error) {
	inputTime, err := datetime2Time(datetime)
	if err != nil {
		return "", err
	}
	//切割字符串
	temp := ""
	if format == "" {
		temp += inputTime.Format("2006-01-02 15:04:05.000")
	}
	//切割字符串
	s := strings.Split(format, " ")
	//判断map中是否存在
	for i := 0; i < len(s); i++ {
		_, ok := formatMap[s[i]]
		if formatMap[s[i]] != "" && !ok {
			return "", fmt.Errorf("输入格式不对")
		}
		switch formatMap[s[i]] {
		case "01": // 月
			temp += inputTime.Format("1月") + " "
		case "02": // 日
			temp += inputTime.Format("2日") + " "
		case "15": // 小时
			temp += inputTime.Format("03")
		case "04": // 分钟
			temp += inputTime.Format(":04")
		case "05": // 秒
			temp += inputTime.Format(":05")
		default: // 没匹配到
			temp += inputTime.Format(formatMap[s[i]]) + " "
		}
	}
	return temp, nil
}
