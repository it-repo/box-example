all: box example

box: echo sola sola.init

echo: ./echo-box
	go build -o ./dist/echo.out ./echo-box
sola: ./sola-box
	go build -o ./dist/sola.out ./sola-box
sola.init: ./sola-box/init
	go build -o ./dist/sola.init.out ./sola-box/init

example: api-doc base-auth config cors custom-http-code-handler favicon-static-backup graphql hello-world logger middleware native new-router proxy restful-api router-auth simple-app websocket
cors: cors.2000 cors.3000 cors.4000 cors.5000

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
restful-api: ./sola-example/restful-api
	go build -o ./dist/sola-example/restful-api.out ./sola-example/restful-api
router-auth: ./sola-example/router-auth
	go build -o ./dist/sola-example/router-auth.out ./sola-example/router-auth
simple-app: ./sola-example/simple-app
	go build -o ./dist/sola-example/simple-app.out ./sola-example/simple-app
websocket: ./sola-example/websocket
	go build -o ./dist/sola-example/websocket.out ./sola-example/websocket

# ---

mysql:
	# TODO: 设置字符集
	docker run --rm -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
adminer:
	docker run --rm -p 8080:8080 adminer:4.7.5
