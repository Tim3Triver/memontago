package memontago

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	SetConfig(Config{Language: "ch", Location: "Asia/Shanghai"})
	a := Parse(time.Now().Add(-20*time.Second), "online")
	fmt.Println(a)
}
func TestA(t *testing.T) {

}

//func A(){
//
//}
func TestFormat(t *testing.T) {
	fmt.Println(time.Now().Weekday())
}
