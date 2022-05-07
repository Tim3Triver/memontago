package memontago

//Parse 解析并判断给定的时间与当前的时间的关系
func Parse(datetime interface{}, options ...string) string {

	return ParseTime(datetime, options...)
}
