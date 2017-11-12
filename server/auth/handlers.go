package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"golang.org/x/net/context"
)


func initAuth()  {
	background := context.Background()

	a, _ = CreateOffice365Auth(background)
}
func HasValidUser(c *gin.Context) {
	if _, err := c.Cookie("lv_session"); err != nil {
		c.SetCookie("lv_session", uuid.NewV4().String(), 800000, "/", "", true, true)
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
	a.Login("state", c.Writer, c.Request)
}
func LogoutHandler(c *gin.Context) {
	a.Login("state", c.Writer, c.Request)
}

func Callback(c *gin.Context) {

}
