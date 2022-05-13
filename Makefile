COMPILER 		:= go
LINTER	 		:= golangci-lint
VERBOSE_FLAG	:= -v

BIN_FOLDER 		:= ./bin
BIN_FILE_PATH 	:= $(BIN_FOLDER)/data_service


build:
	$(COMPILER) build -o $(BIN_FOLDER) $(VERBOSE_FLAG) ./...

run: build
	$(BIN_FILE_PATH)

test:
	$(COMPILER) test $(VERBOSE_FLAG) ./...

lint: 
	$(LINTER) run $(VERBOSE_FLAG) ./...
