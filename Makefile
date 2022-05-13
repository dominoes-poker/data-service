COMPILER 		:= go
BIN_FOLDER 		:= ./bin
BIN_FILE_PATH 	:= $(BIN_FOLDER)/data_service


build:
	$(COMPILER) build -o $(BIN_FOLDER) ./...

run: build
	$(BIN_FILE_PATH)

test:
	$(COMPILER) test ./...
