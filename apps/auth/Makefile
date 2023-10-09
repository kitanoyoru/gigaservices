__default__:
	@echo "Please specify a target to make"

GRPC_GEN=python3 -m grpc_tools.protoc -I. --python_out=./src --grpc_python_out=./src
GRPC_GENERATED=*{_pb2.py,_grpc.py,.pyi}

.PHONY: proto
proto:
	$(GEN) proto/health/v1/health.proto
	$(GEN) proto/auth/v1/auth.proto

.PHONY: update
update:
	@poetry update
