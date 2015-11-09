package main

import (
	"fmt"
	"github.com/wangxin201492/goshell/goshell"
)

func main() {
	fmt.Println("main function")

	map1 := make(map[string]string)
	map1["-r"] = "bj"
	map1["-o"] = "proxy.conf.new"
	map1["-s"] = ""

	// for key, value := range map1 {
	// 	fmt.Printf("key:%s\tvalue:%s\n", key, value)
	// }

	goshell.Goshell("cmd", map1, 1)
}
