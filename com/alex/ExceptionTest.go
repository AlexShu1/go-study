package interface_test

import (
	"errors"
	"fmt"
)

/**
异常的相关知识
*/

// 创建error对象抛出异常信息
func exception1() {
	err := errors.New("this is an error")
	fmt.Println(err)
}
