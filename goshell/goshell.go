package goshell

import (
	"fmt"
	"goshell/debug"
)

const (
	DEBUGLEVEL = debug.NORMAL
)

// func Goshell(cmd string, params map[string]string) {
// 	Goshell(cmd, params, DEBUGLEVEL)
// }

func Goshell(cmd string, params map[string]string, debuglevel int) {
	var param_string string
	for param_flag, param_value := range params {
		param_string = param_string + " " + param_flag
		if len(param_value) != 0 {
			param_string = param_string + " " + param_value
		}
	}

	fmt.Println(cmd, param_string)
	debug.Exec(cmd, param_string, debuglevel)
}
