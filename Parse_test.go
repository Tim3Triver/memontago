package memontago

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	SetConfig(Config{Language: "en", Location: "Asia/Shanghai"})
	a := Parse(time.Now().Add(-1 * time.Minute).Format("2006-01-02 15:04:05"))
	fmt.Println(a)
}
