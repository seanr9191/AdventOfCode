GOCMD=go
GOCUILD=$(GOCMD) build
GORUN=$(GOCMD) run
OUT_PATH=./out/
NAME=AdventOfCode
ENTRY_PATH=./cmd/$(NAME)/main.go

sync:
	$(GOCMD) mod tidy

run: build
	$(OUT_PATH)$(NAME)

build:
	@mkdir -p $(OUT_PATH)
	$(GOCMD) build -o $(OUT_PATH)$(NAME) $(ENTRY_PATH)