package database

import (
	"github.com/itzmeerkat/icmd/infra/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

var dbConn *gorm.DB

func init() {
	folderPath := config.GetAppFolder()
	db, err := gorm.Open(sqlite.Open(filepath.Join(folderPath, "sqlite.db")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbConn = db
}

func MigrateType(v interface{}) {
	err := dbConn.AutoMigrate(v)
	if err != nil {
		panic(err)
	}
}

func CreateRecord(v interface{}) {
	dbConn.Create(v)
}

func GetRecords(dest interface{}, cond interface{}) {
	dbConn.Find(dest, cond)
}
