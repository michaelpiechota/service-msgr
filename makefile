GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD)	test
GOMOD_DOWNLOAD=$(GOCMD)	mod download
GORUN=$(GOCMD)	run
BINARY_NAME="./bin/serve"
MAIN_LOCATION="./cmd/serve"

.PHONY:	build	test	test-ci	vet	fmt	clean	run	deps	build-linux

all:	test	deps	build

build:	
	go build -o $(BINARY_NAME) $(MAIN_LOCATION)

test:	
	go test	-v -race ./...

vet:	
	go vet ./...

fmt:
	go fmt ./...

clean:	
	go clean
	rm -f $(BINARY_NAME)

run:
	go run $(MAIN_LOCATION)/main.go

deps:
	go mod download

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build	-ldflags="-s -w" -o $(BINARY_NAME) -i $(MAIN_LOCATION)