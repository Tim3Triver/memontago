package memontago

// Config 语言和时区
type Config struct {
	Language string // 语言
	Location string // 时区 只在传输字符串表示的事件的时候起作用
	Special  bool   //是否支持特殊标志
}

//默认加载英文
var config = Config{
	Language: "en",
}

// SetConfig 配置语言和时区
// 时区仅在其他函数给定的时间参数为字符串格式的时间起作用
func SetConfig(conf Config) error {
	//语言为空时，设置英语为默认语言
	if conf.Language == "" {
		conf.Language = "en"
	}
	if conf.Location == "" {
		conf.Location = "Asia/Shanghai"
	}
	config = conf
	// 解析对应语言 到结构体中
	return UnMarshalToStruct(config.Language)
}
