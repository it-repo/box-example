mysql:
	# TODO: 设置字符集
	docker run --rm -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
adminer:
	docker run --rm -p 8080:8080 adminer:4.7.5
