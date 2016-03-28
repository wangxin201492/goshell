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

func execute(command string) {
	//cmd := exec.Command(command)
	//err := cmd.Run()
	out, err := exec.Command(command).Output()
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("The date is %s\n", out)
}

func print(command string) {
	log.InfoCyan("(sh)", command)
}

func SetExecLevel(level ExecLevel) {
	execLevel = level
}
