package goshell

import (
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
	_, err := exec.Command(command).Output()
	if err != nil {
		log.Error(err)
	}
}

func print(command string) {
	log.InfoCyan("(sh)", command)
}

func SetExecLevel(level ExecLevel) {
	execLevel = level
}
