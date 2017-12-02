package polls

import (
	"database/sql"
	"github.com/alphapeter/letsvote/server/config"
	"github.com/alphapeter/letsvote/server/users"
	"time"
)

func InitModels() {
	config.DB.AutoMigrate(&Poll{}, &Option{}, &Vote{})
}

type Status int

const (
	REGISTRATING = Status(0)
	VOTING       = Status(5)
	COUNTING     = Status(8)
	ENDED        = Status(10)
)

type VoteDto struct {
	Score1 string `json:"score_1"`
	Score2 string `json:"score_2"`
	Score3 string `json:"score_3"`
}

type Vote struct {
	UserId         string         `gorm:"primary_key" sql:"type:text REFERENCES users(id)"`
	PollId         string         `gorm:"primary_key" sql:"type:text REFERENCES polls(id)"`
	Score1OptionId sql.NullString `sql:"type:text REFERENCES options(id)"`
	Score2OptionId sql.NullString `sql:"type:text REFERENCES options(id)"`
	Score3OptionId sql.NullString `sql:"type:text REFERENCES options(id)"`
}

type Option struct {
	Id              string     `json:"id", gorm:"primary_key"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	CreatedBy       users.User `json:"created_by" gorm:"ForeignKey:Id;AssociationForeignKey:CreatedByUserId"`
	Score           int       `json:"score"`
	CreatedByUserId string     `json:"-" sql:"type:text REFERENCES users(id)"`
	PollId          string     `json:"poll_id" sql:"type:text REFERENCES polls(id)"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type Poll struct {
	Id              string         `json:"id" gorm:"primary_key"`
	Name            string         `json:"name" gorm:"unique"`
	Description     string         `json:"description"`
	HasEnded        bool           `json:"has_ended"`
	CreatedByUserId string         `json:"-" sql:"type:text REFERENCES users(id)"`
	CreatedBy       users.User     `json:"created_by" gorm:"ForeignKey:Id;AssociationForeignKey:CreatedByUserId"`
	WinnerUserId    sql.NullString `json:"-" sql:"type:text REFERENCES users(id)"`
	Winner          users.User     `json:"winner" gorm:"ForeignKey:Id;AssociationForeignKey:WinnerUserId"`
	Options         []Option       `json:"options" gorm:"ForeignKey:PollId;AssociationForeignKey:Id"`
	Votes           []Vote         `json:"votes" gorm:"ForeignKey:PollId;AssociationForeignKey:Id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Status          Status         `json:"status"`
}

type userCreated interface {
	getUserId() string
}

func (p Poll) getUserId() string {
	return p.CreatedByUserId
}

func (o Option) getUserId() string {
	return o.CreatedByUserId
}

func NullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{
			Valid: false,
		}
	}
	return sql.NullString{
		Valid:  true,
		String: s,
	}
}
