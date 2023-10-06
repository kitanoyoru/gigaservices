from __future__ import annotations

from contextlib import asynccontextmanager
from typing import Any, AsyncGenerator

from sqlalchemy.ext.asyncio import AsyncEngine, AsyncSession

from src.database.database import Database


class Service:
    def __init__(self, database: Database):
        self._database = database

    @staticmethod
    @asynccontextmanager
    async def from_engine(engine: AsyncEngine) -> AsyncGenerator[Service, Any]:
        async with AsyncSession(engine) as session:
            async with session.begin():
                yield Service.from_session(session)

    @staticmethod
    def from_session(session: AsyncSession) -> Service:
        database = Database(session)
        return Service(database)
