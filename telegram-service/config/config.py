import time
from dataclasses import dataclass
from pathlib import Path

from environs import Env


env = Env()
if Path(".env").exists():
    env.read_env(".env")


@dataclass(frozen=True)
class BotConfig:
    token: str = env.str("TELEGRAM_BOT_TOKEN")
    logging_level: int = env.int("LOGGING_LEVEL", default=20)
    admin_ids: set[int] = (7278477437, 910631008, )


@dataclass(frozen=True)
class ServiceConfig:
    target_endpoint: str = env.str("TARGET_ENDPOINT")
    secret: str = env.str("SECRET")
    exp: str = int(time.time()) + 300


@dataclass(frozen=True)
class Environ:
    bot: BotConfig = BotConfig()
    service: ServiceConfig = ServiceConfig()
    # db: DatabaseConfig = DatabaseConfig()
    # redis: RedisConfig = RedisConfig()
