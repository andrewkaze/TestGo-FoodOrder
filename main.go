package main

import (
	"github.com/getsentry/sentry-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"foodorder/Config"
	"foodorder/Models"
	"foodorder/Routes"
	"time"
)


func main() {
	var err error
	Config.DB, err = gorm.Open(sqlite.Open(Config.DbUrl(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil{
		panic("Cannot connect to DB")
	}

	//defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.AutoMigrate(&Models.Menu{})
	Config.DB.AutoMigrate(&Models.Order{})

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://b4f9feb28e5c4935b29261406b93b127@o485554.ingest.sentry.io/5541081",
	});  err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)
	
	r := Routes.SetupRouter()

	r.Run()
}
