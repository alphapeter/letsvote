package auth

import (
	"context"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"fmt"
	"github.com/alphapeter/letsvote/server/config"
	"errors"
	"github.com/alphapeter/letsvote/server/users"
)

func CreateOffice365Auth(settings config.OpenIdConnectProvider) (Office365Auth, error) {
	background := context.Background()
	provider, err := oidc.NewProvider(background, settings.IssuerUrl)
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{
		ClientID:          settings.ClientId,
		SkipClientIDCheck: true,
	}
	verifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     settings.ClientId,
		ClientSecret: settings.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  settings.BaseUrl + "/auth/callback/office365",
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return Office365Auth{
		ctx:      background,
		verifier: verifier,
		config:   config,
	}, nil
}

func (auth *Office365Auth) Login(state string, w http.ResponseWriter, r *http.Request) {
	s := auth.config.AuthCodeURL(state)
	http.Redirect(w, r, s, http.StatusFound)
}

func (auth *Office365Auth) AuthResponse(state string, w http.ResponseWriter, r *http.Request) (users.User, error){
	user := users.User{}
	s := r.URL.Query().Get("state")
	if s != state {
		return user, errors.New("state did not match")
	}
	oauth2Token, err := auth.config.Exchange(auth.ctx, r.URL.Query().Get("code"))
	if err != nil {
		return user, errors.New("Failed to exchange token: "+err.Error())
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return user, errors.New("No id_token field in oauth2 token.")
	}
	idToken, err := auth.verifier.Verify(auth.ctx, rawIDToken)
	if err != nil {
		return user, errors.New("Failed to verify ID Token: "+err.Error())
	}

	claims := Office365claims{}
	err = idToken.Claims(&claims)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf(claims.Upn)

	user.Id = claims.Oid
	user.Name = claims.Name
	user.Email = claims.UniqueName

	return user, nil
}
