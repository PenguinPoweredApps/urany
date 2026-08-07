# Variables
PROTO_FILE := ./proto/hello.proto
SERVER_DIR := ./server
CLIENT_DIR := ./client
BIN_DIR := bin

# Tell Make that these aren't actual files
.PHONY: all proto server client clean tidy

# The default target run when you just type 'make'
all: tidy proto server client

# 1. Download/update Go dependencies
tidy:
	@echo "Tidying go modules..."
	go mod tidy

# 2. Generate Go code from the Protocol Buffer file
proto:
	@echo "Generating Protocol Buffer files..."
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	$(PROTO_FILE)

# 3. Build the server binary
server:
	@echo "Building server..."
	mkdir $(BIN_DIR)
	go build -o $(BIN_DIR)/server $(SERVER_DIR)

# 4. Build the client binary
client:
	@echo "Building client..."
	mkdir $(BIN_DIR)
	go build -o $(BIN_DIR)/client $(CLIENT_DIR)

# 5. Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)
	# Optional: Remove generated protobuf files
	# rm -f *.pb.go