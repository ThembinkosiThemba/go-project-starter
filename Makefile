# Define the name of your Go binary
BINARY_NAME=github.com/ThembinkosiThemba/go-project-starter
SOURCES=$(wildcard *.go)

# Define the default target when you just run `make` with no arguments
all: generate-swagger run

# Build the Go binary
build:
	go build -o $(BINARY_NAME)

# Tidy the Go application
tidy:
	go mod tidy
# Run the Go application
run:
	go run cmd/main.go
	
swagger:
	cd cmd && swag init

# Run tests
test:
	go test -list .

# Clean the project by removing the binary
clean:
	rm -f $(BINARY_NAME)

# Define phony targets (these don't represent files)
.PHONY: all build run test clean
