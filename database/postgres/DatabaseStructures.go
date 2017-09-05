package postgres

type Planet struct {
	Id int `gorm:"primary_key"`
	PlanetName string
	XCoordinate int
	YCoordinate int
	IsActive bool
	IsLoaded bool
}

type Citycenter struct {
	Id int `gorm:"primary_key"`
	PlanetName string
	Level int
	PeopleAmount int
	PeopleMaxAmount int
	GrowingPerDay int
}

type IronMine struct {
	Id int `gorm:"primary_key"`
	PlanetName string
	Level int
	MineKPI int
	PeopleAmount int
	PeopleMaxAmount int
}

type CrystalMine struct {
	Id int `gorm:"primary_key"`
	PlanetName string
	Level int
	MineKPI int
	PeopleAmount int
	PeopleMaxAmount int
}

type Cosmodrome struct {
	Id int `gorm:"primary_key"`
	PlanetName string
	ShipAmount int
	ShipMaxAmount int
	FleetsAmount int
	FleetsMaxamount int
}

type Dok struct {
	Id int `gorm:"primary_key"`
	PlanetName string
	Level int
	DokKPI int
}

type Ship struct {
	Id int `gorm:"primary_key"`
	CosmodromeId int
	Level int
	HitRate int
	Health int
	MaxHealth int
	LoadCapacity int
}

type Fleet struct {
	Id int `gorm:"primary_key"`
	CosmodromeId int
	FleetName string
	FirstlvlshipAmount int
	SecondlvlshipAmount int
	ThirdlvlshipAmount int
	IsAlive bool
}

type ShipInFleet struct {
	ShipId int `gorm:"primary_key"`
	FleetId int
}

type Resource struct {
	Id int `gorm:"primary_key"`
	PlanetName string
	IronAmount int
	CrystalAmount int
}

type CitycenterUpdate struct {
	Id int `gorm:"primary_key"`
	CitycenterId int
	NextLevel int
	NextPeoplemaxamount int
	NextCitycenterkpi int
	StartTime int
	EndTime int
}

type IronmineUpdate struct {
	Id int `gorm:"primary_key"`
	IronmineId int
	NextLevel int
	NextPeoplemaxamount int
	NextMinekpi int
	StartTime int
	EndTime int
}

type CrystalmineUpdate struct {
	Id int `gorm:"primary_key"`
	CrystalmineId int
	NextLevel int
	NextPeoplemaxamount int
	NextMinekpi int
	StartTime int
	EndTime int
}

type ShipUpdate struct {
	Id int `gorm:"primary_key"`
	ShipId int
	NextLevel int
	NextHitrate int
	NextLoadcapacity int
	StartTime int
	EndTime int
}

type DokUpdate struct {
	Id int `gorm:"primary_key"`
	DokId int
	NextLevel int
	NextDokkpi int
	StartTime int
	EndTime int
}



type UserSession struct {

}
