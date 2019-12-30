module example

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/ddosakura/sola/v2 v2.1.2
	github.com/go-sql-driver/mysql v1.4.1
	github.com/it-repo/box v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.11
	github.com/labstack/echo/v4 v4.1.11
	github.com/rakyll/statik v0.1.6
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/afero v1.2.2
	github.com/spf13/viper v1.5.0
	github.com/swaggo/swag v1.6.3
	github.com/yuin/gopher-lua v0.0.0-20190514113301-1cd887cd7036
)

replace github.com/it-repo/box => ../box
