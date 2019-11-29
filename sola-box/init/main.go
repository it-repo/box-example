package main

import (
	"fmt"

	"github.com/it-repo/box/service/ac"
	"github.com/it-repo/box/service/route"

	"github.com/ddosakura/sola/v2"
	_ "github.com/go-sql-driver/mysql"
	box "github.com/it-repo/box/middleware/sola-box"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func main() {
	app := sola.New()
	app.LoadConfig()
	dbDriver := viper.GetString("db.driver")
	dbURL := viper.GetString("db.url")
	dbUser := viper.GetString("db.user")
	dbPass := viper.GetString("db.pass")
	solaAuthKey := viper.GetString("sola.auth.key")

	if db, err := gorm.Open(dbDriver, fmt.Sprintf("%s:%s@%s", dbUser, dbPass, dbURL)); err != nil {
		panic(err)
	} else {
		app.CacheORM("default", db)
	}

	box.AC(app.DefaultORM(), solaAuthKey)
	box.Route(app.DefaultORM())

	initDB(app.DefaultORM())
}

func initDB(db *gorm.DB) {
	db.Create(&ac.BoxRole{
		Name: "admin",
		Desc: "Super Administrator. Have access to box all pages.",
		Users: []ac.BoxUser{
			ac.BoxUser{
				Name:   "sakura",
				Nick:   "Super Admin",
				Pass:   "123456",
				Avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
				Desc:   "I am a super administrator",
			},
		}, // 1
		Perms: []ac.BoxPerm{
			ac.BoxPerm{
				Name: "box:admin",
				Desc: "admin page",
			},
			ac.BoxPerm{
				Name: "box:system",
				Desc: "system manage",
			},
		},
	}) // 1
	db.Create(&ac.BoxRole{
		Name: "editor",
		Desc: "Normal Editor. Can see all pages except permission page.",
		Users: []ac.BoxUser{
			ac.BoxUser{
				Name:   "lj",
				Desc:   "I am an editor",
				Nick:   "Normal Editor",
				Pass:   "123456",
				Avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			},
		}, // 2
		Perms: []ac.BoxPerm{
			ac.BoxPerm{
				Name: "box:editor",
				Desc: "editor page",
			},
		},
	}) // 2
	db.Create(&ac.BoxRole{
		Name: "visitor",
		Desc: "Just a visitor. Can only see the home page and the document page.",
		Perms: []ac.BoxPerm{
			ac.BoxPerm{
				Name: "box:visitor",
				Desc: "visitor page",
			},
		},
	}) // 3

	//db.Model(&ac.BoxRole{Model: gorm.Model{ID: 1}}).
	//	Association("Perms").
	//	Append(ac.BoxPerm{Model: gorm.Model{ID: 2}})
	//db.Model(&ac.BoxRole{Model: gorm.Model{ID: 1}}).
	//	Association("Perms").
	//	Append(ac.BoxPerm{Model: gorm.Model{ID: 3}})
	//db.Model(&ac.BoxRole{Model: gorm.Model{ID: 2}}).
	//	Association("Perms").
	//	Append(ac.BoxPerm{Model: gorm.Model{ID: 3}})
	db.Model(&ac.BoxUser{Model: gorm.Model{ID: 1}}).
		Association("Roles").
		Append(&ac.BoxRole{Model: gorm.Model{ID: 2}})
	db.Model(&ac.BoxUser{Model: gorm.Model{ID: 1}}).
		Association("Roles").
		Append(&ac.BoxRole{Model: gorm.Model{ID: 3}})
	db.Model(&ac.BoxUser{Model: gorm.Model{ID: 2}}).
		Association("Roles").
		Append(&ac.BoxRole{Model: gorm.Model{ID: 3}})

	db.Create(&route.BoxRoute{
		FatherID: 0,
		Name:     "box",
		Desc:     "box page",
		Sort:     0,
		Path:     "/box",
		Perm:     "",

		Component: "layout/Layout",
		Redirect:  "noRedirect",
		Title:     "box",
		Icon:      "component",
	}) // 1
	db.Create(&route.BoxRoute{
		FatherID: 1,
		Name:     "admin-page",
		Desc:     "admin page",
		Sort:     0,
		Path:     "admin",
		Perm:     "box:admin",

		Component: "views/box/admin",
		Title:     "admin page",
		NoCache:   true,
	}) // 2
	db.Create(&route.BoxRoute{
		FatherID: 1,
		Name:     "editor-page",
		Desc:     "editor page",
		Sort:     0,
		Path:     "editor",
		Perm:     "box:editor",

		Component: "views/box/editor",
		Title:     "editor page",
		NoCache:   true,
	}) // 3
	db.Create(&route.BoxRoute{
		FatherID: 1,
		Name:     "visitor-page",
		Desc:     "visitor page",
		Sort:     0,
		Path:     "visitor",
		Perm:     "box:visitor",

		Component: "views/box/visitor",
		Title:     "visitor page",
		NoCache:   true,
	}) // 4
	db.Create(&route.BoxRoute{
		FatherID: 1,
		Name:     "free-page",
		Desc:     "free page",
		Sort:     0,
		Path:     "free",
		Perm:     "",

		Component: "views/box/free",
		Title:     "free page",
		NoCache:   true,
	}) // 5

	db.Create(&route.BoxRoute{
		FatherID: 0,
		Name:     "system",
		Desc:     "system page",
		Sort:     0,
		Path:     "/system",
		Perm:     "box:system",

		Component: "layout/Layout",
		Redirect:  "noRedirect",
		Title:     "系统管理",
		Icon:      "component",
	}) // 6
	db.Create(&route.BoxRoute{
		FatherID: 6,
		Name:     "system-info",
		Desc:     "system info",
		Sort:     0,
		Path:     "info",
		Perm:     "",

		Component: "views/system/info",
		Title:     "系统信息",
		Icon:      "component",
		NoCache:   true,
	}) // 7
	db.Create(&route.BoxRoute{
		FatherID: 6,
		Name:     "system-log",
		Desc:     "system log",
		Sort:     0,
		Path:     "log",
		Perm:     "",

		Component: "views/system/log",
		Title:     "系统日志",
		Icon:      "component",
		NoCache:   true,
	}) // 8

	// 404
	db.Create(&route.BoxRoute{
		FatherID: 0,
		Name:     "404",
		Desc:     "404",
		Sort:     0,
		Path:     "*",
		Redirect: "/404",
		Hidden:   true,
	})
}
