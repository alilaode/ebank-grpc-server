ifeq ($(OS), Windows_NT)
	BIN_FILENAME  := ebank-grpc-server.exe
else
	BIN_FILENAME  := ebank-grpc-server
endif

.PHONY: tidy
tidy:
	go mod tidy


.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "bin" rd /s /q bin	
else
	rm -fR ./bin
endif


.PHONY: build
build: clean
	go build -o ./bin/${BIN_FILENAME} ./cmd


.PHONY: execute
execute: clean build
	./bin/${BIN_FILENAME}


## postgresql
.PHONY: postgres
postgres:
	docker run --name ebank_grpc -p 5453:5453 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

.PHONY: createdb
createdb:
	docker exec -it ebank_grpc createdb --username=root --owner=root grpc

.PHONY: dropdb
dropdb:
	docker exec -it ebank_grpc dropdb grpc