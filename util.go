package memontago

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// datetime2Time 传入一个	接口类型的时间，返回time.Time类型
// 只接受 string 标准格式、 int time.Time  参数
// 如果不符合格式返回nil
func datetime2Time(datetime interface{}) time.Time {
	// 返回一个time.Time
	var input time.Time
	//	转换
	switch date := datetime.(type) {
	case int:
		input = time.Unix(int64(date), 0)
	case string:
		input = stringtime2time(date)
	case int64:
		input = time.Unix(date, 0)
	default:
		input = date.(time.Time)
	}
	return input
}

// stringtime2time 字符串格式转time.Time
func stringtime2time(date string) time.Time {
	var input time.Time
	if config.Location == "" {
		parseTime, err := time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			log.Fatalln(err)
		}
		input = parseTime
	} else {
		location, err := time.LoadLocation(config.Location)
		if err != nil {
			panic("时区加载失败")
		}
		parseTime, err := time.ParseInLocation("2006-01-02 15:04:05", date, location)
		if err != nil {
			log.Fatalln(err)
		}
		input = parseTime
	}
	return input
}

func fileExists(filePath string) (bool, error) {
	//加载filePath文件属性
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	// 文件未找到
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	// 其他错误
	return false, err
}

// 解析json字典，为language字典
func parseJsonIntoLang(filePath, lang string) interface{} {
	// 读取文件内容
	context := getContent(filePath)
	switch lang {
	case "ch":
		var langch langCh
		err := json.Unmarshal(context, &langch) // 反序列化json
		if err != nil {
			log.Fatalln(err)
		}
		return langch
	case "en":
		var langen langEn
		err := json.Unmarshal(context, &langen)
		if err != nil {
			log.Fatalln(err)
		}
		return langen
	}
	return nil
}
func getContent(filePath string) []byte {
	context, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	return context
}
