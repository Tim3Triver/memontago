package memontago

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	//time := Add("2022-05-07 15:04:05", 1, "second")
	//time := Add("2022-05-07 15:04:05", 1, "minute")
	//time := Add("2022-05-07 15:04:05", 1, "week")
	//time := Add("2022-05-07 15:04:05", 1, "day")
	//time := Add("2022-05-07 15:04:05", 1, "year")

	durUnit := []string{"second", "minute", "hour", "week", "day", "month", "year"}
	for i := 0; i < 7; i++ {
		time1, err := Add("2022-01-31 15:00:00", 1, durUnit[i])
		//time2 := ("2022-05-07 15:04:00", 1, durUnit[i])
		fmt.Println(durUnit[i], time1, err)
		//fmt.Println(durUnit[i], time2)
	}
}
