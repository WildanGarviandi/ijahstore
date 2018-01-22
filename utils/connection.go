package utils

import (
	"github.com/IjahStore/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// ConnectDB is a function to export DB Setting
func ConnectDB() (*gorm.DB, error) {
	var err error

	db, err = gorm.Open("sqlite3", "./ijahstore.db")

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	var itemsModels = models.Items{}
	var logItemsInboundModels = models.LogItemsInbound{}
	var logItemsOutboundModels = models.LogItemsOutbond{}

	db := db.AutoMigrate(&itemsModels, &logItemsInboundModels, &logItemsOutboundModels)

	return db, err
}

// CloseDB is to defers the Close BD
func CloseDB() {
	defer db.Close()
}
