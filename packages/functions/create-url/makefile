.PHONY: build clean

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bootstrap main.go
	zip bootstrap.zip bootstrap

clean:
	rm -f bootstrap bootstrap.zip