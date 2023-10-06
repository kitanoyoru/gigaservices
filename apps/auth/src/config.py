from dataclasses import dataclass
from enum import Enum
from typing import List

import sqlalchemy
from hydra.core.config_store import ConfigStore


class DatabaseDriver(Enum):
    POSTGRES = "postgresql"


class CacheDriver(Enum):
    REDIS = "redis"


@dataclass
class DatabaseConfig:
    port: int
    host: str
    driver: DatabaseDriver
    username: str
    password: str
    db_name: str

    def get_db_url(self) -> sqlalchemy.engine.url.URL:
        return sqlalchemy.engine.url.URL.create(
            drivername=self.driver.value,
            username=self.username,
            password=self.password,
            host=self.host,
            port=self.port,
            database=self.db_name,
        )


@dataclass
class CacheConfig:
    port: int
    host: str
    driver: CacheDriver
    username: str
    password: str

    def get_cache_url(self):
        return f"{self.driver.value}://{self.host}:{self.port}"


@dataclass
class MessageBrokerConfig:
    brokers_url: List[str]


@dataclass
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
