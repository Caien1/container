package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//dovker run imagen <cmd> <paras>
// go run main.go run

func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("bad command")

	}
}

func run() {
	fmt.Print("Running ", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	cmd.Run()

}
