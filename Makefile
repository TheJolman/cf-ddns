all: build

build:
	go build -o bin/cf-ddns main.go

clean:
	rm -rf bin
