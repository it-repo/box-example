package main

import (
	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Item Model
type Item struct {
	*gorm.Model
	Data string `gorm:"not null"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "test.db")
	db.AutoMigrate(&Item{})
	app := sola.New()
	app.LoadConfig()
	app.CacheORM("default", db)
	r := router.New()

	r.BindFunc("/insert/:data", insert)
	r.BindFunc("/list", list)
	r.BindFunc("/err", err)

	app.Use(r.Routes())
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
