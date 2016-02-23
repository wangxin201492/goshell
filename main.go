package goshell

import (
	"fmt"
)

func main() {
	fmt.Println("main function")
	test()

	Shell("cmd")
}

func test() {
	fmt.Println("test")
}
