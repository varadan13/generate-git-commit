BINARY=generate-git-commit

BIN_DIR=bin

GO=go

all: build

build:
	$(GO) build -o $(BIN_DIR)/$(BINARY) .

install:
	$(GO) install ./cmd/generate-git-commit

run:
	$(GO) run ./cmd/generate-git-commit

clean:
	rm -rf $(BIN_DIR)

build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BIN_DIR)/$(BINARY)-linux ./cmd/generate-git-commit

build-windows:
	GOOS=windows GOARCH=amd64 $(GO) build -o $(BIN_DIR)/$(BINARY).exe ./cmd/generate-git-commit
