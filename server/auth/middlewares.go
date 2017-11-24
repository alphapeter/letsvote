package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/users"
	"net/http"
)

func CookieAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie(SessionCookieName)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		session := users.Session{}
		err = config.DB.First(&session, "id = ?", sessionId).Error
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		user := users.User{}

		err = config.DB.First(&user, "id = ?", session.UserId).Error
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
