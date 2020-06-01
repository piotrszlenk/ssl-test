package main

import (
	"flag"

	"github.com/piotrszlenk/ssl-test/pkg/endpoint"
	"github.com/piotrszlenk/ssl-test/pkg/logz"
)

func main() {
	// Parse CLI arguments
	endpointFile := flag.String("endpointlist", "endpoints.txt", "Flat file with list of endpoints and ports")
	debugFlag := flag.Bool("debug", false, "Enable debug logging.")
	flag.Parse()
	lc := logz.LoggerConfig{*debugFlag}
	logger := logz.GetInstance(lc)
	logger.Debug.Println("Command line arguments: ")
	logger.Debug.Println(" -endpointlist set to:", *endpointFile)
	logger.Debug.Println(" -debug set to:", *debugFlag)

	// Create list of endpoint objects
	logger.Info.Println("Loading endpoints from:", *endpointFile)
	endpoints := endpoint.NewEndpoints(*endpointFile, *logger)
	_, err := endpoints.LoadEndpoints()
	if err != nil {
		logger.Error.Fatalln(err)
	}

}
