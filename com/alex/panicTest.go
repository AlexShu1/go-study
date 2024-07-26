package interface_test

import "fmt"

/**
中断程序
*/

func Div(a, b int) (res int) {
	// defer recover 必须要定义在panic之前
	defer func() {
		if err := recover(); err != nil {
			res = -1
			fmt.Println(err) // 除数不能为0
		}
	}()

	if b == 0 {
		//一旦传入的除数为0, 程序就会终止
		panic("除数不能为0")
	} else {
		res = a / b
	}

	return
}
