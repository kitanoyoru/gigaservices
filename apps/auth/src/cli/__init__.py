import click

from .commands.get_config_command import get_config_command
from .commands.start_server_command import start_server_command


@click.group()
def cli():
    pass


cli.add_command(start_server_command)
cli.add_command(get_config_command)
