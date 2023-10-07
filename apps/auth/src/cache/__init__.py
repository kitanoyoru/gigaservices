import aioredis
from aioredis import Redis
from src.config import CacheConfig


def create_redis_async_engine_from_conf(config: CacheConfig) -> Redis:
    cache_url = config.get_cache_url()
    return aioredis.from_url(cache_url, encoding="utf-8", decode_resposes=True)
