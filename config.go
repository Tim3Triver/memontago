package memontago

type Config struct {
	Language string
	Location string
}

var config = Config{
	Language: "en",
}

var GlobalOptions = []string{}

func SetConfig(conf Config) {
	//设置语言 如果传入的配置为空，设置默认语言
	if conf.Language == "" {
		conf.Language = config.Language
	}
	config = conf
	//	当配置config时，首先看看该语言是否已经加载到langs中了没
	if _, ok := langs[config.Language]; !ok {
		trans(config.Language)
	}
}
