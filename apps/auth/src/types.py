from typing import Any, Mapping, Protocol


class ICheckable(Protocol):
    def __mapping__(self) -> Mapping[str, Any]:
        ...
