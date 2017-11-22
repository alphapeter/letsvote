package main

import (
	"fmt"
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/webui"
	"github.com/alphapeter/letsvote/server/polls"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alphapeter/letsvote/server/auth"
	"github.com/alphapeter/letsvote/server/tap"
)

func main() {
	settings := config.GetSettings()
	config.InitDb()
	users.InitModels()
	polls.InitModels()
	auth.Init(settings.OpenIdConnectProviders)
	defer config.DB.Close()

	router := gin.Default()
	authorized := router.Group("/", auth.CookieAuth())

	router.GET("/", webui.HtmlHandler)
	router.GET("/static/js/app.js", webui.JsHandler)

	router.GET("/api/activeusers", tap.GetConnectedUsers)

	authorized.POST("/api/polls", polls.AddPoll)
	router.GET("/api/polls", polls.GetPolls)

	authorized.PUT("/api/polls/:id", polls.UpdatePoll)
	router.GET("/api/polls/:id", polls.GetPoll)

	router.GET("/api/polls/:id/options", polls.GetOptions)
	router.GET("/api/polls/:id/options/:id", polls.GetOption)
	authorized.POST("/api/polls/:id/options", polls.AddOption)

	router.GET("/auth/login/:provider", auth.LoginHandler)
	router.GET("/auth/callback/:provider", auth.CallbackHandler)
	router.GET("/auth/logout", auth.LogoutHandler)

	tap.Init(router)
	err := http.ListenAndServe(settings.Binding, router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
