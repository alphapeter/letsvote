package auth

import (
	"context"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
}

type userSessionResponse struct {
	LoggedIn bool `json:"logged_in"`
}

type Office365claims struct {
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	IpAddress  string `json:"ipaddr"`
	Name       string `json:"name"`
	Oid        string `json:"oid"`
	OnpremSid  string `json:"onprem_sid"`
	Sub        string `json:"sub"`
	Tid        string `json:"tid"`
	UniqueName string `json:"unique_name"`
	Upn        string `json:"upn"`
	Version    string `json:"ver"`
}

type Office365Auth struct {
	ctx      context.Context
	verifier *oidc.IDTokenVerifier
	config   oauth2.Config
	states   map[string]State
}

type State struct {
	user       users.User
	isLoggedIn bool
}
