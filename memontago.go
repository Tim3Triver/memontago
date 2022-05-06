package memontago

//Parse 解析并判断给定的时间与当前的时间的关系
func Parse(datetime interface{}, options ...string) string {
	return ParseTime(datetime)
}

//func Equals(datetime1, datetime2 interface{}) (bool, error) {
//	return false, nil
//}
//func Before(datetime1, datetime2 interface{}) (bool, error) {
//	return false, nil
//}
//func After(datetime1, datetime2 interface{}) (bool, error) {
//	return false, nil
//}
