package memontago

import "fmt"

func GetWord(kindTime string, number int, status string) string {
	var tmp string
	if config.Special {
		// 特殊标记
		if kindTime == "seconds" && number <= 30 {
			tmp = "justNow"
		}
		if kindTime == "seconds" && number <= 5 {
			tmp = "online"
		}
	}
	if number == 1 {
		kindTime = kindTime[:len(kindTime)-1]
	}
	if status == "upComing" {
		status = "later"
	}
	arr := [3]string{tmp, kindTime, status}
	for i := 0; i < 3; i++ {
		switch arr[i] {
		case "second":
			arr[i] = Lang.Second
		case "seconds":
			arr[i] = Lang.Seconds
		case "minute":
			arr[i] = Lang.Minute
		case "minutes":
			arr[i] = Lang.Minutes
		case "hour":
			arr[i] = Lang.Hour
		case "hours":
			arr[i] = Lang.Hours
		case "day":
			arr[i] = Lang.Day
		case "days":
			arr[i] = Lang.Days
		case "week":
			arr[i] = Lang.Week
		case "weeks":
			arr[i] = Lang.Weeks
		case "month":
			arr[i] = Lang.Month
		case "months":
			arr[i] = Lang.Months
		case "year":
			arr[i] = Lang.Year
		case "years":
			arr[i] = Lang.Years
		case "online":
			arr[i] = Lang.Online
		case "justNow":
			arr[i] = Lang.JustNow
		case "later":
			arr[i] = Lang.Later
		case "ago":
			arr[i] = Lang.Ago
		}

		if config.Special && i == 0 && arr[0] != "" {
			return arr[0]
		}
	}
	if config.Language == "ch" {
		return fmt.Sprintf("%d%s%s", number, arr[1], arr[2])
	}
	return fmt.Sprintf("%d %s %s", number, arr[1], arr[2])
}
