package main

import (
	"fmt"
	"strconv"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/rest"
)

// User Model
type User struct {
	ID   int
	Name string
}

func main() {
	app := sola.New()
	r := rest.New(&rest.Option{
		Path: "/user",
		NewModel: func() interface{} {
			return &User{}
		},
		DefaultPageSize: 10,
		GetFunc: func(id int) interface{} {
			return &User{id, "No." + strconv.Itoa(id)}
		},
		ListFunc: func(page int, size int) interface{} {
			users := make([]User, 0, size)
			for id := (page-1)*size + 1; id <= page*size; id++ {
				users = append(users, User{id, "No." + strconv.Itoa(id)})
			}
			return users
		},
		PostFunc: func(v interface{}) error {
			u := v.(*User)
			fmt.Println("Insert", u.Name)
			return nil
		},
		PutFunc: func(v interface{}) error {
			u := v.(*User)
			fmt.Println("Update", u.ID, "=>", u.Name)
			return nil
		},
		DeleteFunc: func(id string) error {
			fmt.Println("Delete", id)
			return nil
		},
	})

	app.Use(r.Routes())
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
