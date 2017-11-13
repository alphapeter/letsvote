package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"github.com/alphapeter/letsvote/server/users"
)

var office365 Office365Auth
var sessions map[string] users.Session

func Init() {
	office365, _ = CreateOffice365Auth()
}

func getSession() {

}

func HasValidUser(c *gin.Context) {
	if _, err := c.Cookie("letsVoteSession"); err != nil {
		c.SetCookie("letsVoteSession", uuid.NewV4().String(), 800000, "/", "", true, true)
		c.JSON(http.StatusOK, userSessionResponse{
			LoggedIn: false,
		})
		return
	}
	c.JSON(http.StatusOK, userSessionResponse{
		LoggedIn: true,
	})
}

func LoginHandler(c *gin.Context) {
	provider := c.Params.ByName("provider")
	switch provider {
	case "office365":
		office365.Login("state", c.Writer, c.Request)
	}
}

func CallbackHandler(c *gin.Context) {
	provider := c.Params.ByName("provider")
	switch provider {
	case "office365":
		office365.AuthResponse("state", c.Writer, c.Request)
	}

}
