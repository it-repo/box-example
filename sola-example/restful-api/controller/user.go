package controller

import (
	"fmt"
	"strconv"

	"github.com/ddosakura/sola/v2/middleware/rest"
	"github.com/ddosakura/sola/v2/middleware/router"
)

// User Model
type User struct {
	ID   int
	Name string
}

// UserC Router
func UserC(root *router.Router) *router.Router {
	r := rest.New(&rest.Option{
		Root: root,
		Path: "/user",
		NewModel: func() interface{} {
			return &User{}
		},
		DefaultPageSize: 10,
		GetFunc: func(ID string) interface{} {
			id, err := strconv.Atoi(ID)
			if err != nil {
				return nil
			}
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

	return r
}
