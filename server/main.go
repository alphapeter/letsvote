package main

import (
	"context"
	"fmt"
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/webui"
	"github.com/alphapeter/letsvote/server/polls"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"github.com/alphapeter/letsvote/server/auth"
)

func main() {
	settings := config.GetSettings()
	config.InitDb()
	users.InitModels()
	polls.InitModels()
	auth.Init()
	defer config.DB.Close()

	router := gin.Default()

	router.GET("/", webui.HtmlHandler)
	router.GET("/static/js/app.js", webui.JsHandler)

	router.POST("/api/polls", polls.AddPoll) //Auth
	router.GET("/api/polls", polls.GetPolls)

	router.PUT("/api/polls/:id", polls.UpdatePoll) //Auth
	router.GET("/api/polls/:id", polls.GetPoll)

	router.GET("/api/polls/:id/options", polls.GetOptions)
	router.GET("/api/polls/:id/options/:id", polls.GetOption)
	router.POST("/api/polls/:id/options/", polls.AddOption) //Auth

	router.GET("/logout", auth.LogoutHandler)
	router.GET("/auth/login", auth.LoginHandler)

	router.GET("/auth/o365/callback", auth.Callback)

	router.GET("/auth/hasValidUserSession", auth.HasValidUser)

	err := http.ListenAndServe(settings.Binding, router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
