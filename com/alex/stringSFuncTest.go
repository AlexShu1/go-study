package interface_test

import (
	"fmt"
	"regexp"
	"strings"
)

func TestStringSFunc() {
	//length()
	//index()
	//equals()
}

func equals() {
	fmt.Println(strings.EqualFold("hello2", "hello"))
}

func length() {
	str1 := "lnj"
	fmt.Println(len(str1)) // 3
	str2 := "公号代码情缘"
	fmt.Println(len(str2)) // 21  一个汉字3个字节

	// Go语言中rune类型就是专门用于保存汉字的
	arr2 := []rune(str2)
	arr3 := []rune(str1)
	fmt.Println(len(arr2)) // 6
	fmt.Println(len(arr3)) // 3
}

func index() {
	str2 := "公号代码情缘"
	fmt.Println(strings.IndexRune(str2, '号'))
}

func findPhoneNumber(str string) bool {
	// 创建一个正则表达式匹配规则对象
	reg := regexp.MustCompile("^1[1-9]{10}")
	// 利用正则表达式匹配规则对象匹配指定字符串
	res := reg.FindAllString(str, -1)
	if res == nil {
		return false
	}
	return true
}
