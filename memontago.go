package memontago

//Parse
func Parse(datetime interface{}, options ...string) (string, bool) {
	ParseTime(datetime)
	return "", false
}
func Equals(datetime1, datetime2 interface{}) (bool, error) {
	return false, nil
}
func Before(datetime1, datetime2 interface{}) (bool, error) {
	return false, nil
}
func After(datetime1, datetime2 interface{}) (bool, error) {
	return false, nil
}
