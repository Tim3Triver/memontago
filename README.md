# memontago

#### 介绍

该工具基于Go的时间包开发的时间处理库，可以将传入不同格式的时间转换为字符串：”n时间前“，举个例子：比如我们获取一下此刻的时间，再将此刻的时间减去一分钟，然后将这个时间作为momentago中的Parse函数的参数，这个函数能够返回一个字符串：而这个字符串就是“1分钟前”，能够处理三种形式的时间：
1. time.Time
2. Unix时间戳（timestamp）
3. 字符串格式的时间 [+ 时区]

`注意：` Parse函数在解析字符串类新时间时，需要传入Go官方支持的**2006-01-02 15:04:05**格式化解析字符串，
否则传入为空时`""`，默认按照2006-01-02 15:04:05格式去解析

已经支持的语言：
1. Chinese
2. English
- 解析时间，判断给定的一个时间是当前时间的关系

#### 软件架构
  1. memontago.SetConfig() 
  2. memontago.Parse()
#### 安装教程
> go get github.com/Tim3Triver/memontago

#### 使用说明
 第一步：将该依赖中的language目录复制到项目目录下
 第二步：通过momentago.SetConf()，配置语言，时区，是否需要特殊标识
>func SetConfig(conf Config) error {
```go
type Config struct {
	Language string // 语言
	Location string // 时区 只在解析字符串格式表示的时间时起作用
	Special  bool   //是否支持特殊标志
}
```
> 如果不配置，则默认英文，本地时区，不支持特殊标志
> 
> 如果使用标识，则小于等于5秒则以Online显示，大于5秒 小于等于30秒以JustNow显示，而不会显示解析出来的时间差
>

第三步：使用momentago.Parse()，解析时间到的结果
> func Parse(datetime interface{}, format string) (string, error) {
datetime：输入以上三种格式的时间，进行解析。
## 扩展新的显示语言
 如果想加入新的语言可以在language目录下添加yaml格式配置文件，使用的时候在配置文件中指定相应的文件名（不带后缀：`.yaml`）作为Config中的Language的值

例如：
1.添加配置项 ![img](https://user-images.githubusercontent.com/91925733/193413042-c7349d62-ff22-41ab-ad26-b025bbe5a361.png)
2.SetConfig(Config{ Language : "en" }) 
