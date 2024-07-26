package main

import (
	"fmt"
	"go-study/com/alex"
)

type Student struct {
	name string
	age  int
}

/*
*
1. _:代表的是忽略的意思, 若不想接收某个返回值, 可以通过这个标识.
2. 变量定义不使用会报错.
*/
func main() {
	const name = "CONST_NAME"

	alex.Femo1()
	num := 41
	fmt.Printf("Hello World. %d\n", num)
	fmt.Println("This is name:", name, "OK!")
	sPrint := fmt.Sprintln("--> This is name:", name, "OK!")
	fmt.Println(sPrint)

	var num1, num2 int
	fmt.Scanf("%d%d", &num1, &num2)
	fmt.Println(num1, num2)

	if age := 14; age >= 18 {
		fmt.Println("This is 成年人啦 age: ", age)
	} else {
		fmt.Println("还是未成年哦 age: ", age)
	}

	switch num1 {
	case 10:
		fmt.Println("<10")
	case 20:
		fmt.Println("<20")
	default:
		fmt.Println("Default!")

	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	arr := [3]int{1, 3, 5}
	for k, _ := range arr {
		fmt.Println(k, arr[k])
	}

	for k, v := range arr {
		fmt.Println(k, v)
	}
}

func t1() {
	fmt.Printf("Hello World. %d\n", 1)
}
