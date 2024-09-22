package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var dbMigrate *gorm.DB

func SetDB(connection string) {
	var err error
	db, err = gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println("Veritabanına bağlanılamadı:", err)
		panic(err)
	}
	db.SingularTable(true)
	fmt.Println("Veritabanına başarıyla bağlanıldı.")
}

func SetDBMigrate(connection string) {
	var err error
	dbMigrate, err = gorm.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	dbMigrate.SingularTable(true)
}

func GetDB() *gorm.DB {
	if db == nil {
		panic("Veritabanı bağlantısı henüz başlatılmamış")
	}
	return db
}

func GetDBMigrate() *gorm.DB {
	return dbMigrate
}
