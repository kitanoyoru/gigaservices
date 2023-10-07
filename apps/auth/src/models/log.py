from __future__ import annotations

from datetime import date

from sqlalchemy import Boolean, DateTime, Integer, String, func
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column
from src.crawler import LogModelImport


class Base(DeclarativeBase):
    pass


class LogModel(Base):
    __tablename__ = "log"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)

    is_error: Mapped[bool] = mapped_column(Boolean, default=False)
    message: Mapped[str] = mapped_column(String, nullable=False)

    created_at: Mapped[date] = mapped_column(
        DateTime, default=func.now(), nullable=False
    )

    @staticmethod
    def from_import(log_import: LogModelImport) -> LogModel:
        return LogModel(
            id=log_import.id,
            is_error=log_import.is_error,
            message=log_import.message,
            created_at=log_import.created_at,
        )

    def merge_with_import(self, log_import: LogModelImport):
        self.id = log_import.id
        self.is_error = log_import.is_error
        self.message = log_import.message
        self.created_at = log_import.created_at
