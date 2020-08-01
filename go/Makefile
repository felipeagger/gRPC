generate:
	@protoc --proto_path=proto --go_out=plugins=grpc:proto/gen --go_opt=paths=source_relative proto/user.proto

build:
	@echo "---- Building Application ----"
	@go build -o server-bin server/*.go
	@go build -o client-bin client/*.go

run:
	@echo "---- Running Server ----"
	@go run server/*.go

run_client:
	@echo "---- Running Client ----"
	@go run client/*.go