package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/alphapeter/filecommander/server/cfg"
	"github.com/alphapeter/filecommander/server/gui"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

type uuid string
type email string

type user struct {
	name string
	id email
}

type vote struct {
	score int
	user uuid
	option option
}

type option struct{
	id int
	name string
	description string
	createdBy user
	score int
}

type poll struct{
	id uuid
	name string
	description string
	createdBy user
	options []option
	votes []vote
	hasEnded bool
}

func main() {
	settings := cfg.GetSettings()

	mux := httprouter.New()
	mux.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(gui.Html)
	})

	mux.GET("/static/js/app.js", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(gui.Javascript)
	})

	//mux.POST("/login", login)

	mux.POST("/polls", addPoll)
	mux.GET("/polls", getPolls)

	mux.PUT("/polls/:id", updatePoll)
	mux.GET("/polls/:id", getPoll)

	mux.GET("/polls/:id/options", getOptions)
	mux.GET("/polls/:id/options/:id", getOption)

	err := http.ListenAndServe(settings.Binding, mux)
	if err != nil {
		fmt.Println(err.Error())
	}

}

var polls = make(map[uuid]poll)

func addPoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	decoder := json.NewDecoder(r.Body)
	var p poll

	if err := decoder.Decode(&p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error, could not parse json"))
		return
	}

	p.id = "dfsdf"
	polls[p.id] = p

	w.Header().Set("Content-Type", "application/json")
	w.Write(gui.Javascript)
}

func updatePoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(gui.Javascript)
}

func getPolls(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(gui.Javascript)
}

func getPoll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(gui.Javascript)
}

func getOptions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(gui.Javascript)
}

func getOption(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(gui.Javascript)
}
