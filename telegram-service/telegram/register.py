from aiogram import Dispatcher, Bot
from apscheduler.schedulers.asyncio import AsyncIOScheduler

from config.config import Environ
from telegram.handlers import user
from telegram.middlewares.env_middleware import EnvMiddleware


class TgRegister:
    def __init__(self, dp: Dispatcher, bot: Bot, env: Environ):
        self.dp = dp
        self.bot = bot

        self.env = env

    async def register(self):
        #self._create_scheduler()

        self._register_handlers()
        self._register_middlewares()
        #self._register_tasks()

    def _register_handlers(self):
        self.dp.include_routers(user.router)

    def _register_middlewares(self):
        env_middleware = EnvMiddleware(self.env)

        self.dp.callback_query.middleware(env_middleware)

        self.dp.message.middleware(env_middleware)

        self.dp.inline_query.middleware(env_middleware)