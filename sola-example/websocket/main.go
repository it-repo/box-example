package main

import (
	"fmt"
	"net/http"

	"github.com/ddosakura/sola/v2/middleware/native"
	uuid "github.com/satori/go.uuid"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/ws"
	router "github.com/ddosakura/sola/v2/middleware/xrouter"
)

// Msg -
type Msg struct {
	Code int `json:"code"`
	UUID string
	Msg  string `json:"msg"`
}

func main() {
	var send ws.XSend
	o := &ws.Option{
		Handler: ws.HandleWrap(func() interface{} {
			return &Msg{}
		}, func(UUID uuid.UUID, v interface{}) error {
			fmt.Println(UUID, v)
			return nil
		}),

		First: func(UUID uuid.UUID) {
			fmt.Println(UUID, "login")
			send(ws.ALL, &Msg{UUID: UUID.String()})
		},
		ReceiveError: func(UUID uuid.UUID, e error) {
			fmt.Println("receive error:", UUID, e)
		},
		SendError: func(UUID uuid.UUID, e error) {
			fmt.Println("send error:", UUID, e)
		},
		HandlerError: func(UUID uuid.UUID, e error) {
			fmt.Println("handle error:", UUID, e)
		},
	}
	h, _send := ws.New(o)
	send = ws.SendWrap(_send)

	app := sola.New()
	{
		r := router.New("")
		r.Bind("/ws", h)
		{
			sub := r.Sub("/send")
			sub.Bind("/:msg", func(c sola.Context) error {
				msg := router.Param(c, "msg")
				send(ws.ALL, &Msg{Msg: msg})
				return c.String(http.StatusOK, "all success")
			})
			sub.Bind("/:UUID/:msg", func(c sola.Context) error {
				UUID, _ := uuid.FromString(router.Param(c, "UUID"))
				msg := router.Param(c, "msg")
				if UUID == ws.ALL {
					o.SendError(UUID, ws.ErrNoUUID)
					return c.String(http.StatusOK, "fail")
				}
				send(UUID, &Msg{Msg: msg})
				return c.String(http.StatusOK, "success")
			})
		}
		r.Use(native.Static("static", ""))
		app.Use(r.Routes())
	}
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
