build:
	go build -o main ./src/cmd/server/main.go

run:
	go run ./src/cmd/server/main.go

build-and-run:
	go build -o main ./src/cmd/server/main.go
	./main
	
lint:
	golint ./...