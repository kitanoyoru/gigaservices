from dataclasses import dataclass

from hydra.core.config_store import ConfigStore


@dataclass(slots=True, frozen=True)
class DatabaseConfig:
    url: str
    username: str
    password: str


@dataclass(slots=True, frozen=True)
class CacheConfig:
    url: str
    username: str
    password: str


@dataclass(slots=True, frozen=True)
class MessageBrokerConfig:
    brokers_url: list[str]


@dataclass(slots=True, frozen=True)
class AppConfig:
	port: int
	max_grpc_workers: int

    db: DatabaseConfig
    cache: CacheConfig
    broker: MessageBrokerConfig


cs = ConfigStore.instance()

cs.store(name="app_config", node=AppConfig)

cs.store(name="database_config", node=DatabaseConfig)
cs.store(name="cache_config", node=CacheConfig)
cs.store(name="message_broker_config", node=MessageBrokerConfig)
