from __future__ import annotations

from contextlib import asynccontextmanager
from typing import AsyncGenerator

from aioredis import Redis
from aioredis.client import Any


class RedisCache:
    def __init__(self, session: Redis):
        self._session = session

    @staticmethod
    @asynccontextmanager
    async def from_engine(engine: Redis) -> AsyncGenerator[RedisCache, Any]:
        async with engine.client() as session:
            yield RedisCache(session)
