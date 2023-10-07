import abc
import asyncio
import logging
import time
from typing import Awaitable, Callable, Optional, Set

from src.constants import Constants
from src.utils.deadline import Deadline
from src.utils.wrapper import DeadlineWrapper

log = logging.getLogger(__name__)

DEFAULT_CHECK_TTL = 30
DEFAULT_CHECK_TIMEOUT = 10

_Status = Optional[bool]

logger = logging.getLogger(__name__)


class CheckBase(abc.ABC):
    @abc.abstractmethod
    def __status__(self) -> _Status:
        pass

    @abc.abstractmethod
    async def __check__(self) -> _Status:
        pass

    @abc.abstractmethod
    async def __subscribe__(self) -> asyncio.Event:
        pass

    @abc.abstractmethod
    async def __unsubscribe__(self, event: asyncio.Event) -> None:
        pass


class ServiceCheck(CheckBase):
    _value: Optional[_Status] = None
    _poll_task = None
    _last_check: Optional[float] = None

    def __init__(
        self,
        check_func: Callable[[], Awaitable[_Status]],
        *,
        check_ttl: float = Constants.DEFAULT_HEALTH_CHECK_TTL,
        check_timeout: float = Constants.DEFAULT_HEALTH_CHECK_TIMEOUT,
    ):
        self._check_func = check_func
        self._check_ttl = check_ttl
        self._check_timeout = check_timeout

        self._events: Set[asyncio.Event] = set()

        self._check_lock = asyncio.Event()
        self._check_lock.set()

        self._check_wrapper = DeadlineWrapper()

    def __status__(self) -> _Status:
        return self._value

    async def __check__(self) -> _Status:
        if (
            self._last_check is not None
            and time.monotonic() - self._last_check < self._check_ttl
        ):
            return self._value

        if not self._check_lock.is_set():
            await self._check_lock.wait()
            return self._value

        prev_value = self._value
        self._check_lock.clear()

        try:
            deadline = Deadline.from_timeout(self._check_timeout)
            with self._check_wrapper.start(deadline):
                await self._check_func()
        except asyncio.CancelledError:
            raise
        except Exception:
            logger.exception("Health check failed")
            self._value = False
        finally:
            self._check_lock.set()

        self._last_check = time.monotonic()

        if self._value != prev_value:
            log = logger.info if self._value else logger.debug
            log(
                f"Health check {self._check_func.__name__} statuc changed to {self._value}"
            )
            for event in self._events:
                event.set()

        return self._value

    async def _poll(self):
        while True:
            status = await self.__check__()
            if status:
                await asyncio.sleep(self._check_ttl)
            else:
                await asyncio.sleep(self._check_ttl)  # refactor this

    async def __subscribe__(self) -> asyncio.Event:
        if self._poll_task is None:
            loop = asyncio.get_event_loop()
            self._poll_task = loop.create_task(self._poll())

        event = asyncio.Event()
        self._events.add(event)

        return event

    async def __unsubscribe__(self, event: asyncio.Event) -> None:
        self._events.discard(event)

        if not self._events:
            assert self._poll_task is not None
            task = self._poll_task
            self._poll_task = None
            task.cancel()

            try:
                await task
            except asyncio.CancelledError:
                pass


class ServiceStatus(CheckBase):
    def __init__(self):
        self._value: _Status = None
        self._events: Set[asyncio.Event] = set()

    def set(self, value: _Status):
        prev_value = self._value
        self._value = value

        if self._value != prev_value:
            for event in self._events:
                event.set()

    def __status__(self) -> _Status:
        return self._value

    async def __check__(self) -> _Status:
        return self._value

    async def __subscribe__(self) -> asyncio.Event:
        event = asyncio.Event()
        self._events.add(event)
        return event

    async def __unsubscribe__(self, event: asyncio.Event):
        self._events.discard(event)
