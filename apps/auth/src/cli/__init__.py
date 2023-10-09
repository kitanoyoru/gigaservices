import click

# from .commands.get_config_command import get_config_command
from .commands.start_server_command import start_server_command
from .commands.test_grpc_client_command import test_grpc_client


@click.group()
def cli():
    pass


cli.add_command(start_server_command)
# cli.add_command(get_config_command)
cli.add_command(test_grpc_client)
