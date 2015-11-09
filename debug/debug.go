package debug

import (
	"fmt"
	"os/exec"
)

func Exec(cmd string, param string, debugLevel int) {
	fmt.Println("debuglevel:", debugLevel)

	switch debugLevel {
	case DEBUG:
		fmt.Println(cmd, param)
	case NORMAL:
		cmd := exec.Command(cmd, param)
		fmt.Println(cmd.Args, cmd.Dir, cmd.Env, cmd.Args, cmd.Path)
	}

}
