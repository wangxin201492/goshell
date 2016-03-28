package goshell

import "testing"
import "os"
import "fmt"

func TestShell1(t *testing.T) {
	Shell("ls")
}

func TestEnv(t *testing.T) {
	fmt.Println(os.Getenv("PATH"))
}
