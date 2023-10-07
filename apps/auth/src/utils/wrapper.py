import asyncio
from contextlib import contextmanager
from types import TracebackType
from typing import Any, ContextManager, Iterator, Optional, Set, Type

from src.utils.deadline import Deadline


class Wrapper(ContextManager[None]):
    _error: Optional[Exception] = None

    cancelled: Optional[bool] = None
    cancel_failed: Optional[bool] = None

    def __init__(self):
        self._tasks: Set[asyncio.Task[Any]] = set()

    def __enter__(self):
        if self._error is not None:
            raise self._error

        task = asyncio.current_task()
        if task is None:
            raise RuntimeError("Called not inside a task")

        self._tasks.add(task)

    def __exit__(
        self,
        exc_type: Optional[Type[BaseException]],
        _exc_value: Optional[BaseException],
        _exc_traceback: Optional[TracebackType],
    ) -> Optional[bool]:
        task = asyncio.current_task()
        assert task

        self._tasks.discard(task)

        if self._error is not None:
            self.cancel_failed = exc_type is not asyncio.CancelledError
            raise self._error

    def cancel(self, error: Exception):
        self._error = error
        for task in self._tasks:
            task.cancel()
        self.cancelled = True


class DeadlineWrapper(Wrapper):
    @contextmanager
    def start(self, deadline: Deadline) -> Iterator[None]:
        timeout = deadline.time_remaining()
        if not timeout:
            raise asyncio.TimeoutError("Deadline exceed")

        loop = asyncio.get_event_loop()
        timer = loop.call_later(
            timeout, lambda: self.cancel(asyncio.TimeoutError("Deadline exceed"))
        )

        try:
            yield
        finally:
            timer.cancel()
