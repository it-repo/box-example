module example

go 1.13

require (
	github.com/ddosakura/sola/v2 v2.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.11
	github.com/labstack/echo/v4 v4.1.11
)

replace github.com/ddosakura/sola/v2 => ../../lib/sola
