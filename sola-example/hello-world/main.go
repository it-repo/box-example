package main

import (
	"net/http"

	"github.com/ddosakura/sola/v2"
)

// handler
func hw(c sola.Context) error {
	// 输出 Hello World
	return c.String(http.StatusOK, "Hello World")
}

func main() {
	// 将 handler 包装为中间件
	m := sola.Handler(hw).M()
	// 使用中间件
	sola.Use(m)
	// 监听 127.0.0.1:3000
	sola.ListenKeep("127.0.0.1:3000")
}
