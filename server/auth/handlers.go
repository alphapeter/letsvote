package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/alphapeter/letsvote/server/config"
	"encoding/json"
	"net/url"
)

var office365 Office365Auth
var sessions map[string] users.Session

func Init(providers []config.OpenIdConnectProvider) {
	for _, provider := range providers {
		switch provider.Type {
		case "office365":
			office365, _ = CreateOffice365Auth(provider)
		}
	}
}

func LoginHandler(c *gin.Context) {
	provider := c.Params.ByName("provider")
	switch provider {
	case "office365":
		sessionId := newSessionId()
		c.SetCookie(SessionCookieName, sessionId, 600, "/", "",false, true)
		c.SetCookie(UserCookieName, "true", -1, "/", "",false, false)
		office365.Login(sessionId, c.Writer, c.Request)
		break
	default:
		c.String(http.StatusNotFound, "provider " + provider + "not found")
	}
}

func LogoutHandler(c *gin.Context) {
	if sessionId, err := c.Cookie(SessionCookieName); err == nil {
		users.DeleteSession(sessionId)
	}
	c.SetCookie(SessionCookieName, "", -1, "/", "",false, true)
	c.SetCookie(UserCookieName, "", -1, "/", "",false, false)
	c.Redirect(http.StatusFound, "/")
}


func CallbackHandler(ctx *gin.Context) {
	provider := ctx.Params.ByName("provider")
	sessionId, err := ctx.Cookie(SessionCookieName)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid request, no session established or the login took more than 10 minutes")
	}
	var authUser users.User
	switch provider {
	case "office365":
		authUser, err = office365.AuthResponse(sessionId, ctx.Writer, ctx.Request)
	}
	if err != nil {
		invalidateSession(ctx, sessionId, err)
		return
	}

	var user users.User
	user, err = users.GetOrCreateUser(authUser)

	if err != nil {
		invalidateSession(ctx, sessionId, err)
		return
	}

	ctx.SetCookie(SessionCookieName, sessionId, cookieMaxAge, "/", "", false, true)
	userResponse := userToJson(user)
	ctx.SetCookie(UserCookieName, userResponse, cookieMaxAge, "/", "", false, false)
	err = users.SetSession(user, sessionId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}
func userToJson(user users.User) string {
	a, _ := json.Marshal(user)
	s := string(a)
	userResponse := url.PathEscape(s)
	return userResponse
}

func invalidateSession(ctx *gin.Context, sessionId string, err error) {
	ctx.SetCookie(SessionCookieName, sessionId, -1, "/", "", false, true)
	ctx.SetCookie(UserCookieName, "true", -1, "/", "", false, false)
	ctx.String(http.StatusBadRequest, err.Error())
}

func newSessionId() string {
	return uuid.NewV4().String()
}
