package memontago

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	SetConfig(Config{Language: "en", Location: "Asia/Shanghai"})
	a := Parse(time.Now().Add(-50*time.Second).Format("2006-01-02 15:04:05"), "justNow")
	fmt.Println(a)
}
func TestA(t *testing.T) {
	fmt.Println(time.Now().Format(""))
}

//func A(){
//
//}
func TestFormat(t *testing.T) {
	fmt.Println(time.Now().Weekday())
}
