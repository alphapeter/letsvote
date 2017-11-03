package main

import (
	"encoding/json"
	"fmt"
	"github.com/alphapeter/filecommander/server/cfg"
	"github.com/alphapeter/filecommander/server/gui"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

type Uuid string
type Email string

type User struct {
	Name string `json:"name"`
	Id   Email  `json:"id"`
}

type Vote struct {
	Score  int    `json:"score"`
	User   Uuid   `json:"user"`
	Option Option `json:"option"`
}

type Option struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   User   `json:"created_by"`
	Score       int    `json:"score"`
}

type Poll struct {
	Id          Uuid     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CreatedBy   User     `json:"created_by"`
	Options     []Option `json:"options"`
	Votes       []Vote   `json:"votes"`
	HasEnded    bool     `json:"has_ended"`
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

	mux.POST("/api/polls", addPoll)
	mux.GET("/api/polls", getPolls)

	mux.PUT("/api/polls/:id", updatePoll)
	mux.GET("/api/polls/:id", getPoll)

	mux.GET("/api/polls/:id/options", getOptions)
	mux.GET("/api/polls/:id/options/:id", getOption)

	err := http.ListenAndServe(settings.Binding, mux)
	if err != nil {
		fmt.Println(err.Error())
	}

}

var polls = make(map[Uuid]Poll)

func addPoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	decoder := json.NewDecoder(r.Body)
	var p Poll

	if err := decoder.Decode(&p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error, could not parse json"))
		return
	}

	p.Id = "dfsdf"
	polls[p.Id] = p

	w.Header().Set("Content-Type", "application/json")
	w.Write(gui.Javascript)
}

func updatePoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(gui.Javascript)
}

func getPolls(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	poll := createPoll()

	// upprepar sig på alla svar, fixa metod
	polls := []Poll{poll, poll, poll, poll, poll}
	j, _ := json.Marshal(polls)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
var id = 1
func createPoll() Poll {
	uuid := fmt.Sprint("dsdf%d", id)
	id++
	return Poll{
		Id: Uuid(uuid),
		CreatedBy: User{
			Id:   "peter@klaesson.net",
			Name: "324234",
		},
		Description: "best day ever",
		HasEnded:    false,
		Name:        "Innovationday 324",
		Votes:       []Vote{},
		Options: []Option{
			Option{
				Name:        "Die Welle",
				Id:          1,
				Description: "Autokrati",
				CreatedBy: User{
					Id:   "user1",
					Name: "Alfons Åberg",
				},
			},
			Option{
				Name:        "Das experiment",
				Id:          2,
				Description: "",
				CreatedBy: User{
					Id:   "user2",
					Name: "Riddar Kato",
				},
			},
		},
	}
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
