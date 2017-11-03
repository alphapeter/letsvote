package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/alphapeter/filecommander/server/cfg"
	"github.com/alphapeter/filecommander/server/gui"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Vote struct {
	Score  int    `json:"score"`
	User   string `json:"user"`
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
	Id          string     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CreatedBy   User     `json:"created_by"`
	Options     []Option `json:"options"`
	Votes       []Vote   `json:"votes"`
	HasEnded    bool     `json:"has_ended"`
	Winner      User     `json:"winner"`
}

type response struct {
	Success bool `json:"success"`
	Reason string `json:"reason"`
}

var db *sql.DB;
func main() {
	settings := cfg.GetSettings()
	db, _ = sql.Open("sqlite3", "./letsvote.sqlite")
	defer db.Close()
	assertTablesCreated(db)

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
	mux.POST("/api/polls/:id/options/", addOption)

	err := http.ListenAndServe(settings.Binding, mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func assertTablesCreated(db *sql.DB) {
	for _, creation := range tableCreations {
		rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table' AND name='` + creation.tableName + `';`)
		if err != nil {
			panic(err.Error())
		}

		if !rows.Next() {
			_, err := db.Exec(creation.sql)
			checkErr(err)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
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

	p.CreatedBy = User{Id: "peter@stratsys.se"}
	p.Id = uuid.NewV4().String()

	stmt, _ := db.Prepare("insert into polls(id, name, description, createdBy) values(?,?,?,?)")
	_, err := stmt.Exec(p.Id, p.Name, p.Description, p.CreatedBy.Id)

	if err != nil {
		response := struct {
			Success bool `json:"success"`
			Reason string `json:"reason"`
		}{ false, err.Error()}
		a, _ := json.Marshal(response)
		w.Write([]byte(a))
		return
	}
	stmt.Close()
	response := struct {
		Success bool `json:"success"`
		Id string `json:"id"`
	}{ true, p.Id}

	a, _ := json.Marshal(response)
	w.Write([]byte(a))
}

func updatePoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/javascript")
}

func getPolls(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	row, err := db.Query("select id, name, description, createdBy, hasEnded, winner from polls;")
	if err != nil {
		res, _ := json.Marshal(response{false, err.Error()})
		w.Write([]byte(res))
		return
	}
	polls := []Poll{}
	for row.Next() {
		p := Poll{}
		row.Scan(&p.Id, &p.Name, &p.Description, &p.CreatedBy.Id, &p.HasEnded, &p.Winner.Id)
		polls = append(polls, p)
	}
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

type tableCreation struct {
	tableName string
	sql       string
}

var tableCreations = []tableCreation{
	{
		tableName: "users",
		sql: `	CREATE TABLE users
				(
					id TEXT NOT NULL,
					name TEXT NOT NULL
				);`,
	},
	{
		tableName: "polls",
		sql: ` CREATE TABLE polls
				(
					id TEXT PRIMARY KEY DEFAULT "" NOT NULL,
					name TEXT DEFAULT "" NOT NULL,
					description TEXT NOT NULL,
					createdBy TEXT NOT NULL,
					hasEnded INTEGER DEFAULT 0 NOT NULL,
					winner TEXT DEFAULT "" NOT NULL
				);
				CREATE UNIQUE INDEX polls_name_uindex ON polls (name);
				`,
	},
	{
		tableName: "options",
		sql: `CREATE TABLE options
				(
					id int NOT NULL,
					pollId TEXT NOT NULL,
					name TEXT DEFAULT "" NOT NULL,
					description TEXT NOT NULL,
					createdBy TEXT NOT NULL,
					CONSTRAINT table_name_id_pollId_pk PRIMARY KEY (id, pollId),
					CONSTRAINT options_polls_id_fk FOREIGN KEY (pollId) REFERENCES polls (id),
					CONSTRAINT options_users_id_fk FOREIGN KEY (createdBy) REFERENCES users (id)
				);`,
	},
	{
		tableName: "votes",
		sql: `	CREATE TABLE votes
				(
					score INT NOT NULL,
					user TEXT NOT NULL,
					optionId INT NOT NULL,
					CONSTRAINT votes_users_id_fk FOREIGN KEY (user) REFERENCES users (id),
					CONSTRAINT votes_users_id_fk FOREIGN KEY (optionId) REFERENCES users (options)

				);`,
	},
}
