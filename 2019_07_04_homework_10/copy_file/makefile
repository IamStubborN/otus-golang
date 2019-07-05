# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    GOINSTALL=$(GOCMD) install
    GOLINT=golangci-lint
    BINARY_NAME=gocp
    BINARY_UNIX=$(BINARY_NAME)_unix
    BINARY_WINDOWS=$(BINARY_NAME)_windows.exe
    BINARY_DARWIN=$(BINARY_NAME)_darwin
    
    all:
	$(GOTEST) -v ./...
	$(GOLINT) run ./... -v 
	$(GOBUILD) -o $(BINARY_NAME) -v
    build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
    test: 
	$(GOTEST) -v ./...
    clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_WINDOWS)
	rm -f $(BINARY_DARWIN)
	rm -f $(BINARY_UNIX)
    run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
    install:
	$(GOGET) github.com/IamStubborN/otus-golang/2019_07_04_homework_10/copy_file

    # Cross compilation
    build_windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WINDOWS) -v
    build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
    build_darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_DARWIN) -v
    build_docker:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/github.com/IamStubborN/otus-golang/2019_07_04_homework_10/copy_file golang:latest go test /go/src/github.com/IamStubborN/otus-golang/2019_07_04_homework_10/copy_file/gocopy; go build -o "$(BINARY_UNIX)" -v
