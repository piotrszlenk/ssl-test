package main

import (
	"flag"

	"github.com/piotrszlenk/ssl-tst/pkg/logz"
)

func main() {
	// Parse CLI arguments
	urlListFile = flag.String("urllist", "urls.txt", "Flat file with list of URLs")
	debugFlag := flag.Bool("debug", false, "Enable debug logging.")
	flag.Parse()
	lc := logz.LoggerConfig{*debugFlag}
	logger := logz.GetInstance(lc)
	logger.Debug.Println("Command line arguments: ")
	logger.Debug.Println(" -urlist set to:", *urlListFile)
	logger.Debug.Println(" -debug set to:", *debugFlag)

	// Create empty device list and load content from inventory file
	logger.Info.Println("Loading device list from:", *databaseFile)
	devices := device_inventory.NewDevices(*databaseFile, *logger)
	_, err := devices.LoadDevices()
	if err != nil {
		logger.Error.Fatalln(err)
	}

}
