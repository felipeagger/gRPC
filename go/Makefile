generate:
	@protoc --go_out=plugins=grpc:user user.proto

build:
	@echo "---- Building Application ----"
	@go build -o binary main.go

run:
	@echo "---- Running Application ----"
	@go run main.go