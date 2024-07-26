package alex

import "fmt"

type user interface {
	start()
	stop()
}

type Computer struct {
	name  string
	model string
}

func (c *Computer) start() {
	fmt.Println("start")
}
func (c *Computer) stop() {

}
