from typing import Dict, Any, Awaitable, Callable
from aiogram import BaseMiddleware
from aiogram.types import TelegramObject
from config.config import Environ


class EnvMiddleware(BaseMiddleware):
    def __init__(self, env: Environ):
        self.env = env

    async def __call__(
            self,
            handler: Callable[
                [TelegramObject, Dict[str, Any]], Awaitable[Any]],
            event: TelegramObject,
            data: Dict[str, Any]
    ) -> Any:

            data["env"] = self.env

            return await handler(event, data)