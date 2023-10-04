import signal

from abc import ABC, abstractmethod
from concurrent.futures import ThreadPoolExecutor

import grpc

from omegaconf import OmegaConf

from src.config import AppConfig
from src.services import AuthService, HealthService
from src.constants import Constants

from src.proto import auth_service_pb2_grpc, health_pb2_grpc


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
        server = grpc.server(
            ThreadPoolExecutor(max_workers=self._config.max_grpc_workers)
        )

        auth_service_pb2_grpc.add_AuthServiceServicer_to_server(AuthService(), server)
        health_pb2_grpc.add_HealthServicer_to_server(HealthService(), server)

        server.add_insecure_port("[::]:" + str(self._config.port))
        server.wait_for_termination()

        self._register_custom_signal_handlers()

    def _register_custom_signal_handlers(self):
        def sighup_handler():
            OmegaConf.save(self._config, Constants.CONFIG_FULL_PATH)

        signal.signal(signal.SIGHUP, sighup_handler)  # type: ignore
