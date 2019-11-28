package handler

import (
	"net/http"

	"github.com/ddosakura/sola/v2"
)

// Hello godoc
// @Summary     Say Hello
// @Description Print Hello World!
// @Produce     plain
// @Success     200 {string} string "Hello World!"
// @Router      /hello [get]
func Hello(c sola.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
