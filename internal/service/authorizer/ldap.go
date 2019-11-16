package authorizer

import (
	"crypto/tls"
	"gopkg.in/ldap.v2"
)

type ldapAuthorizer struct {
  Address string `json:"address"`
  Domain  string `json:"domain"`
}

func newLDAP() *ldapAuthorizer {
  return &ldapAuthorizer {
    Address: "ad.upr.edu.cu:636",
    Domain: "@upr.edu.cu",
  }
}

func (auth *ldapAuthorizer) CheckCredentials(username string, password string) (error) {
	conn, err := ldap.DialTLS(
		"tcp",
		auth.Address,
		&tls.Config{
			InsecureSkipVerify: true,
		},
	)
  defer conn.Close()
	if err != nil {
		return err
	}
	err = conn.Bind(username+auth.Domain, password)
	if err != nil {
		return err
	}
  return nil
}

func (auth *ldapAuthorizer) CheckToken(token string) (error) {
  return CheckToken(token)
}

func (auth *ldapAuthorizer) GenerateToken(username string, role string) (string) {
  return GenerateToken(username, role)
}

func (auth *ldapAuthorizer) GetUsername(token string) (string, error) {
  return GetUsername(token)
}
