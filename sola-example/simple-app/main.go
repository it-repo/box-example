package main

import (
	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/auth"
	"github.com/ddosakura/sola/v2/middleware/router"
)

func main() {
	defer db.Close()
	app := sola.New()

	{
		r := router.New(&router.Option{
			UseNotFound: true,
		})
		r.Bind("POST /register", register)
		r.Bind("/logout", auth.CleanFunc(success))
		{
			sub := r.Sub(nil)
			sub.Use(auth.LoadAuthCache)
			h := sola.MergeFunc(success, login, jwtSign)
			sub.Bind("POST /login", h)
		}
		{
			sub := r.Sub(nil)
			sub.Use(jwtAuth)
			sub.Bind("GET /welcome", welcome)
			sub.Bind("GET /user/:id", user)
		}
		app.Use(r.Routes())
	}

	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
