package handler

import (
	"example/sola-example/api-doc/model"
	"net/http"

	"github.com/ddosakura/sola/v2"
)

// List godoc
// @Summary     show list
// @Description Get list
// @Produce     json
// @Success     200 {object} model.Response
// @Router      /list [get]
func List(c sola.Context) error {
	return c.JSON(http.StatusOK, &model.Response{
		Code: 0,
		Msg:  "SUCCESS",
		Data: []string{
			"a",
			"b",
			"c",
		}})
}

// Item godoc
// @Summary     Say Item
// @Description Print Item World!
// @Produce     json
// @Param       id path number true "id of item"
// @Success     200 {object} model.Response
// @Router      /item/{id} [get]
func Item(c sola.Context) error {
	return c.JSON(http.StatusOK, &model.Response{
		Code: 0,
		Msg:  "SUCCESS",
		Data: "a",
	})
}
