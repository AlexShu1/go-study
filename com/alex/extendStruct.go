package interface_test

import "fmt"

/**
go 是使用组合的方式, 来达到不同的对象之间的继承关系
*/

type person struct {
	name string
}

type police struct {
	person
	title string
}

type teacher struct {
	person
	title string
}

/*
*
父类的方法
*/
func (p person) say() {
	fmt.Println("person.say(): ", p.name)
}

/*
*
子类的重写方法
*/
func (p teacher) say() {
	fmt.Println("teacher.say(): ", p.name, p.title)
}
