package memontago

import (
	"fmt"
	"log"
)

// 英语字典
type langEn struct {
	Ago     string
	Online  string
	Later   string
	JustNow string
	Second  string
	Seconds string
	Minute  string
	Minutes string
	Hour    string
	Hours   string
	Day     string
	Days    string
	Week    string
	Weeks   string
	Month   string
	Months  string
	Year    string
	Years   string
}

// 汉语字典
type langCh struct {
	Ago     string
	Online  string
	Later   string
	JustNow string
	Second  string
	Minute  string
	Hour    string
	Day     string
	Week    string
	Month   string
	Year    string
}

var langs = map[string]interface{}{} // 语言映射
func trans(lang string) {
	//	设置语言
	//	根据lang找到文件中对应的json
	filePath := fmt.Sprintf("./langs/%s.json", lang)
	isExits, err := fileExists(filePath)
	if !isExits {
		log.Fatalf("%s , doesn't exist\n", filePath)
	}
	if err != nil {
		log.Fatalf("Error while trying to rea file %s, Error: %v\n", filePath, err)
	}

	//	解析为struct
	langInterface := parseJsonIntoLang(filePath, lang)

	//	加入langs
	switch lang {
	case "ch":
		langs["ch"] = langInterface.(langCh)
	case "en":
		langs["en"] = langInterface.(langEn)
	}

}
func initChTrans() map[string]string {
	return map[string]string{
		"ago":     langs["ch"].(langCh).Ago,
		"later":   langs["ch"].(langCh).Later,
		"seconds": langs["ch"].(langCh).Second,
		"minutes": langs["ch"].(langCh).Minute,
		"hours":   langs["ch"].(langCh).Hour,
		"days":    langs["ch"].(langCh).Day,
		"weeks":   langs["ch"].(langCh).Week,
		"months":  langs["ch"].(langCh).Month,
		"years":   langs["ch"].(langCh).Year,
	}
}
func initEnTrans() map[string]string {
	return map[string]string{
		"ago":     langs["en"].(langEn).Ago,
		"later":   langs["en"].(langEn).Later,
		"second":  langs["en"].(langEn).Second,
		"seconds": langs["en"].(langEn).Seconds,
		"minute":  langs["en"].(langEn).Minute,
		"minutes": langs["en"].(langEn).Minutes,
		"hour":    langs["en"].(langEn).Hour,
		"hours":   langs["en"].(langEn).Hours,
		"day":     langs["en"].(langEn).Day,
		"days":    langs["en"].(langEn).Days,
		"week":    langs["en"].(langEn).Week,
		"weeks":   langs["en"].(langEn).Weeks,
		"month":   langs["en"].(langEn).Month,
		"months":  langs["en"].(langEn).Months,
		"year":    langs["en"].(langEn).Year,
		"years":   langs["en"].(langEn).Years,
	}
}

var ChTrans map[string]string

var EnTrans map[string]string
