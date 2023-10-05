import signal
import logging

from abc import ABC, abstractmethod
from concurrent.futures import ThreadPoolExecutor

import grpc

from omegaconf import OmegaConf

from grpclib.server import Server as GRPCServer

from src.config import AppConfig
from src.services import AuthService, HealthService
from src.constants import Constants

from src.proto import auth_service_pb2_grpc, health_pb2_grpc


logger = logging.getLogger(__name__)


class IServer(ABC):
    @abstractmethod
    def serve(self):
        pass

    @abstractmethod
    def _register_custom_signal_handlers(self):
        pass


class Server(IServer):
    def __init__(self, config: AppConfig):
        self._config = config

    def serve(self):
        grpc_server = grpc.server(
            ThreadPoolExecutor(max_workers=Constants.MAX_GRPC_WORKERS),
            options=[
                ("grpc.max_send_message_length", Constants.MAX_GRPC_SEND_MESSAGE_LENGTH),
                ("grpc.max_receive_message_length", Constants.MAX_GRPC_RECEIVE_MESSAGE_LENGTH),
            ],
        )

        auth_service_pb2_grpc.add_AuthServiceServicer_to_server(AuthService(), grpc_server)
        health_pb2_grpc.add_HealthServicer_to_server(HealthService(), grpc_server)

        self._register_custom_signal_handlers()

        grpc_server.add_insecure_port("[::]:" + str(self._config.port))

	    logger.info(f"gRPC is running on port {self._config.port}")
        grpc_server.wait_for_termination()

    def _register_custom_signal_handlers(self):
        def sighup_handler():
            OmegaConf.save(self._config, Constants.CONFIG_FULL_PATH)

        signal.signal(signal.SIGHUP, sighup_handler)  # type: ignore
