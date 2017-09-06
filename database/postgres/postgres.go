package postgres

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"sync"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/loger"
)

var (
	once sync.Once
)

func GetPostgresConnection() (db *gorm.DB, err error) {
	once.Do(func() {
		db, err = gorm.Open("postgres", "user=root dbname=telegramgamedb sslmode=disable password=123")
		if err != nil {
			loger.Log.Errorf("Error during connection to Postgres has occurred %s", err.Error())
		}else {
			loger.Log.Info("connection to Pg has been established")
		}
	})
	return db, err
}
