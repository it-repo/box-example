package main

import (
	"github.com/ddosakura/sola/v2/extension/hot"

	"github.com/ddosakura/sola/v2"
)

// export(s)
var (
	ExportHandler = map[string]sola.Handler{
		"hw": func(c sola.Context) error {
			return hot.Modules(c).Handler("hw2")(c)
		},
	}
)

func main() {}
