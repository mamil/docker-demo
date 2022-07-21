package main

import (
	"docker-demo/container"
	"os"

	log "github.com/sirupsen/logrus"
)

func Run(tty bool, command string) {
	log.Infof("Run, tty:%v, command:%v", tty, command)
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
