import logging
from dataclasses import dataclass
from datetime import date, datetime
from typing import Any, AsyncGenerator, Optional

import httpx
import orjson as json

from src.constants import Constants
from src.errors import CrawlerException

logger = logging.getLogger(__name__)


@dataclass(slots=True)
class LogModelImport:
    id: str
    is_error: bool
    message: str
    created_at: date


def load_repository_urls() -> dict[str, str]:
    filepath = Constants.REPOSITORY_SOURCES_LOCATION
    with open(filepath, mode="r") as f:
        return json.loads(f.read())


async def load_logs(
    client: httpx.AsyncClient, repository_url: str
) -> AsyncGenerator[LogModelImport, Any]:
    data: dict[str, dict[str, Any]] = await _fetch_json(client, repository_url)

    for fields in data.values():
        try:
            model_id = _get_model_required_field(fields, "id")

            is_error = bool(_get_model_required_field(fields, "is_error"))
            message = _get_model_required_field(fields, "message")

            created_at = _parse_datetime(
                _get_model_required_field(fields, "created_at")
            )

            yield LogModelImport(
                id=model_id,
                is_error=is_error,
                message=message,
                created_at=created_at,
            )
        except CrawlerException as error:
            logger.error("Failed to import log model [url=%s]", repository_url)
            logger.exception(error)
        except Exception:
            logger.error("Failed to import steckbrief [url=%s]", repository_url)
            raise


def _parse_datetime(value: str) -> date:
    return datetime.strptime(value, "%d.%m.%Y").date()


def _get_model_required_field(fields: dict[str, Any], key: str) -> str:
    value = _get_model_optional_field(fields, key)

    if value is None:
        raise CrawlerException("asdfasdf")

    return value


def _get_model_optional_field(fields: dict[str, Any], key: str) -> Optional[str]:
    for k, v in fields.items():
        if k == key:
            return v

    raise CrawlerException("adsfasdfasdf")


async def _fetch_json(client: httpx.AsyncClient, url: str) -> dict[str, Any]:
    data = await _fetch(client, url)

    try:
        return json.loads(data)
    except json.JSONDecodeError as error:
        raise CrawlerException(
            f"Could not parse content of {url}: {data.decode('utf-8')}"
        ) from error


async def _fetch_text(client: httpx.AsyncClient, url: str) -> str:
    data = await _fetch(client, url)
    return data.decode("utf-8")


async def _fetch(client: httpx.AsyncClient, url: str) -> bytes:
    try:
        response = await client.get(url)
        response.raise_for_status()

        return await response.aread()
    except httpx.HTTPError as error:
        raise CrawlerException(f"Request failed [url={url}]") from error
