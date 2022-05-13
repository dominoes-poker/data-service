COMPILER 		:= go
BIN_FOLDER 		:= ./bin
BIN_FILE_PATH 	:= $(BIN_FOLDER)/data_service

prepare:
	mkdir -p $(BIN_FOLDER)

build: prepare
	$(COMPILER) build -o $(BIN_FOLDER) ./...

run: build
	$(BIN_FILE_PATH)

test:
	$(COMPILER) test ./...
