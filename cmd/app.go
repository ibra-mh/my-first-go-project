package cmd

import (
	"fmt"
	"my-first-go-project/dto"
	"my-first-go-project/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Application appInitializer = &appIn{}

type appInitializer interface {
	Run() error
	initializeDb()
	// initializeRoutes()
}

type appIn struct {}

func (appIn *appIn) Run() error  {

	return nil
}

func (appIn *appIn) initializeDb() {
	fmt.Print("db initialize")

	DB := Init()
	
    _ = database.New(DB)
	// _ = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Init() *gorm.DB {
    dbURL := "postgres://postgres:123456@localhost:5432/content"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	

    if err != nil {
        fmt.Errorf(err.Error())    }

    db.AutoMigrate(&dto.Path{})

    return db
}