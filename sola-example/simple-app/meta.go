package main

import (
	"github.com/ddosakura/sola/v2/middleware/auth"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDB(dialect string, args ...interface{}) *gorm.DB {
	db, err := gorm.Open(dialect, args...)
	if err != nil {
		panic("Failed to connect to database!")
	}
	return db
}

var (
	db = initDB("sqlite3", "test.db")

	jwtSign, jwtAuth = auth.NewJWT([]byte("sola_key"))
)
