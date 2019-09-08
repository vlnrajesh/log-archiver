package utils

import (
	"flag"
)

func CmdLineParser() {
	// Read Command Line Operations and validate
	var confFile string
	const confFileUsageString = "Configuration file to read"
	flag.StringVar(&confFile, "--conf", "/etc/log-archiver.json", confFileUsageString)
	flag.StringVar(&confFile, "--c", "/etc/log-archiver.json", confFileUsageString)
	flag.Parse()
	flag.PrintDefaults()
}
