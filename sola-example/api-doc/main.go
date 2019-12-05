package main

import (
	"net/http"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/auth"
	"github.com/ddosakura/sola/v2/middleware/router"
	"github.com/ddosakura/sola/v2/middleware/swagger"

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
	jwtSign, jwtAuth := auth.NewJWT([]byte("sola_key"))

	app := sola.New()
	r := router.New(&router.Option{
		UseNotFound: true,
	})
	r.Bind("GET /swagger/*", swagger.WrapHandler)

	{
		sub := r.Sub(&router.Option{
			Pattern: "/api/v1",
		})
		sub.Bind("GET /hello", handler.Hello)
		sub.Bind("/logout", auth.CleanFunc(handler.Hello))
		{
			third := sub.Sub(nil)
			third.Use(auth.LoadAuthCache)
			h := sola.MergeFunc(handler.Hello, login, jwtSign)
			third.Bind("POST /login", h)
		}

		{
			third := sub.Sub(nil)
			third.Use(jwtAuth)
			third.Bind("GET /list", handler.List)
			third.Bind("GET /item/:id", handler.Item)
		}
	}

	app.Use(r.Routes())
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}

func tmp(sola.Handler) sola.Handler {
	return handler.Hello
}

func login(next sola.Handler) sola.Handler {
	return func(c sola.Context) error {
		// 获取 GET 参数
		r := c.Request()
		user := r.PostFormValue("user")
		pass := r.PostFormValue("pass")

		// 校验
		if len(user) == 0 || pass != "123456" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"code": -1,
				"msg":  "FAIL",
			})
		}

		// 储存用户名等信息
		c.Set(auth.CtxClaims, map[string]interface{}{
			"issuer": "sola",
			"user":   user,
		})
		return next(c) // 登录成功
	}
}
