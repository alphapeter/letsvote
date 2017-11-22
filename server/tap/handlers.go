package tap

import (
	"encoding/json"
	"github.com/alphapeter/letsvote/server/auth"
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/users"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"net/http"
	"sync"
)

var socket *melody.Melody

var connectionStatus struct {
	ConnectedUsers map[users.User]map[*melody.Session]bool
	Lock           *sync.Mutex
	Sessions       map[*melody.Session]users.User
	anonymousUsers int
}

type Message struct {
	Event   string
	Payload interface{}
}

func Init(router *gin.Engine) {
	socket = melody.New()
	connectionStatus.Sessions = make(map[*melody.Session]users.User)
	connectionStatus.ConnectedUsers = make(map[users.User]map[*melody.Session]bool)
	connectionStatus.Lock = new(sync.Mutex)

	router.GET("/tap", func(c *gin.Context) {
		socket.HandleRequest(c.Writer, c.Request)
	})

	socket.HandleConnect(func(s *melody.Session) {
		user := getUser(s)
		UserConnected(user, s)

	})

	socket.HandleDisconnect(func(s *melody.Session) {
		UserDisconnected(s)


	})

}
func UserDisconnected(session *melody.Session) {
	connectionStatus.Lock.Lock()

	user := connectionStatus.Sessions[session]
	delete(connectionStatus.Sessions, session)
	delete(connectionStatus.ConnectedUsers[user], session)
	if len(connectionStatus.ConnectedUsers[user]) == 0 {
		delete(connectionStatus.ConnectedUsers, user)
		go broadcast("USER_DISCONNECT", user)
	}
	connectionStatus.Lock.Unlock()
}

func UserConnected(user users.User, s *melody.Session) {
	connectionStatus.Lock.Lock()

	connectionStatus.Sessions[s] = user
	if connectionStatus.ConnectedUsers[user] == nil {
		connectionStatus.ConnectedUsers[user] = make(map[*melody.Session]bool)
		go broadcast("USER_CONNECT", user)
	}
	connectionStatus.ConnectedUsers[user][s] = true

	connectionStatus.Lock.Unlock()
}
func getUser(s *melody.Session) users.User {
	sessionCookie, err := s.Request.Cookie(auth.SessionCookieName)
	if err != nil {
		return users.User{Id: "anonymous"}
	}
	session := users.Session{}
	err = config.DB.First(&session, "id = ?", sessionCookie.Value).Error
	if err != nil { // session valid
		return users.User{Id: "anonymous"}

	}
	user := users.User{}
	if err = config.DB.First(&user, "id = ?", session.UserId).Error; err != nil {
		return users.User{Id: "anonymous"}
	}
	return user
}

func broadcast(event string, payload interface{}) {
	m := Message{
		Event:   event,
		Payload: payload,
	}

	if message, err := json.Marshal(m); err == nil {
		socket.Broadcast([]byte(message))
	}
}

func GetConnectedUsers(c *gin.Context) {
	var users []users.User
	for k, _ := range connectionStatus.ConnectedUsers {
		users = append(users, k)
	}
	c.JSON(http.StatusOK, users)
}