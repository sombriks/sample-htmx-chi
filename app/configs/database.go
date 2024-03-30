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

func (config *Config) InitSchema() {
	config.Db.Exec(`
		create table if not exists todos(
			id integer primary key,
			description text not null,
			done boolean default false,
			created timestamp default current_timestamp
		);
	`)
}
