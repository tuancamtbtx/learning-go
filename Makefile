TARGET=auth

include	./hack/Makefile


default: di build;

all: clean build
clean:
	rm -rf $(TARGET)

di:

di: $(WIRE)
	$(WIRE) gen -output_file_prefix build_server_  ./cmd/server

build:

	go build -o $(TARGET) main.go



proto:

	protoc --go_out=plugins=grpc:. proto/entities/user.proto