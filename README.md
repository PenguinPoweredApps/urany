protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/hello.proto

How to use this Makefile:

make or make all: Runs tidy, proto, server, and client in sequence.

make proto: Only regenerates the hello.pb.go and hello_grpc.pb.go files if you update hello.proto.

make server: Compiles only the server code and places the executable in the bin/ folder.

make clean: Deletes the compiled binaries and the bin/ directory.
