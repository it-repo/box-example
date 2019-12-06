all: dist/init dist/sola-box
	cp ./sola.yml ./dist/sola.yml
clean:
	rm -r dist
	rm -r statik

dist/init: init/*.go
	go build -o dist/init ./init
dist/sola-box: *.go statik/statik.go
	go build -o dist/sola-box .

statik/statik.go: public/index.html
	go generate
