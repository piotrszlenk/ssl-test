package main

import (
	"flag"

	"github.com/piotrszlenk/ssl-tst/pkg/logz"
)

func main() {
	// Parse CLI arguments
	fqdnFile = flag.String("fqdnlist", "fqdns.txt", "Flat file with list of FQDNs and ports")
	debugFlag := flag.Bool("debug", false, "Enable debug logging.")
	flag.Parse()
	lc := logz.LoggerConfig{*debugFlag}
	logger := logz.GetInstance(lc)
	logger.Debug.Println("Command line arguments: ")
	logger.Debug.Println(" -fqdnlist set to:", *fqdnFile)
	logger.Debug.Println(" -debug set to:", *debugFlag)

	// Create list of fqdn objects
	logger.Info.Println("Loading FQDNs from:", *fqdnFile)
	fqdns, err := fqdn.LoadFQDNs(*fqdnFile, *logger)
	if err != nil {
		logger.Error.Fatalln(err)
	}

}
