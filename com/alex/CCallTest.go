package interface_test

// 学习如何调用C语言
/*
#cgo CFLAGS: -I.
#include "Hello.c"
*/
import "C"

func CLanguageCall() {
	C.say()
}
