from dataclasses import dataclass


@dataclass(slots=True, frozen=True)
class Constants:
    CONFIG_PATH = "/etc"
    CONFIG_NAME = "kita-authservice"
    CONFIG_FULL_PATH = "/etc/kita-authservice.yaml"

    REPOSITORY_SOURCES_LOCATION = "./data/sources.json"

    MAX_GRPC_WORKERS = 12
    MAX_GRPC_SEND_MESSAGE_LENGTH = 256 * 1024 * 1024
    MAX_GRPC_RECEIVE_MESSAGE_LENGTH = 256 * 1024 * 1024

    HTTPX_TIMEOUT = 5
    HTTPX_RETRIES = 5

    DEFAULT_HEALTH_CHECK_TTL = 30
    DEFAULT_HEALTH_CHECK_TIMEOUT = 10
