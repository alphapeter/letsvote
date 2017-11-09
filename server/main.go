package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/alphapeter/letsvote/server/cfg"
	"github.com/alphapeter/letsvote/server/gui"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
	"net/http"
)

type User struct {
	Id   string `json:"id", gorm:"primary_key"`
	Name string `json:"name"", gorm:"type:varchar(50)"`
}

type Vote struct {
	Score    int    `json:"score"`
	UserId   string `gorm:"primary_key" sql:"type:text REFERENCES users(id)"`
	User     string `json:"user", gorm:"ForeignKey:Id;AssociationForeignKey:UserId"`
	OptionId string `gorm:"primary_key"`
	Option   Option `json:"option" gorm:"ForeignKey:Id;AssociationForeignKey:OptionId"`
	PollId   string `sql:"type:text REFERENCES polls(id)"`
}

type Option struct {
	Id              uint   `json:"id", gorm:"primary_key"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	CreatedBy       User   `json:"created_by" gorm:"ForeignKey:Id;AssociationForeignKey:CreatedByUserId"`
	Score           uint   `json:"score"`
	CreatedByUserId string `sql:"type:text REFERENCES users(id)"`
	PollId          string `sql:"type:text REFERENCES polls(id)"`
}

type Poll struct {
	Id              string         `json:"id" gorm:"primary_key"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	HasEnded        bool           `json:"has_ended"`
	CreatedByUserId string         `sql:"type:text REFERENCES users(id)"`
	CreatedBy       User           `json:"created_by" gorm:"ForeignKey:Id;AssociationForeignKey:CreatedByUserId"`
	Winner          User           `json:"winner"`
	WinnerUserId    sql.NullString `sql:"type:text REFERENCES users(id)"`
	Options         []Option       `json:"options" gorm:"ForeignKey:PollId;AssociationForeignKey:Id"`
	Votes           []Vote         `json:"votes" gorm:"ForeignKey:PollId;AssociationForeignKey:Id"`
}
type Session struct {
	Key    string `gorm:"primary_key"`
	UserId string `sql:"type:text REFERENCES users(id)"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
}

type userSessionResponse struct {
	LoggedIn bool `json:"logged_in"`
}

var db *gorm.DB

func main() {
	settings := cfg.GetSettings()
	db, _ = gorm.Open("sqlite3", "test.db")
	db.Exec("PRAGMA foreign_keys = ON")

	db.AutoMigrate(&User{}, &Poll{}, &Option{}, &Vote{}, &Session{})
	defer db.Close()

	background := context.Background()
	auth, _ := createOffice365Auth(background)
	mux := httprouter.New()
	mux.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		cookie := &http.Cookie{Name: "hej", Value: "oooh"}
		http.SetCookie(w, cookie)
		fmt.Println(cookie.String())
		w.Header().Set("Content-Type", "text/html")
		w.Write(gui.Html)
	})

	mux.GET("/logout", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	})

	mux.GET("/static/js/app.js", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(gui.Javascript)
	})

	mux.POST("/api/polls", addPoll) //Auth
	mux.GET("/api/polls", getPolls)

	mux.PUT("/api/polls/:id", updatePoll) //Auth
	mux.GET("/api/polls/:id", getPoll)

	mux.GET("/api/polls/:id/options", getOptions)
	mux.GET("/api/polls/:id/options/:id", getOption)
	mux.POST("/api/polls/:id/options/", addOption) //Auth

	mux.GET("/auth/login", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		c := new(http.Cookie)
		http.SetCookie(w, c)
		auth.login("hej", w, r)
	})

	mux.GET("/auth/o365/callback", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		auth.login("hej", w, r)
	})

	mux.GET("/auth/hasValidUserSession", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := r.Cookie("lv_session"); err != nil {
			c := http.Cookie{
				Name:   "lv_session",
				Value:  uuid.NewV4().String(),
			}
			http.SetCookie(w, &c)
			f, _ := json.Marshal(userSessionResponse{
				LoggedIn: false,
			})
			w.Write([]byte(f))
			return
		}

		s, _ := json.Marshal(userSessionResponse{
			LoggedIn: true,
		})
		w.Write([]byte(s))
	})

	err := http.ListenAndServe(settings.Binding, mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func addPoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var p Poll

	if err := decoder.Decode(&p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error, could not parse json"))
		return
	}

	p.CreatedByUserId = "peter@stratsys.se"
	p.Id = uuid.NewV4().String()
	err := db.Create(&p).Error

	if err != nil {
		response := struct {
			Success bool   `json:"success"`
			Reason  string `json:"reason"`
		}{false, err.Error()}
		a, _ := json.Marshal(response)
		w.Write([]byte(a))
		return
	}
	response := struct {
		Success bool   `json:"success"`
		Id      string `json:"id"`
	}{true, p.Id}

	a, _ := json.Marshal(response)
	w.Write([]byte(a))
}

func updatePoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/javascript")
}

func getPolls(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	polls := []Poll{}
	db.Find(&polls)
	res, _ := json.Marshal(polls)
	w.Write([]byte(res))
}
func getPoll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}

func getOptions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}

func getOption(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}
func addOption(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}
