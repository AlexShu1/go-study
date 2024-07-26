package interface_test

import (
	"fmt"
	"math/rand"
	"time"
)

/**
时间相关的方法
*/

func TimeTest() {
	timeFormat()
}

func timeFormat() {
	var t time.Time = time.Now()
	var dateStr = fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println("当前的时间是:", dateStr)

	str1 := t.Format("2006/01/02 15:04:05")
	fmt.Println(str1)

	// time.Sleep(time.Second * 1)

	fmt.Println(rand.Intn(10))
}
