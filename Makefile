all: build

build:
	@echo "Building..."
	
	@go build -o main cmd/api/publish/main.go
	@go build -o main cmd/api/consumer/main.go

run-publish:
	@go run cmd/api/publish/main.go

run-consumer:
	@go run cmd/api/consumer/main.go

clean:
	@echo "Cleaning..."
	@rm -f main

