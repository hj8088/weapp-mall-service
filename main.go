package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"weapp-mall-service/cmd"
)

var (
	VERSION string
	BUILD   string
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var (
		appName = filepath.Base(os.Args[0])
	)

	app := cli.NewApp()
	app.Name = appName
	app.Description = "backend service of wechat app of mall"
	app.Version = fmt.Sprintf("%s, build %s", VERSION, BUILD)
	app.Action = cmd.ServerMain

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:     "debug",
			Required: false,
			Usage:    "show debug info",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("%s execute failed, err:%v.\n", appName, err)
	}
}
