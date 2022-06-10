package memontago

import "time"

// Sub 支持在给定时间上减少 number * DurUnit
func Sub(dateTime interface{}, number int, DurUnit string) (time.Time, error) {
	return Add(dateTime, -number, DurUnit)
}
