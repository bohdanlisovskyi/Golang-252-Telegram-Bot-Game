package postgres

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"sync"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
)

var (
	once sync.Once
)

func GetPostgresConnection() (db *gorm.DB, err error) {
	once.Do(func() {
		db, err = gorm.Open("postgres", "host=myhost user=root dbname=chat sslmode=disable password=ssugolang")
		if err != nil {
			loger.Log.Errorf("Error during connection to Postgres has occurred %s", err.Error())
		}
	})
	return db, err
}
