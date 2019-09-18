# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    GOINSTALL=$(GOCMD) install
    GOLINT=golangci-lint
    BINARY_NAME=gocp
    
    all:
		$(GOTEST) -v ./...
		$(GOLINT) run ./... -v
		$(GOBUILD) -o $(BINARY_NAME) -v
    build: 
		GOOS=$(GOOS) $(GOBUILD) -o $(BINARY_NAME) -v
    test: 
		$(GOTEST) -v ./...
    clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
    run:
		$(GOBUILD) -o $(BINARY_NAME) -v
		./$(BINARY_NAME)
    install:
		$(GOGET) github.com/IamStubborN/otus-golang/2019_07_04_homework_10/copy_file
    build_docker:
		docker run --rm -it -v "$(GOPATH)":/go -w /go/src/github.com/IamStubborN/otus-golang/2019_07_04_homework_10/copy_file golang:latest go test /go/src/github.com/IamStubborN/otus-golang/2019_07_04_homework_10/copy_file/gocopy; go build -o "$(BINARY_UNIX)" -v
