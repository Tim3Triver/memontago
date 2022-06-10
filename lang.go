package memontago

var ChTrans map[string]string
var EnTrans map[string]string

func initChTrans() map[string]string {
	return map[string]string{
		"justNow":   "刚刚",
		"online":    "在线",
		"ago":       "前",
		"later":     "后",
		"seconds":   "秒",
		"minutes":   "分钟",
		"hours":     "小时",
		"days":      "天",
		"weeks":     "周",
		"months":    "月",
		"years":     "年",
		"Monday":    "星期一",
		"Tuesday":   "星期二",
		"Wednesday": "星期三",
		"Thursday":  "星期四",
		"Friday":    "星期五",
		"Saturday":  "星期六",
		"Sunday":    "星期日"}
} // 在配置的时候调用
func initEnTrans() map[string]string {
	return map[string]string{
		"justNow": "Just now",
		"online":  "Online",
		"ago":     "ago",
		"later":   "later",
		"second":  "second",
		"seconds": "seconds",
		"minute":  "minute",
		"minutes": "minutes",
		"hour":    "hour",
		"hours":   "hours",
		"day":     "day",
		"days":    "days",
		"week":    "week",
		"weeks":   "weeks",
		"month":   "month",
		"months":  "months",
		"year":    "year",
		"years":   "years",
	}
} // 在默认初始化配置的时候调用
