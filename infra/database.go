package infra

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"os/user"
	"path/filepath"
)

var dbConn *gorm.DB

func init() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	folderPath := filepath.Join(u.HomeDir, ".icmd")
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
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
