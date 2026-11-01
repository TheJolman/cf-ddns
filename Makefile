all: build

fmt:
	go fmt ./...

build: fmt
	go build -o bin/cf-ddns main.go

clean:
	rm -rf bin
