package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"web20.tk/entries"
)

const dbInfo = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

type Config struct {
	Port int    `json:"port"`
	Host string `json:"host"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

var db *gorm.DB

func Configure(c Config) error {
	err := connect(c)
	if err != nil {
		return err
	}

	autoMigrate()
	return nil
}

func connect(c Config) error {
	var err error
	db, err = gorm.Open("postgres", fmt.Sprintf(dbInfo, c.Host, c.Port, c.User, c.Pass, c.Name))
	if err != nil {
		return err
	}

	return nil
}

func Get() *gorm.DB {
	return db
}

func autoMigrate() {
	db.AutoMigrate(
		&entries.Tag{},
		&entries.Category{},
		&entries.Article{},
	)
}
