# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GORUN=$(GOCMD) run
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    GOINSTALL=$(GOCMD) install
    GOGENERATE=$(GOCMD) generate
    GOLINT=golangci-lint
    GOIMPORTS=goimports
    BINARY_NAME=calendar
    MAIN=./cmd

    all:
		$(GOIMPORTS) -v -w .
		$(GOTEST) -v ./...
		$(GOLINT) run ./... -v
		$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)
    build: 
		GOOS=$(GOOS) $(GOBUILD) -o $(BINARY_NAME) -v
    test:
		$(GOTEST) -v ./...
    clean:
		$(GOCLEAN)
		rm -rf $(BINARY_NAME)
    run:
		$(GORUN) $(MAIN)
    generate:
		$(GOGENERATE) ./...
    imp:
		$(GOIMPORTS) -v -w .
    install:
		$(GOGET) github.com/IamStubborN/otus-golang/2019_07_04_homework_10/copy_file
