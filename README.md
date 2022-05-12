# memontago

#### 介绍

**该软件是go的一个time库**，可以实现对于时间的解析，能够处理三种形式的时间：
1. time.Time
2. Unix时间戳（timestamp）
3. 字符串格式的时间 2006-01-02 15:04:05 [+ 时区]

能够支持两种语言：

1. Chinese
2. English
- 支持解析时间，判断给定的一个时间是当前时间的关系

- 支持在给定时间上增加，减少duration

- 支持time.Time转换为日历格式的时间

- 支持time.Time的日期格式化
#### 软件架构
  1. memontago.Config() 
  2. memontago.Parse()
  3. memontago.Add()
  4. memontago.Sub()
  5. memontago.Format()
  6. memontago.Calender()
#### 安装教程
> go get github.com/Tim3Triver/memontago

#### 使用说明

```go
func Config(Config{Language string, Location string})
```
> 配置语言和时区
```go
func parseTime(datetime interface{}, options ...string) string 
```
> 解析时间：datetime与当前时间的关系
> 
> 例如，datetime是当前时间的前一分钟，返回 1 分钟以前
> >支持两个选项 online、justNow
> >
> > 如果加上选项，给定的时间是当前时间之前30秒内，会返回Online，给定的时间是当前时间之前60秒内并且大于30秒，会返回Just Now



```go
func add(number int, DurUnit string) time.Time
```

> datetime增加number个DurUnit之后的时间
>
> DurUnit string 支持 单位：second/minute/day/week/month/year
>
> 1 month = 31day
>
> > 支持时间的进位
> >
> > 例如：
> >
> > 给定的时间为：19:30:30
> >
> > date：second =40 转换后的时间为 19:31:10



```go 
func Sub (datetime,number, DurUnit string) time.Time
```

> datetime减少number个DurUnit之后的时间
```go
func Format(datetime interface{}, format string) string 
```
> 格式化给定时间为指定格式
> 
```go
func Calender(datetime interface{}) string
```
> 给定一个日期，返回给定的日历日期