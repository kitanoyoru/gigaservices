from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession

from src.crawler import LogModelImport
from src.models.log import LogModel


class Database:
    def __init__(self, session: AsyncSession):
        self._session = session

    async def save_log(self, log_import: LogModelImport):
        log_model = await self._session.execute(
            select(LogModel).where(LogModel.id == log_import.id)
        )

        log_model = log_model.scalar_one_or_none()

        if log_model is None:
            log_model = LogModel.from_import(log_import)
        else:
            log_model.merge_with_import(log_import)

        self._session.add(log_model)
        await self._session.flush()
