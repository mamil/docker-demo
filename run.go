package main

import (
	"docker-demo/cgroups"
	"docker-demo/cgroups/subsystems"
	"docker-demo/container"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Run(tty bool, comArray []string, res *subsystems.ResourceConfig, volume string) {
	parent, writePipe := container.NewParentProcess(tty, volume)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil { // 启动容器
		log.Error(err)
	}

	if res != nil {
		// use docker-demo as cgroup name
		// 设置资源限制
		cgroupManager := cgroups.NewCgroupManager("docker-demo1") // 如果名字是docker-demo，会把执行文件删掉
		defer cgroupManager.Destroy()
		cgroupManager.Set(res)
		cgroupManager.Apply(parent.Process.Pid)
	}

	// 初始化容器
	sendInitCommand(comArray, writePipe)
	parent.Wait()

	log.Infof("Run after wait")

	// vloume
	mntURL := "/root/mnt"
	rootURL := "/root"
	container.DeleteWorkSpace(rootURL, mntURL, volume)

	log.Infof("Run end")
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
