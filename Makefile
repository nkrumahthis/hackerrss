APP_NAME := hackerrss

.PHONY: all
all: build run

.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	go build -o ./hackerrss ./cmd/main.go

.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	./$(APP_NAME)

.PHONY: clean
clean:
	@echo "Cleaning $(APP_NAME)..."
	rm -f $(APP_NAME)