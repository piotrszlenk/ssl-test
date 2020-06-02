package certcheck

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"time"

	"github.com/piotrszlenk/ssl-test/pkg/logz"

	"github.com/piotrszlenk/ssl-test/pkg/endpoint"
)

type TestTargets struct {
	Items  []TestTarget
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
		tt.Items = append(tt.Items, TestTarget{e, NOT_COMPLETED, nil})
	}
	return tt
}

func (t *TestTargets) Test() *TestTargets {
	for _, i := range t.Items {
		logz.Logger().Info.Printf("Testing %s on port %d\n", i.ept.Fqdn, i.ept.Port)
		i.testCertificateChain(t.caPath)
	}
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

	logz.Logger().Info.Printf("Certificate chain sent by the server: %s:%d\n", t.ept.Fqdn, t.ept.Port)
	for _, cert := range conn.ConnectionState().PeerCertificates {
		logz.Logger().Info.Printf("CN: %s OU: %s ValidNotBefore: %s ValidNotAfter: %s\n", cert.Subject.CommonName, cert.Subject.Organization, cert.NotBefore, cert.NotAfter)
		t.res = OK
		t.crts = conn.ConnectionState().PeerCertificates
	}
	return t, nil
}
