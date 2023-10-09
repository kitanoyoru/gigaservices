from __future__ import annotations

from contextlib import asynccontextmanager
from typing import Any, AsyncGenerator, Mapping, Optional

from aioredis import Redis
from src.services.grpc.health_service.check import ServiceCheck


class RedisCache:
    _service_checker: Optional[ServiceCheck] = None
    _service_name = "kita/cache/redis"

    def __init__(self, session: Redis):
        self._session = session

        self._service_checker = ServiceCheck(self._check_service_status)

    @staticmethod
    @asynccontextmanager
    async def from_engine(engine: Redis) -> AsyncGenerator[RedisCache, Any]:
        async with engine.client() as session:
            yield RedisCache(session)

    async def _check_service_status(self):
        response = await self._session.ping()
        if response != "PONG":
            return False

        return True

    def __mapping__(self) -> Mapping[str, Any]:
        return {self._service_name: None}
