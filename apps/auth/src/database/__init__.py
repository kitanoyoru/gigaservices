import alembic.config
from sqlalchemy import Engine, create_engine
from sqlalchemy.ext.asyncio import AsyncEngine, create_async_engine
from src.config import DatabaseConfig


def create_engine_from_conf(config: DatabaseConfig) -> Engine:
    db_url = config.get_db_url()
    return create_engine(db_url)


def create_async_engine_from_conf(config: DatabaseConfig) -> AsyncEngine:
    db_url = config.get_db_url()
    return create_async_engine(db_url)


def create_alembic_config(engine: Engine) -> alembic.config.Config:
    config = alembic.config.Config()

    config.set_main_option("script_location", "sammelrepository:migrations")
    config.attributes["engine"] = engine

    return config
