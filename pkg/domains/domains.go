package domains

import (
	"crypto/tls"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	"time"

	"github.com/lechuhuuha/oneshield/constant"
)

type SSLResponse struct {
	Domain  string `json:"domain"`
	IsSSL   bool   `json:"is_ssl"`
	Expire  string `json:"expire"`
	Message string `json:"message"`
	// not before
	SSLValidAfter string `json:"ssl_valid_after"`
	// not after
	SSLExpireAfter string    `json:"ssl_expire_after"`
	Issuer         pkix.Name `json:"issuer"`
	Subject        pkix.Name `json:"subject"`
	DNSNames       []string  `json:"dns_names"`
	SerialNumber   *big.Int  `json:"serial_number"`
}

func NewDomainSSL(domainName string) *SSLResponse {
	return &SSLResponse{
		Domain: domainName,
		IsSSL:  false,
	}
}

func (s *SSLResponse) TestDomainSSL(domainName string) *SSLResponse {
	domainSSL := NewDomainSSL(domainName)
	nDialer := &net.Dialer{
		Timeout: time.Duration(10) * time.Second,
	}
	conn, err := tls.DialWithDialer(nDialer, "tcp", fmt.Sprintf("%s:%d", domainName, constant.SSLPort), nil)
	if err != nil {
		domainSSL.Message = fmt.Sprintf("Server doesn't support SSL certificate err: " + err.Error())
		return domainSSL
	}
	if len(conn.ConnectionState().PeerCertificates) == 0 {
		return domainSSL
	}
	err = conn.VerifyHostname(domainName)
	if err != nil {
		domainSSL.Message = fmt.Sprintf("Hostname doesn't match with certificate: " + err.Error())
		return domainSSL
	}

	domainSSL.IsSSL = true
	domainSSL.SSLExpireAfter = conn.ConnectionState().PeerCertificates[0].NotAfter.Format(time.DateTime)
	domainSSL.SSLValidAfter = conn.ConnectionState().PeerCertificates[0].NotBefore.Format(time.DateTime)
	domainSSL.Issuer = conn.ConnectionState().PeerCertificates[0].Issuer
	domainSSL.SerialNumber = conn.ConnectionState().PeerCertificates[0].SerialNumber
	domainSSL.Subject = conn.ConnectionState().PeerCertificates[0].Subject
	domainSSL.DNSNames = conn.ConnectionState().PeerCertificates[0].DNSNames

	return domainSSL
}
