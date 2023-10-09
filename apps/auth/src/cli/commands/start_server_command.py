import asyncio
import logging

import click
import hydra
from hydra.utils import instantiate
from src.config import AppConfig
from src.constants import Constants
from src.server import Server

logger = logging.getLogger(__name__)


hydra.initialize(config_path="conf")
raw_config = hydra.compose(Constants.CONFIG_NAME)

config: AppConfig = instantiate(raw_config)


@click.command()
def start_server_command():
    print(config.port)
    s = Server(config)
    loop = asyncio.get_event_loop()

    try:
        loop.run_until_complete(s.serve())
    finally:
        loop.run_until_complete(*s.cleanup_coroutines)
        loop.close()
