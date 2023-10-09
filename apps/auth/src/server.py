import abc
import logging
import signal
from concurrent.futures import ThreadPoolExecutor
from typing import Any, Coroutine

import grpc
from omegaconf import OmegaConf
from src.config import AppConfig
from src.constants import Constants
from src.proto.health.v1 import health_pb2_grpc
from src.services import HealthService

logger = logging.getLogger(__name__)


class IServer(abc.ABC):
    @abc.abstractmethod
    async def serve(self):
        pass

    @abc.abstractmethod
    def register_custom_signal_handlers(self):
        pass


class Server(IServer):
    cleanup_coroutines: list[Coroutine[Any, Any, None]] = []

    def __init__(self, config: AppConfig):
        self._config = config

    async def serve(self):
        grpc_server = grpc.aio.server(
            ThreadPoolExecutor(max_workers=Constants.MAX_GRPC_WORKERS),
            options=[
                (
                    "grpc.max_send_message_length",
                    Constants.MAX_GRPC_SEND_MESSAGE_LENGTH,
                ),
                (
                    "grpc.max_receive_message_length",
                    Constants.MAX_GRPC_RECEIVE_MESSAGE_LENGTH,
                ),
            ],
        )

        #        auth_service_pb2_grpc.add_AuthServiceServicer_to_server(
        #            AuthService(), grpc_server
        #        )
        health_pb2_grpc.add_HealthServicer_to_server(HealthService(), grpc_server)

        self.register_custom_signal_handlers()

        grpc_server.add_insecure_port("[::]:" + str(self._config.port))

        logger.info(f"gRPC server is running on port {self._config.port}")
        await grpc_server.start()

        async def server_graceful_shutdown():
            logger.info("gRPC server is shutting down")
            await grpc_server.stop(5)

        self.cleanup_coroutines.append(server_graceful_shutdown())

        await grpc_server.wait_for_termination()

    def register_custom_signal_handlers(self):
        def sighup_handler():
            OmegaConf.save(self._config, Constants.CONFIG_FULL_PATH)

        signal.signal(signal.SIGHUP, sighup_handler)  # type: ignore
