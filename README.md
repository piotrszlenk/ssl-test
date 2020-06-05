# SSL-TEST

SSL-TEST is a simple tool written in Golang to perform quick scan on SSL endpoints and pull certificate chains sent by them.

## Usage

```./main -endpoints-file example.txt -capath "/usr/local/etc/openssl/cert.pem"```

## Endpoint list

Tool reads a flat text CSV file that contains list of TCP SSL endpoints with port number. For sample endpoint list file, check `example.txt`.

## Sample test output

```
$ ./main -endpoints-file example.txt -capath "/usr/local/etc/openssl/cert.pem"
INFO: 2020/06/05 19:53:01 main.go:32: Loading endpoints from: example.txt
INFO: 2020/06/05 19:53:01 main.go:40: Creating test targets from loaded endpoints.
INFO: 2020/06/05 19:53:01 certcheck.go:48: Testing endpoint google.com on port 443
INFO: 2020/06/05 19:53:01 certcheck.go:48: Testing endpoint gmail.com on port 443
INFO: 2020/06/05 19:53:01 certcheck.go:48: Testing endpoint cnn.com on port 443
INFO: 2020/06/05 19:53:01 certcheck.go:48: Testing endpoint yahoo.com on port 443
INFO: 2020/06/05 19:53:01 certcheck.go:48: Testing endpoint feedly.com on port 443
INFO: 2020/06/05 19:53:01 certcheck.go:104: Certificate chain sent by the endpoint: google.com:443
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: *.google.com OU: [Google LLC] ValidNotBefore: 2020-05-20 12:00:48 +0000 UTC ValidNotAfter: 2020-08-12 12:00:48 +0000 UTC
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: GTS CA 1O1 OU: [Google Trust Services] ValidNotBefore: 2017-06-15 00:00:42 +0000 UTC ValidNotAfter: 2021-12-15 00:00:42 +0000 UTC
INFO: 2020/06/05 19:53:01 certcheck.go:104: Certificate chain sent by the endpoint: cnn.com:443
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: turner-tls.map.fastly.net OU: [Fastly, Inc.] ValidNotBefore: 2020-05-05 20:11:42 +0000 UTC ValidNotAfter: 2021-05-06 20:11:42 +0000 UTC
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: GlobalSign CloudSSL CA - SHA256 - G3 OU: [GlobalSign nv-sa] ValidNotBefore: 2015-08-19 00:00:00 +0000 UTC ValidNotAfter: 2025-08-19 00:00:00 +0000 UTC
INFO: 2020/06/05 19:53:01 certcheck.go:104: Certificate chain sent by the endpoint: gmail.com:443
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: gmail.com OU: [Google LLC] ValidNotBefore: 2020-05-20 12:13:15 +0000 UTC ValidNotAfter: 2020-08-12 12:13:15 +0000 UTC
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: GTS CA 1O1 OU: [Google Trust Services] ValidNotBefore: 2017-06-15 00:00:42 +0000 UTC ValidNotAfter: 2021-12-15 00:00:42 +0000 UTC
INFO: 2020/06/05 19:53:01 certcheck.go:104: Certificate chain sent by the endpoint: yahoo.com:443
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: *.www.yahoo.com OU: [Oath Inc] ValidNotBefore: 2020-05-11 00:00:00 +0000 UTC ValidNotAfter: 2020-11-07 12:00:00 +0000 UTC
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: DigiCert SHA2 High Assurance Server CA OU: [DigiCert Inc] ValidNotBefore: 2013-10-22 12:00:00 +0000 UTC ValidNotAfter: 2028-10-22 12:00:00 +0000 UTC
INFO: 2020/06/05 19:53:01 certcheck.go:104: Certificate chain sent by the endpoint: feedly.com:443
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: *.feedly.com OU: [] ValidNotBefore: 2018-02-06 00:00:00 +0000 UTC ValidNotAfter: 2021-05-01 12:00:00 +0000 UTC
INFO: 2020/06/05 19:53:01 certcheck.go:106: CN: RapidSSL RSA CA 2018 OU: [DigiCert Inc] ValidNotBefore: 2017-11-06 12:23:33 +0000 UTC ValidNotAfter: 2027-11-06 12:23:33 +0000 UTC
```

