CURRENT_DIR = $(shell pwd)
LOCAL_BIN=$(CURRENT_DIR)/bin

run:
	mkdir -p bin
	./bin/main

build:
	CGO_ENABLED=0 go build -o bin/main cmd/main.go 

bin-deps:
	mkdir -p bin
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.5.3

db\:up:
	$(LOCAL_BIN)/goose -dir migrations postgres "postgresql://postgres:1234@localhost:5433/communication?sslmode=disable&timezone=UTC" up

db\:down:
	$(LOCAL_BIN)/goose -dir migrations postgres "postgresql://postgres:1234@localhost:5433/communication?sslmode=disable&timezone=UTC" down

db\:create:
	$(LOCAL_BIN)/goose -dir migrations create "$(NAME)" sql