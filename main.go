package main

import (
	"flag"

	"github.com/piotrszlenk/ssl-test/pkg/certcheck"
	"github.com/piotrszlenk/ssl-test/pkg/endpoint"
	"github.com/piotrszlenk/ssl-test/pkg/logz"
)

func main() {
	// Parse CLI arguments
	endpointFile := flag.String("endpoints-file", "endpoints.txt", "Flat file with list of endpoints and ports")
	debugFlag := flag.Bool("debug", false, "Enable debug logging.")
	caPath := flag.String("capath", "/private/etc/ssl", "Path to OpenSSL CA dir")

	flag.Parse()
	lc := logz.LoggerConfig{*debugFlag}
	logger := logz.GetInstance(lc)
	logger.Debug.Println("Command line arguments: ")
	logger.Debug.Println(" -endpoints-file set to:", *endpointFile)
	logger.Debug.Println(" -debug set to:", *debugFlag)
	logger.Debug.Println(" -capath set to:", *caPath)

	// Create list of endpoint objects
	logger.Info.Println("Loading endpoints from:", *endpointFile)
	endpoints := endpoint.NewEndpoints(*endpointFile, *logger)
	_, err := endpoints.LoadEndpoints()
	if err != nil {
		logger.Error.Fatalln(err)
	}

	//Create test targets
	logger.Info.Println("Creating test targets from loaded endpoints.\n")
	testtargets := certcheck.NewTestTargets(endpoints, caPath, *logger)
	testtargets.Test(*logger)
	//logger.Debug.Print(testtargets)
}
