COMPILER 		:= go
LINTER	 		:= golangci-lint
VERBOSE_FLAG	:= -v

BIN_FOLDER 		:= ./bin
BIN_FILE_PATH 	:= $(BIN_FOLDER)/data_service

prepare_folder:
	mkdir -p $(BIN_FOLDER)

build: prepare_folder
	$(COMPILER) build -o $(BIN_FOLDER) $(VERBOSE_FLAG) ./...

run: build
	$(BIN_FILE_PATH)

test:
	$(COMPILER) test $(VERBOSE_FLAG) ./...

lint: 
	$(LINTER) run $(VERBOSE_FLAG) ./...

check: 
	build, test, lint