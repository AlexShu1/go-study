package main

import (
	"errors"
	"fmt"
	interfacetest "go-study/com/alex"
)

/*
*
1. 值传递：普通变量, 数组, 对象.
2. 引用传递: 指针, slice,map, channel

数组和slice的区别
 1. 创建:
    数组：p := [3] int {1, 3, 5}
    slice：p := []int {1,3,5}
    区别：slice不指定数组大小

2.形参：

	数组： func t(p [3] int)
	slice: func t(p [] int)
	区别：slice不指定数组大小

在Go语言中通过首字母大小写来控制变量、函数、方法、类型的公私有
如果首字母小写, 那么代表私有, 仅在当前包中可以使用
如果首字母大写, 那么代表共有, 其它包中也可以使用
*/
type Person struct {
	name string
	age  int
}

func main() {
	// 值传递 引用传递 数组 slice map
	//funcTest()

	// struct 对象
	//structTest()

	// 指针
	//pointerTest()

	// interface 和  struct
	//interfaceTest()

	// 测试error是否会中断程序: error并不会中断程序的正常运行
	//err := errorTest()
	//if err != nil {
	//	fmt.Println(err)
	//}

	// panic中断测试
	//interfacetest.Div(10, 0)

	// string 相关的方法测试
	interfacetest.TestStringSFunc()
}

func errorTest() error {
	fmt.Println("error...")
	err := errors.New("this is an error")
	fmt.Println("error after")
	return err
}

func interfaceTest() {
	cm := interfacetest.Computer{"戴尔笔记本电脑", "F1234"}
	interfacetest.Working(&cm)

	apple := interfacetest.Phone{"Apple手机 "}
	interfacetest.Working(&apple)
}

func pointerTest() {
	num := 2
	var p = &num
	fmt.Println("num:", num)
	fmt.Println("num:", *p)

	*p = 23
	fmt.Println("num:", num)
	fmt.Println("num:", *p)

	var per = &Person{"John", 20}
	fmt.Println("per:", per)
	per.age = 21
	per.name = "Alex"
	fmt.Println("per:", (*per).name)
	fmt.Println("per:", (*per).age)
}

func structTest() {
	var person Person = Person{"John", 20}
	fmt.Println(person)
	person.personFunc1()
	fmt.Println(person)
	person.personFunc2()
	fmt.Println(person)

	var person2 Person = Person{name: "Alex"}
	fmt.Println(person2)
}

/*
*
方法和对象关联在一起;
这种绑定关系可以看作是一种形参传递, 也分为值传递和引用传递
*/
func (p Person) personFunc1() {
	fmt.Println("person:", p)
	p.age = 233
}

/*
*
方法和对象关联在一起
*/
func (p *Person) personFunc2() {
	fmt.Println("person2:", p)
	p.age = 233
}

func funcTest() {
	p := Person{"lnj", 33}

	// 对象传递
	change(p)
	fmt.Println(p.name) // lnj

	// 指针传递
	p1 := 2
	change1(&p1)
	fmt.Println(p1) // 3

	// slice
	p2 := []int{1, 2}
	change2(p2)
	fmt.Println(p2) /// 100 2

	// map
	p3 := map[string]int{"name": 2, "age": 3}
	change3(p3)
	fmt.Println(p3) /// name: 21

	// 匿名函数
	func(s string) {
		fmt.Println(s)
	}("hello lnj")

	// slice copy
	// 容量为3
	var sce1 = []int{1, 3, 5}
	// 容量为5
	var sce2 = make([]int, 2)
	fmt.Println("拷贝前:", sce2) // [0 0 0 0 0]
	// sce2容量足够, 会将sce1所有内容拷贝到sce2
	copy(sce2, sce1)
	fmt.Println("拷贝后:", sce2) // [1 3 5 0 0]
}

func change3(p3 map[string]int) {
	p3["name"] = 21
}

func change2(p2 []int) {
	p2[0] = 100
}

// 对象传递
func change(p Person) {
	p.name = "zs"
}

// 指针传递
func change1(p *int) {
	*p = *p + 1
}
