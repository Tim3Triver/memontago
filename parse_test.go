package memontago

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	SetConfig(Config{Language: "ch", Location: "Asia/Shanghai"})
	a, err := Parse(time.Now().Add(-20*time.Second), "online")

	fmt.Println(a, err)
}

//func TestCalender(t *testing.T) {
//	SetConfig(Config{Language: "ch"})
//	fmt.Println(Calender(time.Now().AddDate(0, 0, 2)))
//	fmt.Println(Calender(time.Now().AddDate(0, 0, -4)))
//	fmt.Println(Calender(time.Now().AddDate(0, 0, -1)))
//	fmt.Println(Calender(time.Now().AddDate(0, 2, 0)))
//	fmt.Println(Calender(time.Now().AddDate(0, 0, 0)))
//}
func TestCalender(t *testing.T) {
	var a map[string]int
	fmt.Println(len(a))

}
