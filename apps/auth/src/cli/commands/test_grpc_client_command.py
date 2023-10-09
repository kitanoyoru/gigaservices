import click
import grpc
from src.proto.health.v1.health_pb2 import HealthCheckRequest  # type: ignore
from src.proto.health.v1.health_pb2_grpc import HealthStub


@click.command()
def test_grpc_client():
    with grpc.insecure_channel("localhost:16002") as channel:
        stub = HealthStub(channel)
        response = stub.Check(HealthCheckRequest(service="database"))
        print(response.status)
