package users

import (
	"database/sql"
	"github.com/alphapeter/letsvote/server/config"
)

func InitModels() {
	config.DB.AutoMigrate(&User{}, Session{})

}

type User struct {
	Id   string `json:"id", gorm:"primary_key"`
	Name string `json:"name"", gorm:"type:varchar(50)"`
}

type Session struct {
	Key    string         `gorm:"primary_key"`
	UserId sql.NullString `sql:"type:text REFERENCES users(id)"`
}
