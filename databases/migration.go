package databases

import (
	"WhereIsMyDriver/adapters"
	"WhereIsMyDriver/helper"
	"log"
)

//MigrateDB ...
func MigrateDB(v interface{}) {
	db, err := adapters.ConnectDB()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(v)

	if err != nil {
		log.Println("error when migrate", err)
	}

	errorCloseDB := db.Close()
	helper.CheckError("Error when close db ", errorCloseDB)
}
