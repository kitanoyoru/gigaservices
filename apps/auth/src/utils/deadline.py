from __future__ import annotations

import time


class Deadline:
    def __init__(self, *, timestamp: float):
        self._timestamp = timestamp

    def __lt__(self, other: object) -> bool:
        if not isinstance(other, Deadline):
            raise TypeError(
                f"Comparison is not supported between {type(self).__name__} and {type(other).__name__}"
            )

        return self._timestamp < other._timestamp

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, Deadline):
            raise TypeError(
                f"Comparison is not supported between {type(self).__name__} and {type(other).__name__}"
            )

        return self._timestamp == other._timestamp

    @classmethod
    def from_timeout(cls, timeout: float) -> Deadline:
        return cls(timestamp=time.monotonic() + timeout)

    def time_remaining(self) -> float:
        return max(0, self._timestamp - time.monotonic())
