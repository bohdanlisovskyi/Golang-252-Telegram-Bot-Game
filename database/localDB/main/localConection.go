package main

import (
	"github.com/jinzhu/gorm"
	"sync"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/loger"
)

var (
	once sync.Once
)

func LocalPostgresConnection() (db *gorm.DB, err error) {
	once.Do(func() {
		db, err = gorm.Open("postgres", "user=root dbname=testdb sslmode=disable password=root")
		if err != nil {
			loger.Log.Errorf("Error during connection to Postgres has occurred %s", err.Error())
		}else {
			loger.Log.Info("connection established")
		}
	})
	return db, err
}
