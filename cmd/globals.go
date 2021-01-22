package cmd

import "os"

var (
	ResourcePath     = ""
	globalOSSignalCh = make(chan os.Signal, 1)
)
