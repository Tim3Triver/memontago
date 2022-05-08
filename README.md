# memontago

#### 介绍

**该软件是go的一个time库**，可以实现对于时间的解析，判断给定的一个时间是当前时间的关系。
能够处理三种形式的时间：
1. time.Time
2. Unix时间戳（timestamp）
3. 字符串格式的时间 2006-01-02 15:04:05

能够支持两种语言：

1. Chinese
2. English

#### 软件架构
1. memontago.Parse() 

2. memontago.Config()

#### 安装教程
后来再说

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

