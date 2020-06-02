package certcheck

import (
	"fmt"

	"github.com/piotrszlenk/ssl-test/pkg/endpoint"
	"github.com/piotrszlenk/ssl-test/pkg/logz"
	"github.com/spacemonkeygo/openssl"
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
		logger.Info.Printf("Testing %s\n", i.ept.Fqdn)
		i.testCertificateChain(t.caPath, t.logger)
	}
}

func (t *TestTarget) testCertificateChain(caPath *string, logger logz.LogHandler) (*TestTarget, error) {
	ctx, err := openssl.NewCtx()
	if err != nil {
		logger.Error.Printf("Unable to create SSL context for %s\n", t)
	}
	err = ctx.LoadVerifyLocations(*caPath, "")
	if err != nil {
		logger.Error.Printf("Unable to load CA verify location.\n")
	}

	connStr := fmt.Sprintf("%s:%d", t.ept.Fqdn, t.ept.Port)
	logger.Info.Printf("Connecting to %s\n", connStr)

	var flags openssl.DialFlags
	flags = openssl.InsecureSkipHostVerification
	conn, err := openssl.Dial("tcp", connStr, ctx, flags)
	logger.Debug.Print(conn)
	if err != nil {
		logger.Error.Print(err)
		return t, err
	}

	cert, err := conn.PeerCertificate()
	name, _ := cert.GetSubjectName()
	logger.Debug.Print(name)
	return new(TestTarget), nil
}
