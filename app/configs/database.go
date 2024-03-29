package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbConfig struct {
	Db *gorm.DB
}

func (config *Config) Open() {
	var err error
	dburl, set := os.LookupEnv("DB_URL")
	if !set {
		dburl = "todos.db"
	}
	config.Db, err = gorm.Open(sqlite.Open(dburl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
