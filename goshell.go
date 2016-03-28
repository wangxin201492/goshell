package goshell

import (
	"fmt"
	"os/exec"

	log "gitlab.baidu.com/wangxin44/golog"
)

var (
	execLevel ExecLevel = NORMAL
)

func Shell(cmd string) {

	//fmt.Println(cmd)
	switch execLevel {
	case NORMAL:
		execute(cmd)
	case DEBUG:
		print(cmd)
	}
}

func execute(cmd string) {
	print(cmd)
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("The date is %s\n", out)
}

func print(cmd string) {
	log.InfoCyan("(sh) " + cmd)
}

func SetExecLevel(level ExecLevel) {
	execLevel = level
}
