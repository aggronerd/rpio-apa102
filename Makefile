PKG_LIST := $(shell go list ./... | grep -v /vendor/ | grep -v /cmd)

all: build

get: ## Get the dependencies
	@go get golang.org/x/lint/golint@v0.0.0-20200302205851-738671d3881b
	@go get -t ./...

lint: get ## Lint the files
	@PATH="${PATH}:${GOPATH}/bin" golint -set_exit_status ${PKG_LIST}

test: get ## Run unittests
	@go test -short ${PKG_LIST}

race: get ## Run data race detector
	@go test -race -short ${PKG_LIST}

build: get ## Build the example binary files
	@mkdir -p target/arm/bin
	@GOOS=linux GOARCH=arm GOARM=6 go build -o target/arm/bin/basic ./examples/basic
	@GOOS=linux GOARCH=arm GOARM=6 go build -o target/arm/bin/spinner ./examples/spinner

clean: ## Remove previous build
	@rm -f target
