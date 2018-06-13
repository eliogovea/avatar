package auth

import (
	"crypto/tls"
	"errors"
	"net/http"

	"gopkg.in/ldap.v2"
)

var EmptyUsername = errors.New("Empty username")
var EmptyPassword = errors.New("Empty password")

type Auth struct {
	Address     string   `json:"address"` // IP:636
	Domain      string   `json:"domain"`  // @...
	AdminsGroup string   `json:"admins_group"`
	Admins      []string `json:"admins"`
}

func New() *Auth {
	return new(Auth)
}

func (a *Auth) GetUsername(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}
	username := r.FormValue("username")
	if username == "" {
		return username, EmptyUsername
	}
	return username, nil
}

func (a *Auth) GetPassword(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}
	password := r.FormValue("password")
	if password == "" {
		return password, EmptyUsername
	}
	return password, nil
}

func (a *Auth) CheckUserAndPass(username, password string) (error, bool) {
	if username == "" {
		return EmptyUsername, false
	}
	if username == "" {
		return EmptyPassword, false
	}
	conn, err := ldap.DialTLS(
		"tcp",
		a.Address,
		&tls.Config{
			InsecureSkipVerify: true,
		},
	)
	if err != nil {
		return err, false
	}
	err = conn.Bind(username+a.Domain, password)
	if err != nil {
		return err, false
	}

	// TODO use a map
	isManager := false
	for _, admin := range a.Admins {
		if username == admin {
			isManager = true
		}
	}

	return nil, isManager
}
