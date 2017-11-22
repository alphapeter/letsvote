package users

import (
	"crypto/md5"
	"fmt"
	"github.com/alphapeter/letsvote/server/config"
	"time"
	"strings"
)

func InitModels() {
	config.DB.AutoMigrate(&User{})
	config.DB.AutoMigrate(&Session{})
}

type User struct {
	Id        string    `json:"id", gorm:"primary_key"`
	Email     string    `json:"-", gorm:"unique"`
	Gravatar  string    `json:"gravatar"`
	Name      string    `json:"name", gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Session struct {
	Id        string `gorm:"primary_key"`
	UserId    string `sql:"type:text REFERENCES users(id)"`
	CreatedAt time.Time
}

func EmailToHash(email string) string {
	hash := md5.Sum([]byte(strings.ToLower(email)))
	return fmt.Sprintf("%x", hash)
}
