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

// trans 传入语言标识，将对应的json格式的语言文件的解析为结构体 并存入到langs 语言：结构体
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

// 语言映射
var langs = map[string]interface{}{}

var ChTrans map[string]string
var EnTrans map[string]string

func initChTrans() map[string]string {
	return map[string]string{
		"justNow": langs["ch"].(langCh).JustNow,
		"online":  langs["ch"].(langCh).Online,
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
} // 在配置的时候调用
func initEnTrans() map[string]string {
	return map[string]string{
		"justNow": langs["en"].(langEn).JustNow,
		"online":  langs["en"].(langEn).Online,
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
} // 在默认初始化配置的时候调用
