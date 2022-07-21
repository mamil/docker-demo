package main

import (
	"docker-demo/container"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			mydocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		log.Infof("runCommand start, args:%+v", context.Args())
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti") || context.Bool("it")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init start")
		cmd := context.Args().Get(0)
		log.Infof("init command:%s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
