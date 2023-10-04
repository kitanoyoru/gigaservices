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
    s.serve()
