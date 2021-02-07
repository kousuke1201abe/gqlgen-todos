package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	const dataSource = "localuser:localpass@tcp(127.0.0.1:3306)/localdb?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	defer func() {
		if db != nil {
			if err := db.Close(); err != nil {
				panic(err)
			}
		}
	}()
	db.LogMode(true)
	return db
}
