generate:
	@python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. user.proto

run:
	@echo "---- Running Server ----"
	@python app.py

run_client:
	@echo "---- Running Client ----"
	@python client.py