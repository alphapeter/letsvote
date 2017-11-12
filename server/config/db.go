package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func InitDb() {
	DB, _ = gorm.Open("sqlite3", "test.db")
	DB.Exec("PRAGMA foreign_keys = ON")
}
