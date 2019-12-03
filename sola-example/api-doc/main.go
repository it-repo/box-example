package main

import (
	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/auth"
	"github.com/ddosakura/sola/v2/middleware/swagger"
	"github.com/ddosakura/sola/v2/middleware/x/router"

	_ "example/sola-example/api-doc/docs"
	"example/sola-example/api-doc/handler"
)

// @title Swagger Example API
// @version 1.0
// @host localhost:3000
// @BasePath /api/v1
// @description This is a sample server celler server.

// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	_sign := auth.Sign(auth.AuthJWT, []byte("sola_key"))
	_auth := auth.Auth(auth.AuthJWT, []byte("sola_key"))

	app := sola.New()
	r := router.New()

	r.BindFunc("GET /swagger", swagger.WrapHandler)

	sub := router.New()
	sub.Prefix = "/api/v1"
	{
		sub.BindFunc("GET /hello", handler.Hello)
		sub.BindFunc("POST /login", auth.NewFunc(_sign, tmp, handler.Hello))
		sub.BindFunc("/logout", auth.CleanFunc(handler.Hello))

		third := router.New()
		third.Prefix = sub.Prefix
		{
			third.BindFunc("GET /list", handler.List)
			third.BindFunc("GET /item/:id", handler.Item)
		}
		sub.Bind("", auth.New(_auth, nil, third.Routes()))
	}
	r.Bind("/api/v1", sub.Routes())

	app.Use(r.Routes())
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}

func tmp(sola.Handler) sola.Handler {
	return handler.Hello
}
