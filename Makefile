TARGET=auth



all: clean build


clean:

	rm -rf $(TARGET)



build:

	go build -o $(TARGET) main.go



proto:

	protoc --go_out=plugins=grpc:. proto/entities/user.proto