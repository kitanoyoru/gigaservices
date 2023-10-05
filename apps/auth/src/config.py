from typing import Any, Callable, Iterable, Optional, Union
from hydra.core.config_store import ConfigStore

from dataclasses import dataclass, field, fields


class _DefaultType:
    def __repr__(self) -> str:
        return "<default>"


_DEFAULT = _DefaultType()

_ValidatorType = Callable[[str, Any], None]

_ConfigurationType = TypeVar("_ConfigurationType")

_GRPC_WORKER_MIN = 4
_GRPC_WORKER_MAX = 36


def _optional(validator: _ValidatorType) -> _ValidatorType:
    def proc(name: str, value: Any) -> None:
        if value is not None:
            validator(name, value)

    return proc


def _chain(*validators: _ValidatorType) -> _ValidatorType:
    def proc(name: str, value: Any) -> None:
        for validator in validators:
            validator(name, value)

    return proc


def _of_type(*types: type) -> _ValidatorType:
    def proc(name: str, value: Any) -> None:
        if not isinstance(value, types):
            types_repr = " or ".join(str(t) for t in types)
            raise TypeError(f'"{name}" should be of type {types_repr}')

    return proc


def _positive(name: str, value: Union[float, int]) -> None:
    if value <= 0:
        raise ValueError(f'"{name}" should be positive')


def _non_negative(name: str, value: Union[float, int]) -> None:
    if value < 0:
        raise ValueError(f'"{name}" should not be negative')


def _range(min_: int, max_: int) -> _ValidatorType:
    def proc(name: str, value: Union[float, int]) -> None:
        if value < min_:
            raise ValueError(f'"{name}" should be higher or equal to {min_}')
        if value > max_:
            raise ValueError(f'"{name}" should be less or equal to {max_}')

    return proc

def _validate(config: 'AppConfig') -> None:
    for f in fields(config):
        validate_fn = f.metadata.get('validate')
        if validate_fn is not None:
            value = getattr(config, f.name)
            if value is not _DEFAULT:
                validate_fn(f.name, value)


@dataclass(slots=True)
class DatabaseConfig:
    url: str
    username: str
    password: str


@dataclass(slots=True)
class CacheConfig:
    url: str
    username: str
    password: str


@dataclass(slots=True)
class MessageBrokerConfig:
    brokers_url: list[str]


@dataclass(slots=True)
class AppConfig:
    db: DatabaseConfig
    cache: CacheConfig
    broker: MessageBrokerConfig

    port: Optional[int] = field(
        default=16602,
		metadata={
			"validate": _chain(_of_type(int), _positive),
		}
    )

    max_grpc_workers: Optional[int] = field(
		default=12
		metadata= ={
			"validate": _chain(_optional, _of_type(int, _positive, _range(_GRPC_WORKER_MIN, _GRPC_WORKER_MAX))),
		}
	)

	def __post_init__(self):
		_validate(self)



cs = ConfigStore.instance()

cs.store(name="app_config", node=AppConfig)

cs.store(name="database_config", node=DatabaseConfig)
cs.store(name="cache_config", node=CacheConfig)
cs.store(name="message_broker_config", node=MessageBrokerConfig)
