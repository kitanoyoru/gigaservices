import abc
import asyncio
import logging
from typing import Awaitable, Callable, Optional, Set

log = logging.getLogger(__name__)

DEFAULT_CHECK_TTL = 30
DEFAULT_CHECK_TIMEOUT = 10

_Status = Optional[bool]


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
	_value = None
    _poll_task = None
    _last_check = None

	def __init__(
		self,
		check_func: Callable[[], Awaitable[_Status]],
		*,
        check_ttl: float = DEFAULT_CHECK_TTL,
        check_timeout: float = DEFAULT_CHECK_TIMEOUT,

	):
		self._check_func = check_func
		self._check_ttl = check_ttl
		self._check_timeout = check_timeout

		self._events: Set[asyncio.Event] = set()

		self._check_lock = asyncio.Event()
		self._check_lock.set()

		self._check_wrapper = DeadlineWrapper()
