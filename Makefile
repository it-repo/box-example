all: box example

box: echo sola sola.init

echo: ./echo-box
	go build -o ./dist/echo.out ./echo-box
sola: ./sola-box
	go build -o ./dist/sola.out ./sola-box
sola.init: ./sola-box/init
	go build -o ./dist/sola.init.out ./sola-box/init

example: api-doc base-auth config cors custom-http-code-handler favicon-static-backup graphql hello-world hot-plugin hot-update logger middleware native new-router proxy proxy-balance restful-api router-auth simple-app version-test websocket wx-miniapp
cors: cors.2000 cors.3000 cors.4000 cors.5000
version-test: test2.1.2

api-doc: ./sola-example/api-doc
	echo "api-doc need swag"
	go build -o ./dist/sola-example/api-doc.out ./sola-example/api-doc
base-auth: ./sola-example/base-auth
	go build -o ./dist/sola-example/base-auth.out ./sola-example/base-auth
config: ./sola-example/config
	go build -o ./dist/sola-example/config.out ./sola-example/config
cors.2000: ./sola-example/cors/2000
	go build -o ./dist/sola-example/cors.2000.out ./sola-example/cors/2000
cors.3000: ./sola-example/cors/3000
	go build -o ./dist/sola-example/cors.3000.out ./sola-example/cors/3000
cors.4000: ./sola-example/cors/4000
	go build -o ./dist/sola-example/cors.4000.out ./sola-example/cors/4000
cors.5000: ./sola-example/cors/5000
	go build -o ./dist/sola-example/cors.5000.out ./sola-example/cors/5000
custom-http-code-handler: ./sola-example/custom-http-code-handler
	go build -o ./dist/sola-example/custom-http-code-handler.out ./sola-example/custom-http-code-handler
favicon-static-backup: ./sola-example/favicon-static-backup
	go build -o ./dist/sola-example/favicon-static-backup.out ./sola-example/favicon-static-backup
graphql: ./sola-example/graphql
	go build -o ./dist/sola-example/graphql.out ./sola-example/graphql
hello-world: ./sola-example/hello-world
	go build -o ./dist/sola-example/hello-world.out ./sola-example/hello-world
hot-plugin: ./sola-example/hot-plugin
	go build -o ./dist/sola-example/hot-plugin.out ./sola-example/hot-plugin
hot-update: ./sola-example/hot-update
	go build -o ./dist/sola-example/hot-update.out ./sola-example/hot-update
logger: ./sola-example/logger
	go build -o ./dist/sola-example/logger.out ./sola-example/logger
middleware: ./sola-example/middleware
	go build -o ./dist/sola-example/middleware.out ./sola-example/middleware
native: ./sola-example/native
	echo "native need statik"
	go build -o ./dist/sola-example/native.out ./sola-example/native
new-router: ./sola-example/new-router
	go build -o ./dist/sola-example/new-router.out ./sola-example/new-router
proxy: ./sola-example/proxy
	go build -o ./dist/sola-example/proxy.out ./sola-example/proxy
proxy-balance: ./sola-example/proxy-balance
	go build -o ./dist/sola-example/proxy-balance.out ./sola-example/proxy-balance
restful-api: ./sola-example/restful-api
	go build -o ./dist/sola-example/restful-api.out ./sola-example/restful-api
router-auth: ./sola-example/router-auth
	go build -o ./dist/sola-example/router-auth.out ./sola-example/router-auth
simple-app: ./sola-example/simple-app
	go build -o ./dist/sola-example/simple-app.out ./sola-example/simple-app
test2.1.2: ./sola-example/test2.1.2
	go build -o ./dist/sola-example/test2.1.2.out ./sola-example/test2.1.2
websocket: ./sola-example/websocket
	go build -o ./dist/sola-example/websocket.out ./sola-example/websocket
wx-miniapp: ./sola-example/wx-miniapp
	go build -o ./dist/sola-example/wx-miniapp.out ./sola-example/wx-miniapp

# ---

mysql:
	# TODO: 设置字符集
	docker run --rm -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
adminer:
	docker run --rm -p 8080:8080 adminer:4.7.5
