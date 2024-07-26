package interface_test

import "fmt"

type Usb interface {
	Start()
	stop()
}

type Computer struct {
	Name  string
	Model string
}

func (c *Computer) Start() {
	fmt.Println(c.Name + "Start")
}
func (c *Computer) stop() {
	fmt.Println(c.Name + "stop")
}

type Phone struct {
	Name string
}

func (p *Phone) Start() {
	fmt.Println(p.Name + "Start")
}
func (p *Phone) stop() {
	fmt.Println(p.Name + "stop")
}

// 3.使用接口定义的方法
func Working(u Usb) {
	u.Start()
	u.stop()
}
