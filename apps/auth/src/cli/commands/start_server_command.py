import asyncio
import logging

import hydra

from src.config import AppConfig
from src.constants import Constants
from src.server import Server

logger = logging.getLogger(__name__)


@hydra.main(
    version_base=None,
    config_path=Constants.CONFIG_PATH,
    config_name=Constants.CONFIG_NAME,
)
def start_server_command(config: AppConfig):
    s = Server(config)
    loop = asyncio.get_event_loop()

    try:
        loop.run_until_complete(s.serve())
    finally:
        loop.run_until_complete(*s.cleanup_coroutines)
        loop.close()
