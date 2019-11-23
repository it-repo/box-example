package main

import (
	"errors"
	"net/http"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/router"
)

func insert(c sola.Context) error {
	db := c.Sola().DefaultORM()
	data := router.Param(c, "data")
	if err := db.Create(&Item{
		Data: data,
	}).Error; err != nil {
		return err
	}
	return c.String(http.StatusOK, "SUCCESS")
}

func list(c sola.Context) error {
	db := c.Sola().DefaultORM()
	var list []Item
	if err := db.Find(&list).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, list)
}

func err(c sola.Context) error {
	return errors.New("a error")
}
