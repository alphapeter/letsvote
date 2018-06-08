package main

import (
	"fmt"
	"github.com/alphapeter/letsvote/server/auth"
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/polls"
	"github.com/alphapeter/letsvote/server/tap"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alphapeter/letsvote/server/static"
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

	staticHandler := static.CreateHandler(static.Data, "index.html")

	staticHtmlHandler := func(c *gin.Context) {
		staticHandler.ServeHTTP(c.Writer, c.Request)
	}

	staticResourceHandler := func(c *gin.Context) {
		c.Header("Etag", c.Request.RequestURI)
		c.Header("Cache-Control", "max-age=14515200")
		staticHandler.ServeHTTP(c.Writer, c.Request)
	}

	router.GET("/", staticHtmlHandler)
	router.GET("/admin.html", staticHtmlHandler)

	router.GET("/static/*filename", staticResourceHandler)

	router.GET("/api/activeusers", tap.GetConnectedUsers)

	router.GET("/api/polls", polls.GetPolls)
	router.GET("api/voters", polls.GetVoters)


	authorized.POST("/api/polls", polls.AddPoll)

	authorized.DELETE("/api/polls/:pollId", polls.DeletePoll)
	authorized.PATCH("/api/polls/:pollId", polls.UpdatePoll)
	authorized.PUT("/api/polls/:pollId", polls.UpdatePoll)

	authorized.POST("/api/polls/:pollId/options", polls.AddOption)

	authorized.DELETE("/api/polls/:pollId/options/:optionId", polls.DeleteOption)
	authorized.PATCH("/api/polls/:pollId/options/:optionId", polls.UpdateOption)

	authorized.POST("/api/polls/:pollId/vote", polls.HandleVote)

	authorized.GET("api/votes", polls.GetVotes)

	authorized.GET("api/users", users.GetUsers)
	authorized.PATCH("api/users/:userId", users.SetAdminPermission)

	//router.GET("/auth/fakelogin/:fakeuser", auth.FakeLoginHandler)
	router.GET("/auth/login/:provider", auth.LoginHandler)
	router.GET("/auth/callback/:provider", auth.CallbackHandler)
	router.GET("/auth/logout", auth.LogoutHandler)

	tap.Init(router)
	err := http.ListenAndServe(settings.Binding, router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
