build:
	go build -o bin/server main.go app.go

run: build
	./bin/server

unittests:
	go test -v ./...
