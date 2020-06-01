package endpoint

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/piotrszlenk/ssl-test/pkg/logz"
)

type Endpoints struct {
	Items  []Endpoint
	logger logz.LogHandler
	path   string
}

type Endpoint struct {
	Fqdn string
	Port uint64
}

func NewEndpoints(path string, logger logz.LogHandler) *Endpoints {
	e := new(Endpoints)
	e.logger = logger
	e.path = path
	return e
}

func (e *Endpoints) LoadEndpoints() (Endpoints, error) {
	csvfile, err := os.Open(e.path)
	if err != nil {
		e.logger.Error.Fatalln("Unable to read the file.")
	}
	r := csv.NewReader(csvfile)
	for {
		entry, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			e.logger.Error.Fatalln("Unable to read record.")
		}
		e.logger.Debug.Printf("Read record: %s", entry)
		port, _ := strconv.ParseUint(entry[0], 10, 16)
		e.addEndpoint(Endpoint{entry[1], port})
	}

	e.logger.Debug.Printf("Loaded %d records from %s.", len(e.Items), e.path)

	return *e, nil
}

func (e *Endpoints) addEndpoint(ept Endpoint) {
	e.Items = append(e.Items, ept)
}
