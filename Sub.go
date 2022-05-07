package memontago

import "time"

func Sub(dateTime interface{}, number int, DurUnit string) time.Time {
	return Add(dateTime, -number, DurUnit)
}
