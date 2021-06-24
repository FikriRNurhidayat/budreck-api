install:
	go mod tidy

build:
	go build -o bin/budrack main.go

run:
	go run main.go

start: build
	- bin/budrack db:migrate 
	- bin/budrack serve
