package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Println(os.Args)
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("have not define")
	}
}
func run() {
	cmd := exec.Command(os.Args[2])
	cmd.SysProcAttr = &syscall.SysProcAttr{ // 需要sudo权限运行
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
		// Credential: &syscall.Credential{
		// 	Uid: uint32(1),
		// 	Gid: uint32(1),
		// },
		GidMappingsEnableSetgroups: true,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(-1)
}
