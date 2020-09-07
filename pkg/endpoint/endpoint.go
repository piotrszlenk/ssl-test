package endpoint

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/piotrszlenk/ssl-test/pkg/logz"
)

type Endpoints struct {
	Items []Endpoint
	path  string
}

type Endpoint struct {
	Fqdn string
	Port uint64
}

func NewEndpoints(path string) *Endpoints {
	e := new(Endpoints)
	e.path = path
	return e
}

func (e *Endpoints) LoadEndpoints() (Endpoints, error) {
	csvfile, err := os.Open(e.path)
	if err != nil {
		logz.Logger().Error.Fatalln("Unable to read the file.")
	}
	r := csv.NewReader(csvfile)
	for {
		entry, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logz.Logger().Error.Printf("Unable to parse entry: %s", entry)
			continue
		}
		logz.Logger().Debug.Printf("Read record: %s", entry)
		if len(entry) == 2 {
			port, err := strconv.ParseUint(entry[0], 10, 16)
			if err != nil {
				logz.Logger().Error.Printf("Unable to parse port number from line: %s", entry)
				continue
			}
			e.addEndpoint(Endpoint{entry[1], port})
		} else {
			logz.Logger().Error.Printf("Unable to parse entry: %s", entry)
		}
	}

	logz.Logger().Debug.Printf("Loaded %d records from %s.", len(e.Items), e.path)

	return *e, nil
}

func (e *Endpoints) addEndpoint(ept Endpoint) {
	e.Items = append(e.Items, ept)
}
