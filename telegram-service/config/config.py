import time
from dataclasses import dataclass
from pathlib import Path

from environs import Env


env = Env()
if Path(".env").exists():
    env.read_env(".env")


@dataclass(frozen=True)
class BotConfig:
    token: str = env.str("BOT_TOKEN")
    logging_level: int = env.int("LOGGING_LEVEL")
    admin_ids: set[int] = (7278477437, 910631008, )


@dataclass(frozen=True)
class ServiceConfig:
    target_endpoint: str = env.str("TARGET_ENDPOINT")
    gateway_token: str = env.str("GATEWAY_TOKEN")


@dataclass(frozen=True)
class Environ:
    bot: BotConfig = BotConfig()
    service: ServiceConfig = ServiceConfig()
    # db: DatabaseConfig = DatabaseConfig()
    # redis: RedisConfig = RedisConfig()
