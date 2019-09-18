# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GORUN=$(GOCMD) run
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    GOINSTALL=$(GOCMD) install
    GOLINT=golangci-lint
    BINARY_NAME=microservice

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
    run:
		$(GORUN) .
    install:
		$(GOGET) github.com/IamStubborN/otus-golang/2019_07_04_homework_10/copy_file
