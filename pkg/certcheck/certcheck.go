package certcheck

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/piotrszlenk/ssl-test/pkg/logz"

	"github.com/piotrszlenk/ssl-test/pkg/endpoint"
)

type TestTargets struct {
	Items  []*TestTarget
	caPath *string
}

type TestStatus int32

const (
	OK            TestStatus = 1
	FAILED        TestStatus = 0
	NOT_COMPLETED TestStatus = 2
)

type TestTarget struct {
	ept  endpoint.Endpoint
	res  TestStatus
	crts []*x509.Certificate
}

func NewTestTargets(epts *endpoint.Endpoints, caPath *string) *TestTargets {
	tt := new(TestTargets)
	tt.caPath = caPath

	for _, e := range epts.Items {
		tt.Items = append(tt.Items, &TestTarget{e, NOT_COMPLETED, nil})
	}
	return tt
}

func Test(c <-chan *TestTarget, caPath *string, wg *sync.WaitGroup) {
	defer wg.Done()
	for tt := range c {
		logz.Logger().Info.Printf("Testing endpoint %s on port %d\n", tt.ept.Fqdn, tt.ept.Port)
		tt.testCertificateChain(caPath)
	}
}

func (t *TestTargets) Test() *TestTargets {
	var numWorkers = 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	c := make(chan *TestTarget, len(t.Items))
	for _, i := range t.Items {
		c <- i
		logz.Logger().Debug.Printf("Sending TestTarget (%s on port %d) to the channel.\n", i.ept.Fqdn, i.ept.Port)
	}

	for i := 0; i < numWorkers; i++ {
		go Test(c, t.caPath, &wg)
	}

	close(c)
	wg.Wait()
	return t
}

func (t *TestTarget) testCertificateChain(caPath *string) (*TestTarget, error) {
	cfg := tls.Config{InsecureSkipVerify: true}
	connStr := fmt.Sprintf("%s:%d", t.ept.Fqdn, t.ept.Port)

	duration, _ := time.ParseDuration("15s")
	dialer := net.Dialer{Timeout: duration}
	conn, err := tls.DialWithDialer(&dialer, "tcp", connStr, &cfg)
	if err != nil {
		logz.Logger().Error.Printf("Failed to connect: %s\n", err.Error())
		t.res = FAILED
		return t, nil
	}
	defer conn.Close()

	logz.Logger().Debug.Printf("Certificate chain sent by the server: %s:%d\n", t.ept.Fqdn, t.ept.Port)
	t.res = OK
	t.crts = conn.ConnectionState().PeerCertificates
	return t, nil
}

func (t *TestTargets) PrintResults() {
	for _, testtarget := range t.Items {
		if testtarget.res == FAILED {
			logz.Logger().Info.Printf("Test failed for endpoint: %s:%d\n", testtarget.ept.Fqdn, testtarget.ept.Port)
		}
		if testtarget.res == OK {
			logz.Logger().Info.Printf("Certificate chain sent by the endpoint: %s:%d\n", testtarget.ept.Fqdn, testtarget.ept.Port)
			for _, cert := range testtarget.crts {
				logz.Logger().Info.Printf("CN: %s OU: %s ValidNotBefore: %s ValidNotAfter: %s\n", cert.Subject.CommonName, cert.Subject.Organization, cert.NotBefore, cert.NotAfter)
			}
		}
	}
}
