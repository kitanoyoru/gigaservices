from dataclasses import dataclass


@dataclass(slots=True)
class Constants:
	CONFIG_PATH = "/etc/"
	CONFIG_NAME = "kita-authservice"
	CONFIG_FULL_PATH = "/etc/kita-authservice.yaml"


