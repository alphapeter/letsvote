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

type Vote struct {
	UserId         string `json:"user_id" gorm:"primary_key" sql:"type:text REFERENCES users(id)"`
	PollId         string `json:"poll_id" gorm:"primary_key" sql:"type:text REFERENCES polls(id)"`
	Score1OptionId string `json="score_1" sql:"type:text REFERENCES options(id)"`
	Score2OptionId string `json="score_2" sql:"type:text REFERENCES options(id)"`
	Score3OptionId string `json="score_3" sql:"type:text REFERENCES options(id)"`
}

type Option struct {
	Id              string     `json:"id", gorm:"primary_key"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	CreatedBy       users.User `json:"created_by" gorm:"ForeignKey:Id;AssociationForeignKey:CreatedByUserId"`
	Score           uint       `json:"score"`
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
}
