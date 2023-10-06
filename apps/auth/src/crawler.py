


def load_repository_urls() -> dict[str, str]:
    filepath = Constants.REPOSITORY_SOURCES_LOCATION 
    with open(filepath, mode="r") as f:
        return json.loads(f.read())

