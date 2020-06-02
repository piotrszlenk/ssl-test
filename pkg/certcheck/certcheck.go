package certcheck

import (
	"crypto/tls"
	"fmt"

	"github.com/piotrszlenk/ssl-test/pkg/endpoint"
	"github.com/piotrszlenk/ssl-test/pkg/logz"
)

type TestTargets struct {
	Items  []TestTarget
	logger logz.LogHandler
	caPath *string
}

type TestTarget struct {
	ept endpoint.Endpoint
	res string
}

func NewTestTargets(epts *endpoint.Endpoints, caPath *string, logger logz.LogHandler) *TestTargets {
	tt := new(TestTargets)
	tt.caPath = caPath
	tt.logger = logger

	for _, e := range epts.Items {
		tt.Items = append(tt.Items, TestTarget{e, ""})
	}
	return tt
}

func (t *TestTargets) Test(logger logz.LogHandler) {
	for _, i := range t.Items {
		logger.Info.Printf("Testing %s on port %d\n", i.ept.Fqdn, i.ept.Port)
		i.testCertificateChain(t.caPath, t.logger)
	}
}

func (t *TestTarget) testCertificateChain(caPath *string, logger logz.LogHandler) (*TestTarget, error) {
	cfg := tls.Config{InsecureSkipVerify: true}
	connStr := fmt.Sprintf("%s:%d", t.ept.Fqdn, t.ept.Port)
	conn, err := tls.Dial("tcp", connStr, &cfg)
	if err != nil {
		logger.Error.Printf("Failed to connect: %s\n", err.Error())
		return t, nil
	}
	logger.Info.Printf("Certificate chain sent by the server: %s:%d\n", t.ept.Fqdn, t.ept.Port)
	for _, cert := range conn.ConnectionState().PeerCertificates {
		logger.Info.Printf("CN: %s OU: %s ValidNotBefore: %s ValidNotAfter: %s\n", cert.Subject.CommonName, cert.Subject.Organization, cert.NotBefore, cert.NotAfter)
	}
	conn.Close()
	return t, nil
}
