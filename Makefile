BIN_NAME = gouser

build:
	@echo ""
	@echo "[MAKE:BUILD] ${BIN_NAME}..."
	@echo ""

	mkdir -p bin
	go build -o ./bin/${BIN_NAME} ./cmd/gouserApi/main.go

	@echo ""

run:
	@echo ""
	@echo "[MAKE:RUNNING] ${BIN_NAME}..."
	@echo ""

	go build -o ./bin/${BIN_NAME} ./cmd/gouserApi/main.go
	./bin/${BIN_NAME} $(ARGS)

	@echo ""

test:
	@echo ""
	@echo "[MAKE:Testing] ${BIN_NAME}..."
	@echo ""

	go test -v ./... -count=1

	@echo ""

clean:
	
	@echo ""
	@echo "[MAKE:CLEANUP]..."
	@echo ""

	go clean
	rm -rf ./bin

	@echo ""
