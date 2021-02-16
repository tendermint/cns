PACKAGES=$(shell go list ./...)

###############################################################################
###                           Basic Golang Commands                         ###
###############################################################################

all: install

install: go.sum
	go install -mod=readonly ./cmd/cnsd

install-debug: go.sum
	go build -mod=readonly -gcflags="all=-N -l" ./cmd/cnsd

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)

lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify