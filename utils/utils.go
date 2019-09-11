package utils

import (
	"flag"
)

//CmdLineParser is parse command line options and validate inputs
func CmdLineParser() {
	var confFile string
	const confFileUsageString = "Configuration file to read"
	flag.StringVar(&confFile, "--conf", "/etc/log-archiver.json", confFileUsageString)
	flag.StringVar(&confFile, "--c", "/etc/log-archiver.json", confFileUsageString)
	flag.Parse()
	flag.PrintDefaults()
}
