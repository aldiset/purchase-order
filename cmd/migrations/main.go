package main

import (
	"log"

	"purchase-order/config"
	"purchase-order/config/db"
	"purchase-order/modules/database/user"
	"purchase-order/utils/logger"
)

func main() {
	cfg := config.LoadConfingAPI()

	mlog, err := logger.New("info")
	if err != nil {
		log.Fatalln(err.Error())
	}

	db, err := db.NewMySQL(&cfg.Database)
	if err != nil {
		log.Fatalln("open mysql failed:", err.Error())
	}

	mlog.Info("start migration model ...")

	// Add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user.User{})

	mlog.Info("success migration model ...")
}
