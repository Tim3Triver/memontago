package memontago

import (
	"fmt"
	"log"
)

// 设置的语言为英语
type langEn struct {
}
type langCh struct {
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
