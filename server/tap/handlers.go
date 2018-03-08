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
	"time"
)

var socket *melody.Melody

type connectionInfo struct {
	user     users.User
	sessions map[*melody.Session]bool
}

var connectionStatus struct {
	ConnectedUsers map[string]*connectionInfo
	Lock           *sync.Mutex
	Sessions       map[*melody.Session]users.User
	anonymousUsers int
}

type Message struct {
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
}

func Init(router *gin.Engine) {
	socket = melody.New()
	connectionStatus.Sessions = make(map[*melody.Session]users.User)
	connectionStatus.ConnectedUsers = make(map[string]*connectionInfo)
	connectionStatus.Lock = new(sync.Mutex)

	heartbeat := time.NewTicker(20 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-heartbeat.C:
				Broadcast("heartbeat", "")
			case <-quit:
				heartbeat.Stop()
				return
			}
		}
	}()

	router.GET("/tap", func(c *gin.Context) {
		socket.HandleRequest(c.Writer, c.Request)
	})

	socket.HandleConnect(func(s *melody.Session) {
		user := getUser(s)
		UserConnected(user, s)
		users := GetConnectedUserList()
		message := Message{
			Event:   "CONNECTED_USERS",
			Payload: users,
		}
		if json, err := json.Marshal(message); err == nil {
			s.Write(json)
		}

	})

	socket.HandleDisconnect(func(s *melody.Session) {
		UserDisconnected(s)
	})
}
func UserDisconnected(session *melody.Session) {
	connectionStatus.Lock.Lock()

	user := connectionStatus.Sessions[session]
	delete(connectionStatus.Sessions, session)
	delete(connectionStatus.ConnectedUsers[user.Id].sessions, session)
	if len(connectionStatus.ConnectedUsers[user.Id].sessions) == 0 {
		delete(connectionStatus.ConnectedUsers, user.Id)
		go Broadcast("USER_DISCONNECT", user)
	}
	connectionStatus.Lock.Unlock()
}

func UserConnected(user users.User, s *melody.Session) {
	connectionStatus.Lock.Lock()

	connectionStatus.Sessions[s] = user
	if connectionStatus.ConnectedUsers[user.Id] == nil {
		connectionStatus.ConnectedUsers[user.Id] = &connectionInfo{
			user:     user,
			sessions: make(map[*melody.Session]bool),
		}

		go Broadcast("USER_CONNECT", user)
	}
	connectionStatus.ConnectedUsers[user.Id].sessions[s] = true

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

func Broadcast(event string, payload interface{}) {
	m := Message{
		Event:   event,
		Payload: payload,
	}

	if message, err := json.Marshal(m); err == nil {
		socket.Broadcast([]byte(message))
	}
}

func GetConnectedUsers(c *gin.Context) {
	users := GetConnectedUserList()
	c.JSON(http.StatusOK, users)
}

func GetConnectedUserList() []users.User {
	var users []users.User
	for _, v := range connectionStatus.ConnectedUsers {
		users = append(users, v.user)
	}
	return users
}
