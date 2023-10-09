import asyncio
from itertools import chain
from typing import TYPE_CHECKING, Any, Collection, Dict, Mapping, Optional, Set

import grpc
from src.proto.health.v1.health_pb2 import HealthCheckRequest  # type: ignore
from src.proto.health.v1.health_pb2 import HealthCheckResponse
from src.proto.health.v1.health_pb2_grpc import HealthServicer

if TYPE_CHECKING:
    from src.types import ICheckable

    from .check import CheckBase


def _status(checks: Set["CheckBase"]) -> "HealthCheckResponse.ServingStatus.ValueType":  # type: ignore
    statuses = {check.__status__() for check in checks}
    if statuses == {None}:
        return HealthCheckResponse.SERVING_STATUS_UNKNOWN
    elif statuses == {True}:
        return HealthCheckResponse.SERVING_STATUS_SERVING
    else:
        return HealthCheckResponse.SERVING_STATUS_NOT_SERVING


def _service_name(service: "ICheckable") -> str:
    methods = service.__mapping__()

    method_name = next(iter(methods), None)
    assert method_name is not None
    _, service_name, _ = method_name.split("/")

    return service_name


def _reset_waits(
    events: Collection[asyncio.Event],
    waits: Mapping[asyncio.Event, "asyncio.Task[bool]"],
) -> Dict[asyncio.Event, "asyncio.Task[bool]"]:
    new_waits = {}
    for event in events:
        wait = waits.get(event)
        if wait is None or wait.done():
            event.clear()
            wait = asyncio.ensure_future(event.wait())
        new_waits[event] = wait

    return new_waits


class _All:
    def __mapping__(self) -> Dict[str, Any]:
        return {"//": None}


ALL = _All()

_CheckConfig = Mapping["ICheckable", Collection["CheckBase"]]


class HealthService(HealthServicer):
    def __init__(self, checks: Optional[_CheckConfig] = None):
        if not checks:
            checks = {ALL: []}
        elif ALL not in checks:
            checks = dict(checks)
            checks[ALL] = list(chain.from_iterable(checks.values()))

        self._checks = {
            _service_name(s): set(check_list) for s, check_list in checks.items()
        }

    async def Check(self, request: HealthCheckRequest, _context: grpc.ServicerContext) -> HealthCheckResponse:  # type: ignore
        checks = self._checks.get(request.service)

        if checks is None:
            return HealthCheckResponse(
                status=HealthCheckResponse.SERVING_STATUS_UNKNOWN
            )
        elif len(checks) == 0:
            return HealthCheckResponse(
                status=HealthCheckResponse.SERVING_STATUS_SERVING
            )
        else:
            for check in checks:
                await check.__check__()
            return HealthCheckResponse(status=_status(checks))

    async def Watch(self, request: HealthCheckRequest, context: grpc.ServicerContext) -> HealthCheckResponse:  # type: ignore
        checks = self._checks.get(request.service)

        if checks is None:
            yield HealthCheckResponse(status=HealthCheckResponse.SERVING_STATUS_UNKNOWN)
        elif len(checks) == 0:
            yield HealthCheckResponse(status=HealthCheckResponse.SERVING_STATUS_SERVING)
        else:
            events = []
            for check in checks:
                events.append(await check.__subscribe__())

            waits = _reset_waits(events, {})

            try:
                yield HealthCheckResponse(status=_status(checks))
                while True:
                    await asyncio.wait(
                        waits.values(), return_when=asyncio.FIRST_COMPLETED
                    )
                    waits = _reset_waits(events, waits)
                    yield HealthCheckResponse(status=_status(checks))
            finally:
                for check, event in zip(checks, events):
                    await check.__unsubscribe__(event)
                for wait in waits.values():
                    if not wait.done():
                        wait.cancel()
