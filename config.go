package memontago

type Config struct {
	Language string // 语言
	Location string // 时区 只在传输字符串表示的事件的时候起作用
}

//默认加载英文
var config = Config{
	Language: "en",
}

// 初始化时默认加载英文格式
func init() {
	trans("en")
	EnTrans = initEnTrans()
}
func SetConfig(conf Config) {
	//设置语言 如果传入的配置为空，设置默认语言
	if conf.Language == "" {
		conf.Language = config.Language
	}
	config = conf
	//	当配置config时，首先看看该语言是否已经加载到langs中了没
	if _, ok := langs[config.Language]; !ok {
		trans(config.Language)
		ChTrans = initChTrans()
	}
}
