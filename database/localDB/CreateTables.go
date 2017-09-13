package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/loger"
	"sync"
	"github.com/jinzhu/gorm"
)

var (
	once sync.Once
)

func Conn() (db *gorm.DB, err error) {
	once.Do(func() {
		db, err = gorm.Open("postgres", "user=root dbname=testdb sslmode=disable password=root")
		if err != nil {
			loger.Log.Errorf("Error during connection to Postgres has occurred %s", err.Error())
		} else {
			loger.Log.Info("connection to TEST DATABASE established")
		}
	})
	return db, err
}

type Planet struct {
	Id          int    `gorm:"primary_key;AUTO_INCREMENT"`
	UserName    string `gorm:"not null;unique"`
	PlanetName  string `gorm:"not null;unique"`
	XCoordinate int
	YCoordinate int
	IsActive    bool
	IsLoaded    bool
}

type Citycenter struct {
	Id              int    `gorm:"primary_key;AUTO_INCREMENT"`
	PlanetName      string `gorm:"not null;unique"`
	Level           int
	PeopleAmount    int
	PeopleMaxAmount int
	CitycenterKPI   int
}

type IronMine struct {
	Id              int    `gorm:"primary_key;AUTO_INCREMENT"`
	PlanetName      string `gorm:"not null;unique"`
	Level           int
	MineKPI         int
	PeopleAmount    int
	PeopleMaxAmount int
}

type CrystalMine struct {
	Id              int    `gorm:"primary_key;AUTO_INCREMENT"`
	PlanetName      string `gorm:"not null;unique"`
	Level           int
	MineKPI         int
	PeopleAmount    int
	PeopleMaxAmount int
}

type Cosmodrome struct {
	Id              int    `gorm:"primary_key;AUTO_INCREMENT"`
	PlanetName      string `gorm:"not null;unique"`
	Level           int
	CosmodrKPI      int
	ShipAmount      int
	ShipMaxAmount   int
	FleetsAmount    int
	FleetsMaxamount int
}

type Ship struct {
	Id           int `gorm:"primary_key;AUTO_INCREMENT"`
	CosmodromeId int
	Level        int
	HitRate      int
	Health       int
	LoadCapacity int
}

type Fleet struct {
	Id                  int `gorm:"primary_key;AUTO_INCREMENT"`
	CosmodromeId        int
	FleetName           string
	FirstlvlshipAmount  int
	SecondlvlshipAmount int
	ThirdlvlshipAmount  int
	IsAlive             bool
	InCosmodrome        bool
}

type ShipInFleet struct {
	ShipId  int `gorm:"primary_key;not null;unique"`
	FleetId int `gorm:"primary_key;not null;unique"`
}

type Resource struct {
	Id            int    `gorm:"primary_key;AUTO_INCREMENT"`
	PlanetName    string `gorm:"not null;unique"`
	IronAmount    int
	CrystalAmount int
}

type CitycenterUpdate struct {
	Id                  int `gorm:"primary_key;AUTO_INCREMENT"`
	CitycenterId        int
	NextLevel           int
	NextPeoplemaxamount int
	NextCitycenterkpi   int
	StartTime           int
	EndTime             int
}

type IronmineUpdate struct {
	Id                  int `gorm:"primary_key;AUTO_INCREMENT"`
	IronmineId          int
	NextLevel           int
	NextPeoplemaxamount int
	NextMinekpi         int
	StartTime           int
	EndTime             int
}

type CrystalmineUpdate struct {
	Id                  int `gorm:"primary_key;AUTO_INCREMENT"`
	CrystalmineId       int
	NextLevel           int
	NextPeoplemaxamount int
	NextMinekpi         int
	StartTime           int
	EndTime             int
}

type ShipUpdate struct {
	Id               int `gorm:"primary_key;AUTO_INCREMENT"`
	ShipId           int
	NextLevel        int
	NextHitrate      int
	NextLoadcapacity int
	StartTime        int
	EndTime          int
}

type CosmodromeUpdate struct {
	Id                  int `gorm:"primary_key;AUTO_INCREMENT"`
	CosmodromeId        int
	NextLevel           int
	NextCosmodrkpi      int
	NextShipmaxamount   int
	NextFleetsmaxamount int
	StartTime           int
	EndTime             int
}

// You can use script (gamebotdb.backup)for creation of database or use following functionality.
// Before create tables, create database on your local machine manually,
// using following credentials:
// database name - testdb;
// user - root;
// password - root
func main() {
	db, err := Conn()
	if err != nil {
		loger.Log.Errorf("error: ", err)
	}
	db.CreateTable(&Planet{})
	db.CreateTable(&Citycenter{})
	db.CreateTable(&IronMine{})
	db.CreateTable(&CrystalMine{})
	db.CreateTable(&Cosmodrome{})
	db.CreateTable(&Fleet{})
	db.CreateTable(&Ship{})
	db.CreateTable(&ShipInFleet{})
	db.CreateTable(&CitycenterUpdate{})
	db.CreateTable(&IronmineUpdate{})
	db.CreateTable(&CrystalmineUpdate{})
	db.CreateTable(&ShipUpdate{})
	db.CreateTable(&Resource{})
	db.CreateTable(&CosmodromeUpdate{})
}
