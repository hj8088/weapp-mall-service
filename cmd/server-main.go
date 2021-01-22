package cmd

import (
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func ServerMain(c *cli.Context) error {

	exePath := filepath.Dir(os.Args[0])
	ResourcePath = filepath.Join(exePath, "resource")

	signal.Notify(globalOSSignalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go handleSignals()

	webHandler, _ := configureServerHandler()

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: webHandler,
	}

	httpServer.ListenAndServe()

	return nil
}
