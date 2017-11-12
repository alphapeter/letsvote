package auth

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)


func Init() {

}
func CreateOffice365Auth(ctx context.Context) (Office365Auth, error) {
	provider, err := oidc.NewProvider(ctx, "")
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{
		ClientID:          "",
		SkipClientIDCheck: true,
	}
	verifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:5556/auth/o365/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return Office365Auth{
		ctx:      ctx,
		verifier: verifier,
		config:   config,
	}, nil
}

//http.HandleFunc("/",
func (auth *Office365Auth) Login(state string, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, auth.config.AuthCodeURL(state), http.StatusFound)
}

//http.HandleFunc("/auth/o365/callback",
func (auth *Office365Auth) AuthResponse(state string, w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("state") != state {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	oauth2Token, err := auth.config.Exchange(auth.ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}
	idToken, err := auth.verifier.Verify(auth.ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	v := Office365claims{}
	err = idToken.Claims(&v)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf(v.Upn)

	//set cookie
	w.Write([]byte(v.Name))
	//http.Redirect(w, r, "/", http.StatusFound)

	//(data)
}
