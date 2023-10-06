from __future__ import annotations

from datetime import date

from sqlalchemy import Boolean, DateTime, Integer, String, func
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column


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
